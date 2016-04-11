package redisql

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//how to use?
//eg: INTO("user").FIELDS("name, age, city").VALUES("xzj",26,"sh").INSERT()

type Insert struct {
	Into   string
	Fields []string
	Values []interface{}
}

func INTO(tablename string) *Insert {
	fmt.Println("into start...")
	if len(database) <= 0 {
		panic("you have not choose database, please call func 'ChangeDatabase'.")
	}
	if len(strings.Trim(tablename, "")) <= 0 {
		panic("tablename can not be null.")
	}
	return &Insert{
		Into: strings.ToLower(strings.Trim(tablename, "")),
	}
}

func (ist *Insert) FIELDS(fields ...string) *Insert {
	var tmpFields []string
	for _, f := range fields {
		tmpf := strings.Split(f, ",")
		for _, ff := range tmpf {
			tmpFields = append(tmpFields, strings.ToLower(strings.Trim(ff, " ")))
		}
	}

	if len(tmpFields) <= 0 {
		panic("can not call this func without fields.")
	}

	return &Insert{
		Into:   ist.Into,
		Fields: tmpFields,
	}
}

func (ist *Insert) VALUES(values ...interface{}) *Insert {
	if len(ist.Fields) <= 0 {
		panic("fields is null, please call func 'FIELDS'.")
	}
	if len(ist.Fields) != len(values) {
		panic("Field and value are not correspondence, please check.")
	}
	return &Insert{
		Into:   ist.Into,
		Fields: ist.Fields,
		Values: values,
	}
}

func (ist *Insert) INSERT() error {
	fmt.Println("insert start...data:", *ist)

	conn := getConn()
	defer conn.Close()

	if existsTable(ist.Into) == false {
		return errors.New(fmt.Sprintf("table %s is not exist.", ist.Into))
	}

	//get table max id
	tmpid, err := getNextId(ist.Into)
	if err != nil {
		return err
	}

	//get data Info
	var params []interface{}
	params = append(params, fmt.Sprintf(REDISQL_DATAS, database, ist.Into, strconv.Itoa(tmpid)))
	params = append(params, "id")
	params = append(params, tmpid)
	for i := 0; i < len(ist.Fields); i++ {
		if existsField(ist.Into, ist.Fields[i]) == false {
			return errors.New(fmt.Sprintf("no field %s in table %s.", ist.Fields[i], ist.Into))
		}
		params = append(params, ist.Fields[i])
		params = append(params, ist.Values[i])
	}

	//get table indexs
	indexs, err := getIndexs(ist.Into)
	if err != nil {
		return err
	}

	//begin work
	_, err = conn.Do("MULTI")
	if err != nil {
		fmt.Println("MULTI", err)
		return err
	}

	//insert new data
	_, err = conn.Do("HMSET", params...)
	if err != nil {
		conn.Do("DISCARD")
		return err
	}

	//update max id
	_, err = conn.Do("HSET", fmt.Sprintf(REDISQL_TABLES, database), ist.Into, tmpid)
	if err != nil {
		conn.Do("DISCARD")
		return err
	}

	//update table count
	_, err = conn.Do("HINCRBY", fmt.Sprintf(REDISQL_COUNT, database), ist.Into, 1)
	if err != nil {
		conn.Do("DISCARD")
		return err
	}

	//add new indexdata
	ist.Fields = append(ist.Fields, "id")
	ist.Values = append(ist.Values, tmpid)
	for _, v := range indexs {
		flag := 0
		var indexdata string
		for _, ixf := range v {
			for i, f := range ist.Fields {
				if ixf == f {
					if flag >= 1 {
						indexdata += "."
					}
					flag += 1
					indexdata += fmt.Sprintf("%s.%v", f, ist.Values[i])
				}
			}
		}
		var fieldtype string
		if len(v) == 1 {
			fieldtype, err = getFieldType(ist.Into, v[0])
			if err != nil {
				return err
			}
		}
		if fieldtype == "" || fieldtype == REDISQL_TYPE_STRING {
			_, err = conn.Do("SADD", fmt.Sprintf(REDISQL_INDEX_DATAS, database, ist.Into, indexdata), tmpid)
			if err != nil {
				conn.Do("DISCARD")
				return err
			}
		} else {
			kv := strings.Split(indexdata, ".")
			var score string
			if fieldtype == REDISQL_TYPE_DATE {
				t, err := time.Parse("2006-01-02 15:04:05", kv[1])
				if err != nil {
					return err
				}
				score = t.Format("20060102150405")
			} else {
				score = kv[1]
			}
			fmt.Printf("score = %s", score)
			_, err := conn.Do("ZADD", fmt.Sprintf(REDISQL_INDEX_DATAS, database, ist.Into, kv[0]), score, tmpid)
			if err != nil {
				return err
			}
		}
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		conn.Do("DISCARD")
		return err
	}
	return nil
}
