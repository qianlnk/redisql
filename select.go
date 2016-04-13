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
	limit := make([]int, 2)
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

func (slt *Select) getEsDataIds(cartesianKey, left, opt, right string) (string, error) {
	fmt.Printf("getEsDataIds start, cartesianKey = %s, left = %s, opt = %s, right = %s\n", cartesianKey, left, opt, right)
	conn := getConn()
	defer conn.Close()

	lefts := strings.Split(left, ".")
	if len(lefts) != 2 {
		return "", errors.New(fmt.Sprintf("unknow field '%s'.", left))
	}
	if existsTable(slt.Froms[lefts[0]]) == false {
		return "", errors.New(fmt.Sprintf("table %s not exist.", slt.Froms[lefts[0]]))
	}
	if existsField(slt.Froms[lefts[0]], lefts[1]) == false {
		return "", errors.New(fmt.Sprintf("field %s not found in table %s.", lefts[1], slt.Froms[lefts[0]]))
	}
	lefttype, err := getFieldType(slt.Froms[lefts[0]], lefts[1])
	if err != nil {
		return "", err
	}
	rights := strings.Split(right, ".")
	if len(rights) != 2 {
		return "", errors.New(fmt.Sprintf("unknow field '%s'.", right))
	}
	if existsTable(slt.Froms[rights[0]]) == false {
		return "", errors.New(fmt.Sprintf("table %s not exist.", slt.Froms[rights[0]]))
	}
	if existsField(slt.Froms[rights[0]], rights[1]) == false {
		return "", errors.New(fmt.Sprintf("field %s not found in table %s.", rights[1], slt.Froms[rights[0]]))
	}
	righttype, err := getFieldType(slt.Froms[rights[0]], rights[1])
	if err != nil {
		return "", err
	}
	if lefttype != righttype {
		return "", errors.New(fmt.Sprintf("can not use type '%s' = '%s'.", lefttype, righttype))
	}
	fmt.Println(lefttype, righttype)
	cdtSn, err := getNextConditionSn()
	if err != nil {
		return "", err
	}
	cdtKey := fmt.Sprintf(REDISQL_TMP_CONDITION, database, slt.Froms[lefts[0]]+"."+slt.Froms[rights[0]], strconv.Itoa(cdtSn))
	switch lefttype {
	case REDISQL_TYPE_NUMBER:
		if lefts[1] == "id" {
			key := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[rights[0]], rights[1])
			fmt.Println("key:", key)
			tmpids, err := redigo.Strings(conn.Do("ZRANGE", key, 0, -1, "withscores"))
			if err != nil {
				fmt.Println("err:", err)
				return "", err
			}
			for i := 1; i < len(tmpids); i = i + 2 {
				tmpkey1 := lefts[0] + "_" + tmpids[i] + "_" + rights[0] + "_" + tmpids[i-1] + "_"
				tmpkey2 := rights[0] + "_" + tmpids[i-1] + "_" + lefts[0] + "_" + tmpids[i] + "_"
				fmt.Println(cartesianKey, tmpkey1)
				hkeys, err := redigo.Strings(conn.Do("HKEYS", cartesianKey))
				if err != nil {
					return "", err
				}
				for _, hk := range hkeys {
					if strings.Contains(hk, tmpkey1) || strings.Contains(hk, tmpkey2) {
						_, err = conn.Do("SADD", cdtKey, hk)
						if err != nil {
							return "", err
						}
					}
				}
			}
		} else if rights[1] == "id" {
			key := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[lefts[0]], lefts[1])
			tmpids, err := redigo.Strings(conn.Do("ZRANGE", key, 0, -1, "withsocres"))
			if err != nil {
				return "", err
			}
			for i := 1; i < len(tmpids); i = i + 2 {
				tmpkey1 := rights[0] + "_" + tmpids[i] + "_" + lefts[0] + " " + tmpids[i-1] + "_"
				tmpkey2 := lefts[0] + " " + tmpids[i-1] + "_" + rights[0] + "_" + tmpids[i] + "_"
				hkeys, err := redigo.Strings(conn.Do("HKEYS", cartesianKey))
				if err != nil {
					return "", err
				}
				for _, hk := range hkeys {
					if strings.Contains(hk, tmpkey1) || strings.Contains(hk, tmpkey2) {
						_, err = conn.Do("SADD", cdtKey, hk)
						if err != nil {
							return "", err
						}
					}
				}
			}
		} else {
			leftkey := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[lefts[0]], lefts[1])
			leftids, err := redigo.Strings(conn.Do("ZRANGE", leftkey, 0, -1, "withsocres"))
			if err != nil {
				return "", err
			}
			for i := 1; i < len(leftids); i = i + 2 {
				rightkey := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[rights[0]], rights[1])
				rightids, err := redigo.Strings(conn.Do("ZRANGEBYSCORE", rightkey, leftids[i], leftids[i]))
				if err != nil {
					return "", err
				}
				for _, id := range rightids {
					tmpkey1 := lefts[0] + "_" + leftids[i-1] + "_" + rights[0] + "_" + id + "_"
					tmpkey2 := rights[0] + "_" + id + "_" + lefts[0] + "_" + leftids[i-1] + "_"
					hkeys, err := redigo.Strings(conn.Do("HKEYS", cartesianKey))
					if err != nil {
						return "", err
					}
					for _, hk := range hkeys {
						if strings.Contains(hk, tmpkey1) || strings.Contains(hk, tmpkey2) {
							_, err = conn.Do("SADD", cdtKey, hk)
							if err != nil {
								return "", err
							}
						}
					}
				}
			}
		}
		break
	case REDISQL_TYPE_STRING:
		leftkey := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[lefts[0]], lefts[1]+".*")
		tmpkeys, err := redigo.Strings(conn.Do("KEYS", leftkey))
		if err != nil {
			return "", err
		}
		for _, k := range tmpkeys {
			leftids, err := redigo.Strings(conn.Do("SMEMBERS", k))
			if err != nil {
				return "", err
			}
			values := strings.Split(k, ".")
			value := values[len(values)-1]
			rightkey := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[rights[0]], rights[1]+"."+value)
			rightids, err := redigo.Strings(conn.Do("SMEMBERS", rightkey))
			if err != nil {
				return "", err
			}
			if len(rightids) <= 0 || len(leftids) <= 0 {
				continue
			}
			for _, l := range leftids {
				for _, r := range rightids {
					tmpkey1 := lefts[0] + "_" + l + "_" + rights[0] + "_" + r + "_"
					tmpkey2 := rights[0] + "_" + r + "_" + lefts[0] + "_" + l + "_"
					hkeys, err := redigo.Strings(conn.Do("HKEYS", cartesianKey))
					if err != nil {
						return "", err
					}
					for _, hk := range hkeys {
						if strings.Contains(hk, tmpkey1) || strings.Contains(hk, tmpkey2) {
							_, err = conn.Do("SADD", cdtKey, hk)
							if err != nil {
								return "", err
							}
						}
					}
				}

			}
		}
		break
	case REDISQL_TYPE_DATE:
		leftkey := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[lefts[0]], lefts[1])
		leftids, err := redigo.Strings(conn.Do("ZRANGE", leftkey, 0, -1, "withsocres"))
		if err != nil {
			return "", err
		}
		for i := 1; i < len(leftids); i = i + 2 {
			rightkey := fmt.Sprintf(REDISQL_INDEX_DATAS, database, slt.Froms[rights[0]], rights[1])
			rightids, err := redigo.Strings(conn.Do("ZRANGEBYSCORE", rightkey, leftids[i], leftids[i]))
			if err != nil {
				return "", err
			}
			for _, id := range rightids {
				tmpkey1 := lefts[0] + "_" + leftids[i-1] + "_" + rights[0] + "_" + id + "_"
				tmpkey2 := rights[0] + "_" + id + "_" + lefts[0] + "_" + leftids[i-1] + "_"
				hkeys, err := redigo.Strings(conn.Do("HKEYS", cartesianKey))
				if err != nil {
					return "", err
				}
				for _, hk := range hkeys {
					if strings.Contains(hk, tmpkey1) || strings.Contains(hk, tmpkey2) {
						_, err = conn.Do("SADD", cdtKey, hk)
						if err != nil {
							return "", err
						}
					}
				}
			}
		}
		break
	default:
		return "", errors.New("data type err.")
	}
	return cdtKey, nil
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
			fmt.Println("tmpkeys:", tmpkeys)
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
			fmt.Println(tmpRight)
			t, err := time.Parse("20060102150405", tmpRight)
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

		if strings.Contains(right, ".") == true {
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
		} else {
			return true, nil
		}
	}
	return false, errors.New("unknow, no more message.")
}

