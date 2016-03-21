package redisql

import (
	"errors"
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"os"
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
	return &Insert{
		Into: strings.ToLower(tablename),
	}
}

func (ist *Insert) FIELDS(fields ...string) *Insert {
	if ist.Into == "" {
		fmt.Errorf("table name is null, please call func 'INTO'.")
		os.Exit(1)
	}

	if len(fields) <= 0 {
		fmt.Errorf("can not call this func without fields.")
		os.Exit(1)
	}

	var tmpFields []string
	for _, f := range fields {
		tmpf := strings.Split(f, ",")
		for _, ff := range tmpf {
			tmpFields = append(tmpFields, strings.ToLower(strings.Trim(ff, " ")))
		}
	}

	return &Insert{
		Into:   ist.Into,
		Fields: tmpFields,
	}
}

func (ist *Insert) VALUES(values ...interface{}) *Insert {
	if len(ist.Fields) <= 0 {
		fmt.Errorf("fields is null, please call func 'FIELDS'.")
		os.Exit(1)
	}
	if len(ist.Fields) != len(values) {
		fmt.Errorf("Field and value are not correspondence, please check.")
		os.Exit(1)
	}
	return &Insert{
		Into:   ist.Into,
		Fields: ist.Fields,
		Values: values,
	}
}

func (ist *Insert) INSERT() error {
	fmt.Println("insert start...data:", *ist)

	conn := getConn(redisdb)
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("EXISTS", fmt.Sprintf(SYS_FIELDS, ist.Into)))
	if err != nil {
		return err
	}

	if exists == false {
		return errors.New(fmt.Sprintf("table %s is not exist.", ist.Into))
	}

	//get table max id
	var tmpid int
	row, err := conn.Do("HGET", SYS_TABLES_MAX_ID, ist.Into)
	if err != nil {
		return err
	} else {
		if row == nil {
			tmpid = 0
		} else {
			tmpid, err = redigo.Int(row, nil)
			if err != nil {
				return err
			}
		}
	}

	tmpid = tmpid + 1

	//get data Info
	var params []interface{}
	params = append(params, fmt.Sprintf(USER_TABLE, ist.Into, strconv.Itoa(tmpid)))
	for i := 0; i < len(ist.Fields); i++ {
		exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(SYS_FIELDS, ist.Into), ist.Fields[i]))
		if err != nil {
			return err
		}
		if exists == false {
			return errors.New(fmt.Sprintf("no field %s in table %s.", ist.Fields[i], ist.Into))
		}
		params = append(params, ist.Fields[i])
		params = append(params, ist.Values[i])
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
		fmt.Println("HMSET", err)
		conn.Do("DISCARD")
		return err
	}

	//update max id
	_, err = conn.Do("HSET", SYS_TABLES_MAX_ID, ist.Into, tmpid)
	if err != nil {
		fmt.Println("HSET", err)
		conn.Do("DISCARD")
		return err
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		fmt.Println("EXEC", err)
		conn.Do("DISCARD")
		return err
	}
	return nil
}
