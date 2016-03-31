package redisql

import (
	"errors"
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"time"
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
	Opertion    = []string{"(", ")", "and", "or", "#"}
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
	specialChar1 := []string{"!=", ">=", "<="}
	specialChar2 := []string{"=", ">", "<", "(", ")"}
	specialChar3 := []string{"! =", ">  =", "<  ="}
	var temCondition []string
	tmpcdt := strings.ToLower(condition)
	for _, c := range specialChar1 {
		tmpcdt = strings.Replace(tmpcdt, c, " "+c+" ", -1)
	}
	for _, c := range specialChar2 {
		tmpcdt = strings.Replace(tmpcdt, c, " "+c+" ", -1)
	}
	for i, c := range specialChar3 {
		tmpcdt = strings.Replace(tmpcdt, c, specialChar1[i], -1)
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
		tmpWhere := append(slt.Where, "#")
		for _, val := range tmpWhere {
			fmt.Println(val)
			if inarray(Opertion, val) == false && inarray(CompareSign, val) == false {
				if inarray(CompareSign, snStack.GetPOP()) == false {
					esStack.PUSH(val)
				} else {
					left := esStack.POP()
					opt := snStack.POP()
					right := val
					ok, err := slt.judgeRight(left, right)
					if err != nil {
						return nil, err
					}
					if ok {
						alias, tmpids, err := slt.getDataIds(left, opt, right)
						if err != nil {
							return nil, err
						}
						var idstring string
						idstring = alias
						for _, v := range tmpids {
							idstring = idstring + " " + v
						}
						fmt.Println(idstring)
						esStack.PUSH(idstring)
					} else {
						tmpEs := left + " " + opt + " " + right
						esStack.PUSH(tmpEs)
					}
				}
			} else if inarray(CompareSign, val) == true {
				snStack.PUSH(val)
			} else if inarray(Opertion, val) == true {
				res := Compare(snStack.GetPOP(), val)
				fmt.Println(snStack.GetPOP(), val, "res=", res)
				switch res {
				case REDISQL_PRIORITY_LESS:
					snStack.PUSH(val)
					break
				case REDISQL_PRIORITY_EQUAL:
					snStack.POP()
					break
				case REDISQL_PRIORITY_GREATER:
					right := esStack.POP()
					left := esStack.POP()
					opt := snStack.POP()
					fmt.Printf("%s %s %s", left, opt, right)
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

func (slt *Select) getDataIds(left, option, right string) (string, []string, error) {
	conn := getConn()
	defer conn.Close()
	fields := strings.Split(left, ".")
	if len(fields) != 2 {
		return "", nil, errors.New(fmt.Sprintf("unknow field '%s'.", left))
	}
	if existsTable(slt.Froms[fields[0]]) == false {
		return "", nil, errors.New(fmt.Sprintf("table %s not exist.", slt.Froms[fields[0]]))
	}
	if existsField(slt.Froms[fields[0]], fields[1]) == false {
		return "", nil, errors.New(fmt.Sprintf("field %s not found in table %s.", fields[1], slt.Froms[fields[0]]))
	}
	fieldtype, err := getFieldType(slt.Froms[fields[0]], fields[1])
	if err != nil {
		return "", nil, err
	}
	if fieldtype == REDISQL_TYPE_STRING {
		tmpRight := strings.Replace(right, "'", "", -1)
		switch option {
		case "=":
			index := fmt.Sprintf("%s.%s", fields[1], tmpRight)
			key := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[fields[0]], index)
			res, err := redigo.Strings(conn.Do("SMEMBERS", key))
			return fields[0], res, err
		case "!=":
			index := fmt.Sprintf("%s.%s", fields[1], tmpRight)
			key := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[fields[0]], index)
			outids, err := redigo.Strings(conn.Do("SMEMBERS", key))
			if err != nil {
				return "", nil, err
			}
			tmpkeys, err := redigo.Strings(conn.Do("KEYS", fmt.Sprintf(REDISQL_DATAS, database, slt.Froms[fields[0]], "*")))
			if err != nil {
				return "", nil, err
			}
			var ids []string
			for _, tmp := range tmpkeys {
				ids = append(ids, strings.Split(tmp, ".")[len(strings.Split(tmp, "."))-1])
			}
			return fields[0], outer(ids, outids), nil
		case "like":
			index := fmt.Sprintf("%s.%s", fields[1], strings.Replace(tmpRight, "%", "*", -1))
			key := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[fields[0]], index)
			fmt.Println(fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[fields[0]], fmt.Sprintf("%s.%s", fields[1], strings.Replace(strings.Replace(right, "'", "", -1), "%", "*", -1))))
			tmpkeys, err := redigo.Strings(conn.Do("KEYS", key))
			if err != nil {
				return "", nil, err
			}
			var ids []string
			for _, tmp := range tmpkeys {
				tmpids, err := redigo.Strings(conn.Do("SMEMBERS", tmp))
				if err != nil {
					return "", nil, err
				}
				ids = append(ids, tmpids...)
			}
			return fields[0], ids, nil
		default:
			return "", nil, errors.New("type string not support sign " + option)
		}
	} else {
		var score string
		if fieldtype == REDISQL_TYPE_NUMBER {
			_, err = strconv.Atoi(right)
			if err != nil {
				return "", nil, err
			}
			score = right
		} else {
			tmpRight := strings.Replace(right, "'", "", -1)
			t, err := time.Parse("2006-01-02 15:04:05", tmpRight)
			if err != nil {
				return "", nil, err
			}
			score = t.Format("20060102150405")
		}
		fmt.Printf("score:%s\n", score)
		key := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[fields[0]], fields[1])
		switch option {
		case "=":
			res, err := redigo.Strings(conn.Do("ZRANGEBYSCORE", key, score, score))
			return fields[0], res, err
		case "!=":
			ids1, err := redigo.Strings(conn.Do("ZRANGEBYSCORE", key, "-inf", "("+score))
			if err != nil {
				return "", nil, err
			}
			ids2, err := redigo.Strings(conn.Do("ZRANGEBYSCORE", key, "("+score, "+inf"))
			if err != nil {
				return "", nil, err
			}
			return fields[0], union(ids1, ids2), nil
		case ">":
			res, err := redigo.Strings(conn.Do("ZRANGEBYSCORE", key, "("+score, "+inf"))
			return fields[0], res, err
		case ">=":
			res, err := redigo.Strings(conn.Do("ZRANGEBYSCORE", key, score, "+inf"))
			return fields[0], res, err
		case "<":
			res, err := redigo.Strings(conn.Do("ZRANGEBYSCORE", key, "-inf", "("+score))
			return fields[0], res, err
		case "<=":
			res, err := redigo.Strings(conn.Do("ZRANGEBYSCORE", key, "-inf", score))
			return fields[0], res, err
		default:
			return "", nil, errors.New("type number or date not support sign " + option)
		}
	}
	return "", nil, nil
}

