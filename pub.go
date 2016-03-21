package redisql

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
)

func getConn() redigo.Conn {
	return DB.pool.Get()
}

//database operation
func existsDatabase(dbname string) bool {
	fmt.Println("exists database %s start...", dbname)

	conn := getConn()
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

	conn := getConn()
	defer conn.Close()

	dbs, err := redigo.Strings(conn.Do("HKEYS", REDISQL_DATABASES))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return dbs
}

func getTableNumber(dbname string) (int, error) {
	fmt.Println("get " + dbname + " table number start...")
	conn := getConn()
	defer conn.Close()

	num, err := redigo.Int(conn.Do("HGET", REDISQL_DATABASES, dbname))
	if err != nil {
		return 0, err
	}

	return num, nil
}

//table opertion
func existsTable(dbname, tablename string) bool {
	fmt.Println("exists table %s.%s start...", dbname, tablename)

	conn := getConn()
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDISQL_TABLES, dbname), tablename))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	return exists
}

func getTables(dbname string) []string {
	fmt.Println("get %s tables start...", dbname)

	conn := getConn()
	defer conn.Close()

	tables, err := redigo.Strings(conn.Do("HKEYS", fmt.Sprintf(REDISQL_TABLES, dbname)))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return tables
}

func getNextId(dbname, tablename string) (int, error) {
	fmt.Println("get %s %s last id start...", dbname, tablename)

	conn := getConn()
	defer conn.Close()

	tmpid, err := redigo.Int(conn.Do("HGET", fmt.Sprintf(REDISQL_TABLES, dbname), tablename))
	if err != nil {
		return 0, err
	}

	return tmpid + 1, nil
}

//field opertion
func existsField(dbname, tablename, fieldname string) bool {
	fmt.Println("exists field:%s start...", fieldname)

	conn := getConn()
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDISQL_FIELDS, dbname, tablename), fieldname))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}

	return exists
}

//index opertion
func existsIndex(dbname, tablename, indexname string) bool {
	fmt.Println("exists %s %s index %s start...", dbname, tablename, indexname)

	conn := getConn()
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDISQL_INDEXS, dbname, tablename), indexname))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}

	return exists
}

func getIndexs(dbname, tablename string) map[string][]string {
	fmt.Println("get %s %s indexs start...", dbname, tablename)
	indexs := make(map[string][]string)

	conn := getConn()
	defer conn.Close()

	indexnames, err := redigo.Strings(conn.Do("HKEYS", fmt.Sprintf(REDISQL_INDEXS, dbname, tablename)))
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}

	for _, ix := range indexnames {
		fieldnames, err := redigo.Strings(conn.Do("HGET", fmt.Sprintf(REDISQL_INDEXS, dbname, tablename), ix))
		if err != nil {
			fmt.Errorf(err.Error())
			return nil
		}
		indexs[ix] = fieldnames
	}

	return indexs
}