func merge(left, opt, right string) string {
	if strings.Contains(left, "&") == false {
		if opt == "and" {
			return left + "&" + right
		} else if opt == "or" {
			lefttab := strings.Fields(left)
			righttab := strings.Fields(right)
			return left + "&" + righttab[0] + " *^" + lefttab[0] + " * &" + right
		}
	} else {

	}
	return ""
}

func (slt *Select) SELECT() ([]byte, error) {
	fmt.Println(slt)
	err := slt.check()
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	conn := getConn()
	defer conn.Close()
	defer clearTmp()

	tabNum := len(slt.Froms)
	if tabNum == 1 { //single table query
		var ids []string
		var err error
		if len(slt.Where) == 0 { //no condition
			//get all ids
			var tableName string
			for _, v := range slt.Froms {
				tableName = v
			}
			key := fmt.Sprintf(REDISQL_INDEX_DATAS, database, tableName, "id")
			ids, err = redigo.Strings(conn.Do("ZRANGE", key, 0, -1))
			if err != nil {
				fmt.Println("err:", err)
				return nil, err
			}
		} else { //single or more condition
			esStack := new_stack() //表达式栈
			snStack := new_stack() //关系符栈
			snStack.PUSH("#")      //start and end with #
			tmpWhere := append(slt.Where, "#")
			i := 0
			val := tmpWhere[i]
			for {
				if val == "#" && snStack.GetPOP() == "#" {
					break
				}
				fmt.Println("val:", val)
				if inarray(Opertion, val) == false && inarray(CompareSign, val) == false { //not opertion or relationship sign
					if inarray(CompareSign, snStack.GetPOP()) == false { //no relationship sign at top snStack
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
							fmt.Println("getDataIds:", alias, tmpids, err)
							if err != nil {
								return nil, err
							}
							//get condition sn
							cdtSn, err := getNextConditionSn()
							if err != nil {
								return nil, err
							}

							key := fmt.Sprintf(REDISQL_TMP_CONDITION, database, slt.Froms[alias], strconv.Itoa(cdtSn))
							for _, id := range tmpids {
								_, err = conn.Do("SADD", key, id)
								if err != nil {
									return nil, err
								}
							}
							esStack.PUSH(key)
						} else {
							return nil, errors.New("condition wrong.")
						}
					}
					i++
				} else if inarray(CompareSign, val) == true {
					snStack.PUSH(val)
					i++
				} else if inarray(Opertion, val) == true {
					res := Compare(snStack.GetPOP(), val)
					fmt.Println(snStack.GetPOP(), val, "res=", res)
					switch res {
					case REDISQL_PRIORITY_LESS:
						snStack.PUSH(val)
						i++
						break
					case REDISQL_PRIORITY_EQUAL:
						snStack.POP()
						i++
						break
					case REDISQL_PRIORITY_GREATER:
						right := esStack.POP()
						left := esStack.POP()
						opt := snStack.POP()
						fmt.Printf("%s %s %s\n", left, opt, right)
						reskey, err := singleMerge(left, opt, right)
						if err != nil {
							return nil, err
						}
						esStack.PUSH(reskey)
						break
					case REDISQL_PRIORITY_ERROR:
						return nil, errors.New("wrong ralationship.")
					}
				}
				val = tmpWhere[i]
			}
			mykey := esStack.POP()
			ids, err = redigo.Strings(conn.Do("SMEMBERS", mykey))
			if err != nil {
				return nil, err
			}

		}
		fmt.Println(ids)
		var fields []interface{}
		for _, f := range slt.Fields {
			fields = append(fields, f.Name)
		}

		var datas [][]string
		for _, id := range ids {
			var params []interface{}
			params = append(params, fmt.Sprintf(REDISQL_DATAS, database, slt.Froms[slt.Fields[0].TableAlias], id))
			params = append(params, fields...)
			fmt.Println(params)
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
				if slt.Top > 0 && j >= slt.Top {
					break
				}
				if slt.Limit != nil && j < slt.Limit[0] {
					continue
				}
				if slt.Limit != nil && j > slt.Limit[1] {
					break
				}
				if slt.Limit == nil && j > 0 {
					resstring += ","
				} else if slt.Limit != nil && j > slt.Limit[0] {
					resstring += ","
				}
				resstring += fmt.Sprintf("\"%s\"", data[i])
			}
			resstring += "]"
		}
		resstring += "}"
		fmt.Println(resstring)
		return []byte(resstring), nil
	} else { //two or more tables query
		cartesiankey, err := slt.cartesianProduct()
		if err != nil {
			return nil, err
		}
		fmt.Println(cartesiankey)
		var ids []string
		if len(slt.Where) == 0 { //no condition
			//get all ids
			ids, err = redigo.Strings(conn.Do("HKEYS", cartesiankey))
			if err != nil {
				fmt.Println("err:", err)
				return nil, err
			}
		} else { //single or more condition
			esStack := new_stack() //表达式栈
			snStack := new_stack() //关系符栈
			snStack.PUSH("#")      //start and end with #
			tmpWhere := append(slt.Where, "#")
			i := 0
			val := tmpWhere[i]
			for {
				if val == "#" && snStack.GetPOP() == "#" {
					break
				}
				fmt.Println("val:", val)
				if inarray(Opertion, val) == false && inarray(CompareSign, val) == false { //not opertion or relationship sign
					if inarray(CompareSign, snStack.GetPOP()) == false { //no relationship sign at top snStack
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
							fmt.Println("getDataIds:", alias, tmpids, err)
							if err != nil {
								return nil, err
							}
							//get condition sn
							cdtSn, err := getNextConditionSn()
							if err != nil {
								return nil, err
							}

							key := fmt.Sprintf(REDISQL_TMP_CONDITION, database, slt.Froms[alias], strconv.Itoa(cdtSn))
							for _, id := range tmpids {
								tmpkey := alias + "_" + id + "_"
								hkeys, err := redigo.Strings(conn.Do("HKEYS", cartesiankey))
								if err != nil {
									return nil, err
								}
								for _, hk := range hkeys {
									if strings.Contains(hk, tmpkey) {
										_, err = conn.Do("SADD", key, hk)
										if err != nil {
											return nil, err
										}
									}
								}
							}
							esStack.PUSH(key)
						} else {
							key, err := slt.getEsDataIds(cartesiankey, left, opt, right)
							if err != nil {
								return nil, err
							}
							esStack.PUSH(key)
						}
					}
					i++
				} else if inarray(CompareSign, val) == true {
					snStack.PUSH(val)
					i++
				} else if inarray(Opertion, val) == true {
					res := Compare(snStack.GetPOP(), val)
					fmt.Println(snStack.GetPOP(), val, "res=", res)
					switch res {
					case REDISQL_PRIORITY_LESS:
						snStack.PUSH(val)
						i++
						break
					case REDISQL_PRIORITY_EQUAL:
						snStack.POP()
						i++
						break
					case REDISQL_PRIORITY_GREATER:
						right := esStack.POP()
						left := esStack.POP()
						opt := snStack.POP()
						fmt.Printf("%s %s %s\n", left, opt, right)
						reskey, err := singleMerge(left, opt, right)
						if err != nil {
							return nil, err
						}
						esStack.PUSH(reskey)
						break
					case REDISQL_PRIORITY_ERROR:
						return nil, errors.New("wrong ralationship.")
					}
				}
				val = tmpWhere[i]
			}
			mykey := esStack.POP()
			ids, err = redigo.Strings(conn.Do("SMEMBERS", mykey))
			if err != nil {
				return nil, err
			}

		}
		fmt.Println(ids)
		var datas [][]string
		for _, id := range ids {
			tmpids := strings.Split(id, "_")
			idmap := make(map[string]string)
			for i := 1; i < len(tmpids); i = i + 2 {
				idmap[tmpids[i-1]] = tmpids[i]
			}
			var dataid []string
			for _, f := range slt.Fields {
				key := fmt.Sprintf(REDISQL_DATAS, database, slt.Froms[f.TableAlias], idmap[f.TableAlias])
				dataf, err := redigo.String(conn.Do("HGET", key, f.Name))
				if err != nil {
					return nil, err
				}
				dataid = append(dataid, dataf)
			}
			datas = append(datas, dataid)
		}

		var resstring string
		resstring += "{"
		for i, f := range slt.Fields {
			if i > 0 {
				resstring += ","
			}
			resstring += fmt.Sprintf("\"%s\":[", f.Alias)
			for j, data := range datas {
				if slt.Top > 0 && j >= slt.Top {
					break
				}
				if slt.Limit != nil && j < slt.Limit[0] {
					continue
				}
				if slt.Limit != nil && j > slt.Limit[1] {
					break
				}
				if slt.Limit == nil && j > 0 {
					resstring += ","
				} else if slt.Limit != nil && j > slt.Limit[0] {
					resstring += ","
				}
				resstring += fmt.Sprintf("\"%s\"", data[i])
			}
			resstring += "]"
		}
		resstring += "}"
		fmt.Println(resstring)
		return []byte(resstring), nil
	}

}