//return true when right is value, return false where right is a table field.
func (slt *Select) judgeRight(left, right string) (bool, error) {
	if strings.Contains(right, "'") == true { //string or time, include "'" must be value
		return true, nil
	} else {
		tmp := strings.Split(left, ".")
		if len(tmp) != 2 {
			return false, errors.New(fmt.Sprintf("unknow field '%s'.", left))
		}
		fieldtype, err := getFieldType(slt.Froms[tmp[0]], tmp[1])
		if err != nil {
			return false, err
		}
		if fieldtype == "" {
			return false, errors.New(fmt.Sprintf("field %s not found in table %s.", tmp[1], slt.Froms[tmp[0]]))
		}
		if fieldtype == REDISQL_TYPE_NUMBER {
			_, err := strconv.Atoi(right)
			if err != nil {
				return false, err
			} else {
				return true, nil
			}
		} else {
			rights := strings.Split(right, ".")
			if len(rights) != 2 {
				return false, errors.New(fmt.Sprintf("unknow field '%s'.", right))
			} else {
				if existsTable(slt.Froms[rights[0]]) == false {
					return false, errors.New(fmt.Sprintf("table %s not exist.", slt.Froms[rights[0]]))
				}
				if existsField(slt.Froms[rights[0]], rights[1]) == false {
					return false, errors.New(fmt.Sprintf("field %s not found in table %s.", rights[1], slt.Froms[rights[0]]))
				}
				return false, nil
			}
		}
	}
	return false, errors.New("unknow, no more message.")
}
