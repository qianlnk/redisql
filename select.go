package redisql

import (
	"errors"
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"strings"
)

/*
create table user(
id int(11) auto_increment PRIMARY KEY,
name varchar(64) not null default '',
age int(11) not null default 0,
city varchar(32) not null default ''
)

create table log(
id int(11) auto_increment PRIMARY KEY,
userid int(11) not null default 0,
operate varchar(128) not null default '',
detail VARCHAR(128)not null default '',
`date` datetime not null default '1900-01-01 00:00:00'
)
*/

//how to use
//eg: FIELDS("a.name, a.age, a.city, b.userid, b.operate, b.detail, b.data")
//	  .FROM("user, log").ALIAS("a, b")
//	  .WHERE(a.id = b.userid AND a.age = 24) OR (b.operate = "login").SELECT()
type Inner struct {
	Inner string
	Alias string
	On    string
	And   string
	Or    string
}

type Left struct {
	Left  string
	Alias string
	On    string
	And   string
	Or    string
}

type Right struct {
	Right string
	Alias string
	On    string
	And   string
	Or    string
}

type SltTable map[string]string

type SltField struct {
	TableAlias string
	Name       string
	Alias      string
}

var (
	CompareSign = []string{"=", "!=", ">", ">=", "<", "<=", "like"}
	Opertion    = []string{"(", ")", "and", "or"}
)

type Select struct {
	Froms  SltTable   //query table must have an alias
	Fields []SltField //query field must like tablealias.fieldname, it can have an alias also.
	Where  []string
	Top    int
	Limit  []int
}

func FROM(tables ...string) *Select {
	tmptables := make(SltTable)
	for _, t := range tables {
		tabs := strings.Split(t, ",")
		for _, tt := range tabs {
			tas := strings.Fields(tt)
			if len(tas) != 2 {
				panic(errors.New("tables wrong."))
			} else {
				if v, ok := tmptables[tas[0]]; ok {
					panic(fmt.Sprintf("alias '％s' is used by '%s', can't as table %s's alias again.", tas[0], v, tas[1]))
				} else {
					tmptables[tas[1]] = tas[0]
				}
			}
		}
	}

	return &Select{
		Froms: tmptables,
	}
}

func (slt *Select) FIELDS(fields ...string) *Select {
	var tmpfields []SltField
	for _, f := range fields {
		fs := strings.Split(f, ",")
		for _, ff := range fs {
			var field SltField
			fas := strings.Fields(ff)
			tabField := strings.Split(fas[0], ".")
			if len(tabField) != 2 {
				panic(errors.New("fields wrong1."))
			} else {
				if len(fas) == 2 {
					field.TableAlias = tabField[0]
					field.Name = tabField[1]
					field.Alias = fas[1]
				} else if len(fas) == 1 { //if no define alias,field's alias is it self
					field.TableAlias = tabField[0]
					field.Name = tabField[1]
					field.Alias = tabField[1]
				} else {
					panic(errors.New("fields wrong2."))
				}
				tmpfields = append(tmpfields, field)
			}
		}
	}

	return &Select{
		Froms:  slt.Froms,
		Fields: tmpfields,
		Where:  slt.Where,
		Top:    slt.Top,
		Limit:  slt.Limit,
	}
}

func (slt *Select) WHERE(condition string) *Select {
	//data value can't include data as follow
	specialChar := []string{"=", "!=", ">", ">=", "<", "<=", "(", ")", "and(", "or(", ")and", ")or"}
	var temCondition []string
	tmpcdt := strings.ToLower(condition)
	for _, c := range specialChar {
		tmpcdt = strings.Replace(tmpcdt, c, " "+c+" ", -1)
	}
	temCondition = strings.Fields(tmpcdt)
	return &Select{
		Froms:  slt.Froms,
		Fields: slt.Fields,
		Where:  temCondition,
		Top:    slt.Top,
		Limit:  slt.Limit,
	}
}

func (slt *Select) TOP(top int) *Select {
	return &Select{
		Froms:  slt.Froms,
		Fields: slt.Fields,
		Where:  slt.Where,
		Top:    top,
		Limit:  slt.Limit,
	}
}

func (slt *Select) LIMIT(start, end int) *Select {
	if start < 0 {
		panic("limit start can't less than 0.")
	}
	var limit []int
	if start > end && end > 0 {
		limit[0] = end
		limit[1] = start
	} else {
		limit[0] = start
		limit[1] = end
	}
	return &Select{
		Froms:  slt.Froms,
		Fields: slt.Fields,
		Where:  slt.Where,
		Top:    slt.Top,
		Limit:  limit,
	}
}