func clearTmp() {
	conn := getConn()
	defer conn.Close()
	keys, err := redigo.Strings(conn.Do("keys", "*.tmp.condition.*"))
	if err != nil {
		return
	}
	for _, v := range keys {
		_, err := conn.Do("DEL", v)
		if err != nil {
			return
		}
	}
}

func singleMerge(left, opt, right string) (string, error) {
	conn := getConn()
	defer conn.Close()
	cdtSn, err := getNextConditionSn()
	if err != nil {
		return "", err
	}
	tmp := strings.Split(left, ".tmp.condition.")
	key := tmp[0] + ".tmp.condition." + strconv.Itoa(cdtSn)
	fmt.Println(key)
	switch opt {
	case "and":
		_, err = conn.Do("SINTERSTORE", key, left, right)
		if err != nil {
			return "", err
		}
		break
	case "or":
		_, err = conn.Do("SUNIONSTORE", key, left, right)
		if err != nil {
			return "", err
		}
		break
	default:
		return "", errors.New("relationship sign wrong.")
	}
	return key, nil
}

//get Cartesian Product
func (slt *Select) cartesianProduct() (string, error) {
	conn := getConn()
	defer conn.Close()
	var alias []string
	ids := make(map[string][]string)
	var err error
	for k, v := range slt.Froms {
		key := fmt.Sprintf(REDISQL_INDEX_DATAS, database, v, "id")
		ids[k], err = redigo.Strings(conn.Do("ZRANGE", key, 0, -1))
		if err != nil {
			return "", err
		}
		alias = append(alias, k)
	}
	fmt.Println("ids:", ids)
	hkeys := make([]string, 0, 5)
	for i, a := range alias {
		fmt.Println("hkeys:", hkeys, "   i:", i, "   a:", a)
		if len(ids[a]) <= 0 { //some table is null, return an unexist key
			return "", errors.New("no result.")
		}
		if i == 0 {
			for _, id := range ids[a] {
				tmphkey := a + "_" + id + "_"
				hkeys = append(hkeys, tmphkey)
			}
		} else {
			tmphkeys := hkeys
			hkeys = make([]string, 0, 5)
			for _, id := range ids[a] {
				tmphkey := a + "_" + id + "_"
				for _, k := range tmphkeys {
					tmphk := k + tmphkey
					fmt.Println("k:", k, "   tmphkey:", tmphk)
					hkeys = append(hkeys, tmphk)
				}
			}
		}
	}
	cdtSn, err := getNextConditionSn()
	if err != nil {
		return "", nil
	}
	key := fmt.Sprintf(REDISQL_TMP_CARTESIAN, database, strconv.Itoa(cdtSn))
	for _, v := range hkeys {
		_, err = conn.Do("HSET", key, v, 0)
		if err != nil {
			return "", err
		}
	}
	return key, nil
}
