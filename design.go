package redisql

import (
	"errors"
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
)

func (r *Redisql) TABLE(table string) *Redisql {
	return &Redisql{
		Database: selectdb,
		Table:    table,
		Fields:   r.Fields,
		Types:    r.Types,
		Values:   r.Values,
	}
}

func (r *Redisql) FIELDS(fields ...string) *Redisql {
	return &Redisql{
		Database: selectdb,
		Table:    r.Table,
		Fields:   fields,
		Types:    r.Types,
		Values:   r.Values,
	}
}

func (r *Redisql) TYPES(types ...string) *Redisql {
	return &Redisql{
		Database: selectdb,
		Table:    r.Table,
		Fields:   r.Fields,
		Types:    types,
		Values:   r.Values,
	}
}

func (r *Redisql) INTO(table string) *Redisql {
	return r.TABLE(table)
}

func (r *Redisql) VALUES(values ...interface{}) *Redisql {
	return &Redisql{
		Database: selectdb,
		Table:    r.Table,
		Fields:   r.Fields,
		Types:    r.Types,
		Values:   values,
	}
}

func (r *Redisql) check(operate string) error {
	fmt.Println("check start...data:", operate)
	if r.Table == "" {
		return errors.New("not set table name, please call func 'INTO' or 'TABLE'.")
	}

	if len(r.Fields) <= 0 {
		return errors.New("not set fields, please call func 'FIELDS'.")
	}

	if len(r.Types) <= 0 && operate == SQL_CREATE {
		return errors.New("not set types, please call func 'TYPES'.")
	}

	if len(r.Values) <= 0 && operate == SQL_INSERT {
		return errors.New("not set values, please call func 'VALUES'.")
	}

	if len(r.Fields) != len(r.Types) && operate == SQL_CREATE {
		return errors.New("Field and type are not correspondence, please check.")
	}

	if len(r.Fields) != len(r.Values) && operate == SQL_INSERT {
		return errors.New("Field and value are not correspondence, please check.")
	}

	return nil
}

func (r *Redisql) UNIQUE() error {
	fmt.Println("unique start...data:", *r)
	err := r.check(SQL_UNIQUE)
	if err != nil {
		return err
	}

	conn := DB.pool.Get()
	defer conn.Close()

	//change db
	conn.Do("SELECT", r.Database)

	var params []interface{}
	params = append(params, fmt.Sprintf(REDIS_KEY_UNIQUES, r.Table))
	for i := 0; i < len(r.Fields); i++ {
		params = append(params, r.Fields[i])
	}

	_, err = conn.Do("SADD", params...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Redisql) CREATE() error {
	fmt.Println("create start...data:", *r)
	//check Redisql
	err := r.check(SQL_CREATE)
	if err != nil {
		return err
	}

	conn := DB.pool.Get()
	defer conn.Close()

	//change db
	conn.Do("SELECT", r.Database)

	//judge table is exists?
	exists, err := redigo.Bool(conn.Do("EXISTS", fmt.Sprintf(REDIS_KEY_FIELDS, r.Table)))
	if err != nil {
		return err
	}

	if exists == true {
		return errors.New(fmt.Sprintf("table %s is exist.", r.Table))
	}

	//add table info
	var params []interface{}
	params = append(params, fmt.Sprintf(REDIS_KEY_FIELDS, r.Table))
	for i := 0; i < len(r.Fields); i++ {
		params = append(params, r.Fields[i])
		params = append(params, r.Types[i])
	}

	_, err = conn.Do("HMSET", params...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Redisql) INSERT() error {
	fmt.Println("insert start...data:", *r)

	//check Redisql
	err := r.check(SQL_INSERT)
	if err != nil {
		return err
	}

	conn := DB.pool.Get()
	defer conn.Close()

	//change db
	conn.Do("SELECT", r.Database)
	exists, err := redigo.Bool(conn.Do("EXISTS", fmt.Sprintf(REDIS_KEY_FIELDS, r.Table)))
	if err != nil {
		return err
	}

	if exists == false {
		return errors.New(fmt.Sprintf("table %s is not exist.", r.Table))
	}

	//get table max id
	var tmpid int
	row, err := conn.Do("HGET", REDIS_KEY_TABLES_MAX_ID, r.Table)
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
	params = append(params, fmt.Sprintf("%s.%d", r.Table, tmpid))
	for i := 0; i < len(r.Fields); i++ {
		exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDIS_KEY_FIELDS, r.Table), r.Fields[i]))
		if err != nil {
			return err
		}
		if exists == false {
			return errors.New(fmt.Sprintf("no field %s in table %s.", r.Fields[i], r.Table))
		}
		params = append(params, r.Fields[i])
		params = append(params, r.Values[i])
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
	_, err = conn.Do("HSET", REDIS_KEY_TABLES_MAX_ID, r.Table, tmpid)
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
