package redisql

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
		_, err = conn.Do("SADD", fmt.Sprintf(REDISQL_INDEX_DATAS, database, ist.Into, indexdata), tmpid)
		if err != nil {
			conn.Do("DISCARD")
			return err
		}
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		conn.Do("DISCARD")
		return err
	}
	return nil
}