func (slt *Select) check() error {
	//check table
	for _, t := range slt.Froms {
		if existsTable(t) == false {
			return errors.New(fmt.Sprintf("table '%s' not exists.", t))
		}
	}
	//check field
	for _, f := range slt.Fields {
		var tmptab string
		if v, ok := slt.Froms[f.TableAlias]; ok {
			tmptab = v
		}

		if tmptab == "" {
			return errors.New(fmt.Sprintf("Unknow alias '%s' int table list.", f.TableAlias))
		} else {
			if existsField(tmptab, f.Name) == false {
				return errors.New(fmt.Sprintf("Unknow cloumn '%s.%s' int field list.", f.TableAlias, f.Name))
			}
		}
	}

	return nil
}

//support single table query first，not support select * now.
func (slt *Select) SELECT() ([]byte, error) {
	fmt.Println(slt)
	err := slt.check()
	if err != nil {
		return nil, err
	}

	conn := getConn()
	defer conn.Close()

	var ids map[string][]string
	if len(slt.Where) == 0 {
		for k, v := range slt.Froms {
			tmpdatas, err := redigo.Strings(conn.Do("KEYS", fmt.Sprintf(REDISQL_DATAS, database, v, "*")))
			if err != nil {
				return nil, err
			}
			for _, tmp := range tmpdatas {
				ids[k] = append(ids[k], strings.Split(tmp, ".")[len(strings.Split(tmp, "."))-1])
			}
		}
	} else {
		//get ids
		esStack := new_stack()
		snStack := new_stack()
		snStack.PUSH("#")
		for _, val := range slt.Where {
			if inarray(Opertion, val) == false && inarray(CompareSign, val) == false {
				if inarray(CompareSign, snStack.GetPOP()) == false {
					esStack.PUSH(val)
				} else {
					left := esStack.POP()
					opt := snStack.POP()
					right := val
					tmpids, err := slt.getDataIds(left, opt, right)
					if err != nil {
						return nil, err
					}
					var idstring string
					for _, v := range tmpids {
						idstring = idstring + " " + v
					}
					esStack.PUSH(idstring)
				}
			}
		}
	}

	//	} else if slt.And == nil && slt.Or == nil {
	//		//get data id
	//		indexData := fmt.Sprintf("%s.%s", slt.Where.Field, slt.Where.Value)
	//		ids, err = redigo.Strings(conn.Do("SMEMBERS", fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[slt.Where.TableAlias], indexData)))
	//		if err != nil {
	//			return nil, err
	//		}
	//	}

	//	var fields []interface{}
	//	for _, f := range slt.Fields {
	//		fields = append(fields, f.Name)
	//	}

	//	var datas [][]string
	//	for _, id := range ids {
	//		var params []interface{}
	//		params = append(params, fmt.Sprintf(REDISQL_DATAS, database, slt.Froms[slt.Fields[0].TableAlias], id))
	//		params = append(params, fields...)
	//		//fmt.Println(params)
	//		datas1, err := redigo.Strings(conn.Do("HMGET", params...))
	//		if err != nil {
	//			return nil, err
	//		}
	//		datas = append(datas, datas1)
	//	}
	//	var resstring string
	//	resstring += "{"
	//	for i, f := range slt.Fields {
	//		if i > 0 {
	//			resstring += ","
	//		}
	//		resstring += fmt.Sprintf("\"%s\":[", f.Alias)
	//		for j, data := range datas {
	//			if j > 0 {
	//				resstring += ","
	//			}
	//			resstring += fmt.Sprintf("\"%s\"", data[i])
	//		}
	//		resstring += "]"
	//	}
	//	resstring += "}"

	//	return []byte(resstring), nil
	return nil, nil
}

func (slt *Select) getDataIds(left, option, right string) ([]string, error) {
	conn := getConn()
	defer conn.Close()
	fields := strings.Split(left, ".")
	if existsTable(slt.Froms[fields[0]]) == false {
		return nil, errors.New(fmt.Sprintf("table %s not exist.", slt.Froms[fields[0]]))
	}
	if existsField(slt.Froms[fields[0]], fields[1]) == false {
		return nil, errors.New(fmt.Sprintf("field %s not found in table %s.", fields[1], slt.Froms[fields[0]]))
	}
	fieldtype, err := getFieldType(slt.Froms[fields[0]], fields[1])
	if err != nil {
		return nil, err
	}
	if fieldtype == REDISQL_TYPE_STRING {
		switch option {
		case "=":
			indexdata := fmt.Sprintf("%s.%s", fields[1], right)
			ids, err := redigo.Strings(conn.Do("SMEMBERS", fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[fields[0]], indexdata)))
			if err != nil {
				return nil, err
			}
			return ids, nil
		}
	} else {
		return nil, nil
	}
	return nil, nil
}
