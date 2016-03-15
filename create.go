package redisql

import (
	"errors"
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"os"
	"strings"
)

//how to use?
//eg: TABLE("user").FIELDS("name, age, city").TYPES("xzj",26,"sh").CREATE()

type Table struct {
	Database int
	Name     string
	Fields   []string
	Types    []string
}

func TABLE(tablename string) *Table {
	if tablename == "" {
		fmt.Errorf("tablename can not be null.")
		os.Exit(1)
	}
	return &Table{
		Database: selectdb,
		Name:     tablename,
	}
}

func (tab *Table) FIELDS(fields ...string) *Table {
	if tab.Name == "" {
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

	return &Table{
		Database: tab.Database,
		Name:     tab.Name,
		Fields:   tmpFields,
	}
}

func (tab *Table) TYPES(types ...string) *Table {
	if len(tab.Fields) <= 0 {
		fmt.Errorf("fields is null, please call func 'FIELDS'.")
		os.Exit(1)
	}
	if len(tab.Fields) != len(types) {
		fmt.Errorf("Field and types are not correspondence, please check.")
		os.Exit(1)
	}
	return &Table{
		Database: tab.Database,
		Name:     tab.Name,
		Fields:   tab.Fields,
		Types:    types,
	}
}

func (tab *Table) CREATE() error {
	fmt.Println("create start...data:", *tab)

	conn := DB.pool.Get()
	defer conn.Close()

	//change db
	conn.Do("SELECT", tab.Database)

	//judge table is exists?
	exists, err := redigo.Bool(conn.Do("EXISTS", fmt.Sprintf(SYS_FIELDS, tab.Name)))
	if err != nil {
		return err
	}

	if exists == true {
		return errors.New(fmt.Sprintf("table %s is exist.", tab.Name))
	}

	//add table info
	var params []interface{}
	params = append(params, fmt.Sprintf(SYS_FIELDS, tab.Name))
	for i := 0; i < len(tab.Fields); i++ {
		params = append(params, tab.Fields[i])
		params = append(params, tab.Types[i])
	}

	_, err = conn.Do("HMSET", params...)
	if err != nil {
		return err
	}

	return nil
}

func (tab *Table) INDEX() error {
	fmt.Println("index start...date:", tab)

	conn := DB.pool.Get()
	defer conn.Close()

	conn.Do("SELECT", tab.Database)
	//judge table is exists?
	exists, err := redigo.Bool(conn.Do("EXISTS", fmt.Sprintf(SYS_FIELDS, tab.Name)))
	if err != nil {
		return err
	}

	if exists == false {
		return errors.New(fmt.Sprintf("table %s is not exist.", tab.Name))
	}

	indexname := "index"
	for _, f := range tab.Fields {
		exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(SYS_FIELDS, tab.Name), f))
		if err != nil {
			return err
		}
		if exists == false {
			return errors.New(fmt.Sprintf("no field %s in table %s.", f, tab.Name))
		}
		indexname += f
	}
	//judge index is exists?
	exists, err = redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(SYS_INDEXS, tab.Name), indexname))
	if err != nil {
		return err
	}

	if exists == true {
		return errors.New(fmt.Sprintf("index %s is exist.", tab.Name))
	}
	//save index
	_, err = conn.Do("HSET", fmt.Sprintf(SYS_INDEXS, tab.Name), indexname, tab.Fields)
	if err != nil {
		return err
	}

	//add index
	//get all data
	fmt.Println(fmt.Sprintf(USER_TABLE, tab.Name, "*"))
	rows, err := redigo.Strings(conn.Do("KEYS", fmt.Sprintf(USER_TABLE, tab.Name, "*")))
	if err != nil {
		return err
	}
	fmt.Println("row:", rows)
	for _, r := range rows {
		fmt.Println(r)
		var indexField string
		for i := 0; i < len(tab.Fields); i++ {
			value, err := redigo.String(conn.Do("HGET", r, tab.Fields[i]))
			if err != nil {
				return err
			}
			indexField += fmt.Sprintf(".%s:%s", tab.Fields[i], value)
		}
		fmt.Println(indexField)
		ids := strings.Split(r, ":")
		_, err := conn.Do("SADD", fmt.Sprintf(USER_INDEX, tab.Name, indexname, indexField), ids[len(ids)-1])
		if err != nil {
			return err
		}
	}
	return nil
}
