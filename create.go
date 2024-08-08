package redisql

import (
	"errors"
	"fmt"
	"strings"
	"time"

	redigo "github.com/garyburd/redigo/redis"
)

//how to use?
//eg: 	create table user
//		TABLE("user").FIELDS("name, age, city").TYPES("xzj",26,"sh").CREATE()

//eg:	create index on user(name)
//		TABLE("user").FIELDS("name").INDEX()

func CreateDatabase(dbname string) error {
	db := strings.ToLower(strings.Trim(dbname, " "))
	if len(db) <= 0 {
		return errors.New("database name can not be null.")
	}

	if existsDatabase(db) == true {
		return errors.New(fmt.Sprintf("database:%s exists.\n", db))
	}

	conn := getConn()
	defer conn.Close()

	_, err := conn.Do("HSET", REDISQL_DATABASES, db, 0)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", fmt.Sprintf(REDISQL_CONDITION_SN, db), 0)
	if err != nil {
		return err
	}
	return nil
}

type Table struct {
	Name   string
	Fields []string
	Types  []string
}

func TABLE(tablename string) *Table {
	if len(database) <= 0 {
		panic("you have not choose database, please call func 'ChangeDatabase'.")
	}
	if len(strings.Trim(tablename, "")) <= 0 {
		panic("tablename can not be null.")
	}
	return &Table{
		Name: strings.ToLower(strings.Trim(tablename, "")),
	}
}

func (tab *Table) FIELDS(fields ...string) *Table {
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

	return &Table{
		Name:   tab.Name,
		Fields: tmpFields,
	}
}

func (tab *Table) TYPES(types ...string) *Table {
	if len(tab.Fields) <= 0 {
		panic("fields is null, please call func 'FIELDS'.")
	}
	var tmpTypes []string
	for _, t := range types {
		tmpt := strings.Split(t, ",")
		for _, tt := range tmpt {
			if strings.ToLower(strings.Trim(tt, " ")) != REDISQL_TYPE_NUMBER &&
				strings.ToLower(strings.Trim(tt, " ")) != REDISQL_TYPE_STRING &&
				strings.ToLower(strings.Trim(tt, " ")) != REDISQL_TYPE_DATE {
				panic("Redisql not such type like " + strings.ToLower(strings.Trim(tt, " ")))
			}
			tmpTypes = append(tmpTypes, strings.ToLower(strings.Trim(tt, " ")))
		}
	}
	if len(tab.Fields) != len(tmpTypes) {
		panic("Field and types are not correspondence, please check.")
	}
	return &Table{
		Name:   tab.Name,
		Fields: tab.Fields,
		Types:  tmpTypes,
	}
}

func (tab *Table) CREATE() error {
	conn := getConn()
	defer conn.Close()

	//judge table is exists?
	exists := existsTable(tab.Name)
	if exists == true {
		return errors.New(fmt.Sprintf("table %s is exist.", tab.Name))
	}

	//add table info
	var params []interface{}
	params = append(params, fmt.Sprintf(REDISQL_FIELDS, database, tab.Name))
	//add field 'id' default.
	params = append(params, "id")
	params = append(params, REDISQL_TYPE_NUMBER)
	for i := 0; i < len(tab.Fields); i++ {
		if tab.Fields[i] == "id" {
			continue
		}
		params = append(params, tab.Fields[i])
		params = append(params, tab.Types[i])
	}

	_, err := conn.Do("MULTI")
	if err != nil {
		return err
	}
	_, err = conn.Do("HMSET", params...)
	if err != nil {
		return err
	}

	_, err = conn.Do("HSET", fmt.Sprintf(REDISQL_TABLES, database), tab.Name, 0)
	if err != nil {
		conn.Do("DISCARD")
		return err
	}

	_, err = conn.Do("HSET", fmt.Sprintf(REDISQL_COUNT, database), tab.Name, 0)
	if err != nil {
		conn.Do("DISCARD")
		return err
	}

	//tablenumber +1
	_, err = conn.Do("HINCRBY", REDISQL_DATABASES, database, 1)
	if err != nil {
		conn.Do("DISCARD")
		return err
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		conn.Do("DISCARD")
		return err
	}
	//default create index on id
	err = TABLE(tab.Name).FIELDS("id").INDEX()
	if err != nil {
		return err
	}
	return nil
}

func (tab *Table) INDEX() error {
	conn := getConn()
	defer conn.Close()

	//judge table is exists?
	exists := existsTable(tab.Name)
	if exists == false {
		return errors.New(fmt.Sprintf("table %s is not exist.", tab.Name))
	}

	indexname := "index"
	for _, f := range tab.Fields {
		if existsField(tab.Name, f) == false {
			return errors.New(fmt.Sprintf("no field %s in table %s.", f, tab.Name))
		}
		indexname += "_"
		indexname += f
	}
	//judge index is exists?
	if existsIndex(tab.Name, indexname) == true {
		return errors.New(fmt.Sprintf("index %s is exist.", tab.Name))
	}
	//save index
	_, err := conn.Do("HSET", fmt.Sprintf(REDISQL_INDEXS, database, tab.Name), indexname, tab.Fields)
	if err != nil {
		return err
	}

	var fieldtype string
	if len(tab.Fields) == 1 {
		fieldtype, err = getFieldType(tab.Name, tab.Fields[0])
		if err != nil {
			return err
		}
	}

	//add index
	//get all data
	rows, err := redigo.Strings(conn.Do("KEYS", fmt.Sprintf(REDISQL_DATAS, database, tab.Name, "*")))
	if err != nil {
		return err
	}

	for _, r := range rows {

		var indexField string
		for i := 0; i < len(tab.Fields); i++ {
			value, err := redigo.String(conn.Do("HGET", r, tab.Fields[i]))
			if err != nil {
				return err
			}
			if i >= 1 {
				indexField += "."
			}
			indexField += fmt.Sprintf("%s.%s", tab.Fields[i], value)
		}

		ids := strings.Split(r, ".")
		if fieldtype == "" || fieldtype == REDISQL_TYPE_STRING {
			_, err := conn.Do("SADD", fmt.Sprintf(REDISQL_INDEX_DATAS, database, tab.Name, indexField), ids[len(ids)-1])
			if err != nil {
				return err
			}
		} else {
			kv := strings.Split(indexField, ".")
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

			_, err := conn.Do("ZADD", fmt.Sprintf(REDISQL_INDEX_DATAS, database, tab.Name, kv[0]), score, ids[len(ids)-1])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
