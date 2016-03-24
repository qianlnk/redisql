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
//	  .WHERE(a.id = b.userid).AND(a.age = 24).OR(b.operate = "login").SELECT()
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

type SingleCondition struct {
	TableAlias string
	Field      string
	Compare    string
	Value      string
}

var (
	CompareSign = [...]string{"=", "!=", ">", ">=", "<", "<="}
)

type ComplexCondition struct {
	SgCdt []SingleCondition
	Union string //or | and
}

type Condition struct {
	CplCdt []ComplexCondition
	Union  string
}

type Select struct {
	Froms  SltTable   //query table must have an alias
	Fields []SltField //query field must like tablealias.fieldname, it can have an alias also.
	Where  SingleCondition
	And    []SingleCondition
	Or     []SingleCondition
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
	}
}

func (slt *Select) WHERE(condition string) *Select {
	var temCondition SingleCondition
	cdt := strings.Fields(condition)
	if len(cdt) != 3 {
		panic(errors.New(fmt.Sprintf("wrong condition: '%s'.", condition)))
	}
	tabField := strings.Split(cdt[0], ".")
	if len(tabField) != 2 {
		panic(errors.New(fmt.Sprintf("wrong condition: '%s'.", condition)))
	} else {
		temCondition.TableAlias = tabField[0]
		temCondition.Field = tabField[1]
	}

	for _, cp := range CompareSign {
		if cdt[1] == cp {
			temCondition.Compare = cdt[1]
		}
	}

	if temCondition.Compare == "" {
		panic(errors.New(fmt.Sprintf("wrong compare sign '%s'.", cdt[1])))
	}

	value := strings.Replace(cdt[2], "'", "", -1)
	temCondition.Value = value

	return &Select{
		Froms:  slt.Froms,
		Fields: slt.Fields,
		Where:  temCondition,
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

	//check where
	if slt.Where.Field != "" {
		if tab, ok := slt.Froms[slt.Where.TableAlias]; ok {
			if existsField(tab, slt.Where.Field) == false {
				return errors.New(fmt.Sprintf("Unknow cloumn '%s.%s' int field list.", slt.Where.TableAlias, slt.Where.Field))
			} else {
				if existsIndex(tab, fmt.Sprintf("index_%s", slt.Where.Field)) == false {
					return errors.New(fmt.Sprintf("not index for clumn '%s'.", slt.Where.Field))
				}
			}
		} else {
			return errors.New(fmt.Sprintf("Unknow alias '%s' int table list used in condition.", slt.Where.TableAlias))
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

	var ids []string
	if slt.Where.Field == "" {
		for _, v := range slt.Froms {
			tmpdatas, err := redigo.Strings(conn.Do("KEYS", fmt.Sprintf(REDISQL_DATAS, database, v, "*")))
			if err != nil {
				return nil, err
			}
			for _, v := range tmpdatas {
				ids = append(ids, strings.Split(v, ".")[len(strings.Split(v, "."))-1])
			}
		}
	} else if slt.And == nil && slt.Or == nil {
		//get data id
		indexData := fmt.Sprintf("%s.%s", slt.Where.Field, slt.Where.Value)
		ids, err = redigo.Strings(conn.Do("SMEMBERS", fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[slt.Where.TableAlias], indexData)))
		if err != nil {
			return nil, err
		}
	}

	var fields []interface{}
	for _, f := range slt.Fields {
		fields = append(fields, f.Name)
	}

	var datas [][]string
	for _, id := range ids {
		var params []interface{}
		params = append(params, fmt.Sprintf(REDISQL_DATAS, database, slt.Froms[slt.Fields[0].TableAlias], id))
		params = append(params, fields...)
		//fmt.Println(params)
		datas1, err := redigo.Strings(conn.Do("HMGET", params...))
		if err != nil {
			return nil, err
		}
		datas = append(datas, datas1)
	}
	var resstring string
	resstring += "{"
	for i, f := range slt.Fields {
		if i > 0 {
			resstring += ","
		}
		resstring += fmt.Sprintf("\"%s\":[", f.Alias)
		for j, data := range datas {
			if j > 0 {
				resstring += ","
			}
			resstring += fmt.Sprintf("\"%s\"", data[i])
		}
		resstring += "]"
	}
	resstring += "}"

	return []byte(resstring), nil
}
