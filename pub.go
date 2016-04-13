package redisql

import (
	"container/list"
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"strings"
)

func getConn() redigo.Conn {
	conn := DB.pool.Get()
	conn.Do("SELECT", redisdb)
	return conn
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

func GetDatabases() []string {
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

func getTableNumber() (int, error) {
	fmt.Println("get " + database + " table number start...")
	conn := getConn()
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

	conn := getConn()
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDISQL_TABLES, database), tablename))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	return exists
}

func GetTables() []string {
	fmt.Println("get %s tables start...", database)

	conn := getConn()
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

	conn := getConn()
	defer conn.Close()

	return redigo.Int(conn.Do("HGET", fmt.Sprintf(REDISQL_COUNT, database), tablename))
}

func getNextId(tablename string) (int, error) {
	fmt.Println("get %s %s last id start...", database, tablename)

	conn := getConn()
	defer conn.Close()

	tmpid, err := redigo.Int(conn.Do("HGET", fmt.Sprintf(REDISQL_TABLES, database), tablename))
	if err != nil {
		return 0, err
	}

	return tmpid + 1, nil
}

func getNextConditionSn() (int, error) {
	conn := getConn()
	defer conn.Close()

	cdtSn, err := redigo.Int(conn.Do("GET", fmt.Sprintf(REDISQL_CONDITION_SN, database)))
	if err != nil {
		return 0, err
	}
	cdtSn = cdtSn + 1
	_, err = conn.Do("INCRBY", fmt.Sprintf(REDISQL_CONDITION_SN, database), 1)
	if err != nil {
		return 0, err
	}
	return cdtSn, nil
}

//field opertion
func existsField(tablename, fieldname string) bool {
	fmt.Println("exists field:%s start...", fieldname)

	conn := getConn()
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDISQL_FIELDS, database, tablename), fieldname))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}

	return exists
}

func getFieldType(tablename, fieldname string) (string, error) {
	fmt.Printf("get field %s.%s's type start...\n", tablename, fieldname)

	conn := getConn()
	defer conn.Close()

	return redigo.String(conn.Do("HGET", fmt.Sprintf(REDISQL_FIELDS, database, tablename), fieldname))
}

//index opertion
func existsIndex(tablename, indexname string) bool {
	fmt.Println("exists %s %s index %s start...", database, tablename, indexname)

	conn := getConn()
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

	conn := getConn()
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

//return res when res int array1 and array2
func inter(array1, array2 []string) []string {
	var arrRes []string
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if v1 == v2 {
				arrRes = append(arrRes, v1)
				continue
			}
		}
	}
	return arrRes
}

//return res when res in array1 or array2
func union(array1, array2 []string) []string {
	var arrRes []string
	arrRes = append(arrRes, array1...)
	arrRes = append(arrRes, array2...)
	return arrRes
}

//return res when res in array1 but not in array2
func outer(array1, array2 []string) []string {
	var arrRes []string
	for _, v1 := range array1 {
		if inarray(array2, v1) == false {
			arrRes = append(arrRes, v1)
		}
	}
	return arrRes
}

//judge is sub in array
func inarray(array []string, sub string) bool {
	for _, v := range array {
		if sub == v {
			return true
		}
	}
	return false
}

//create stack
type stack struct {
	s *list.List
}

func new_stack() *stack {
	return &stack{
		s: list.New(),
	}
}
func (s *stack) PUSH(value string) {
	s.s.PushBack(value)
}

func (s *stack) POP() string {
	res := s.s.Back()
	s.s.Remove(res)
	return res.Value.(string)
}

func (s *stack) GetPOP() string {
	res := s.s.Back()
	return res.Value.(string)
}

//
func Compare(oldsign, newsign string) int {
	switch oldsign {
	case "#":
		switch newsign {
		case "(":
			return REDISQL_PRIORITY_LESS
		case "and":
			return REDISQL_PRIORITY_LESS
		case "or":
			return REDISQL_PRIORITY_LESS
		case "#":
			return REDISQL_PRIORITY_GREATER
		default:
			return REDISQL_PRIORITY_ERROR
		}
	case "(":
		switch newsign {
		case "(":
			return REDISQL_PRIORITY_LESS
		case ")":
			return REDISQL_PRIORITY_EQUAL
		case "or":
			return REDISQL_PRIORITY_LESS
		case "and":
			return REDISQL_PRIORITY_LESS
		default:
			return REDISQL_PRIORITY_ERROR
		}
		break
	case ")":
		switch newsign {
		case "or":
			return REDISQL_PRIORITY_GREATER
		case "and":
			return REDISQL_PRIORITY_GREATER
		default:
			return REDISQL_PRIORITY_ERROR
		}
		break
	case "and":
		switch newsign {
		case "or":
			return REDISQL_PRIORITY_GREATER
		case "and":
			return REDISQL_PRIORITY_GREATER
		case "(":
			return REDISQL_PRIORITY_LESS
		case ")":
			return REDISQL_PRIORITY_GREATER
		case "#":
			return REDISQL_PRIORITY_GREATER
		default:
			return REDISQL_PRIORITY_ERROR
		}
		break
	case "or":
		switch newsign {
		case "or":
			return REDISQL_PRIORITY_GREATER
		case "and":
			return REDISQL_PRIORITY_GREATER
		case "(":
			return REDISQL_PRIORITY_LESS
		case ")":
			return REDISQL_PRIORITY_GREATER
		case "#":
			return REDISQL_PRIORITY_GREATER
		default:
			return REDISQL_PRIORITY_ERROR
		}
		break
	default:
		return REDISQL_PRIORITY_ERROR
	}
	return REDISQL_PRIORITY_ERROR
}
