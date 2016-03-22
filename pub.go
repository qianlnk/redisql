package redisql

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"strings"
)

func getConn(db int) redigo.Conn {
	conn := DB.pool.Get()
	conn.Do("SELECT", db)
	return conn
}

//database operation
func existsDatabase(dbname string) bool {
	fmt.Println("exists database %s start...", dbname)

	conn := getConn(redisdb)
	defer conn.Close()
	exists, err := redigo.Bool(conn.Do("HEXISTS", REDISQL_DATABASES, dbname))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}

	return exists
}

func getDatabases() []string {
	fmt.Println("get databaases start...")

	conn := getConn(redisdb)
	defer conn.Close()

	dbs, err := redigo.Strings(conn.Do("HKEYS", REDISQL_DATABASES))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return dbs
}

func getTableNumber() (int, error) {
	fmt.Println("get " + database + " table number start...")
	conn := getConn(redisdb)
	defer conn.Close()

	num, err := redigo.Int(conn.Do("HGET", REDISQL_DATABASES, database))
	if err != nil {
		return 0, err
	}

	return num, nil
}

//table opertion
func existsTable(tablename string) bool {
	fmt.Println("exists table %s.%s start...", database, tablename)

	conn := getConn(redisdb)
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDISQL_TABLES, database), tablename))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	return exists
}

func getTables() []string {
	fmt.Println("get %s tables start...", database)

	conn := getConn(redisdb)
	defer conn.Close()

	tables, err := redigo.Strings(conn.Do("HKEYS", fmt.Sprintf(REDISQL_TABLES, database)))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return tables
}

func getCount(tablename string) (int, error) {
	fmt.Println("get " + tablename + " count start...")

	conn := getConn(redisdb)
	defer conn.Close()

	return redigo.Int(conn.Do("HGET", fmt.Sprintf(REDISQL_COUNT, database), tablename))
}

func getNextId(tablename string) (int, error) {
	fmt.Println("get %s %s last id start...", database, tablename)

	conn := getConn(redisdb)
	defer conn.Close()

	tmpid, err := redigo.Int(conn.Do("HGET", fmt.Sprintf(REDISQL_TABLES, database), tablename))
	if err != nil {
		return 0, err
	}

	return tmpid + 1, nil
}

//field opertion
func existsField(tablename, fieldname string) bool {
	fmt.Println("exists field:%s start...", fieldname)

	conn := getConn(redisdb)
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDISQL_FIELDS, database, tablename), fieldname))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}

	return exists
}

//index opertion
func existsIndex(tablename, indexname string) bool {
	fmt.Println("exists %s %s index %s start...", database, tablename, indexname)

	conn := getConn(redisdb)
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDISQL_INDEXS, database, tablename), indexname))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}

	return exists
}

func getIndexs(tablename string) (map[string][]string, error) {
	fmt.Println("get %s %s indexs start...", database, tablename)
	indexs := make(map[string][]string)

	conn := getConn(redisdb)
	defer conn.Close()

	indexnames, err := redigo.Strings(conn.Do("HKEYS", fmt.Sprintf(REDISQL_INDEXS, database, tablename)))
	if err != nil {
		return nil, err
	}

	for _, ix := range indexnames {
		fieldnames, err := redigo.String(conn.Do("HGET", fmt.Sprintf(REDISQL_INDEXS, database, tablename), ix))
		if err != nil {
			return nil, err
		}
		indexs[ix] = strings.Split(strings.Replace(strings.Replace(fieldnames, "[", "", -1), "]", "", -1), " ")
	}

	return indexs, nil
}
