package redisql

import (
	"container/list"
	"errors"
	"fmt"
	"strings"
	"time"

	redigo "github.com/garyburd/redigo/redis"
)

func getConn() redigo.Conn {
	conn := DB.pool.Get()
	conn.Do("SELECT", redisdb)
	return conn
}

// database operation
func existsDatabase(dbname string) bool {
	conn := getConn()
	defer conn.Close()
	exists, err := redigo.Bool(conn.Do("HEXISTS", REDISQL_DATABASES, dbname))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}

	return exists
}

func GetDatabases() ([]string, float64, error) {
	conn := getConn()
	defer conn.Close()
	start := time.Now()
	dbs, err := redigo.Strings(conn.Do("HKEYS", REDISQL_DATABASES))
	end := time.Now()
	usetime := end.Sub(start).Seconds()
	if err != nil {
		return nil, usetime, errors.New("ERROR 1045: " + err.Error())
	}
	return dbs, usetime, nil
}

func GetTables() ([]string, float64, error) {
	if database == "" {
		return nil, 0, errors.New("ERROR 1046: No database selected")
	}
	conn := getConn()
	defer conn.Close()
	start := time.Now()
	tables, err := redigo.Strings(conn.Do("HKEYS", fmt.Sprintf(REDISQL_TABLES, database)))
	end := time.Now()
	usetime := end.Sub(start).Seconds()
	if err != nil {
		return nil, usetime, errors.New("ERROR 1045: " + err.Error())
	}
	return tables, usetime, nil
}

func GetTableInfo(tablename string) ([][]string, float64, error) {
	if database == "" {
		return nil, 0, errors.New("ERROR 1046: No database selected")
	}
	if existsTable(tablename) == false {
		return nil, 0, errors.New(fmt.Sprintf("ERROR 1146: Table '%s.%s' doesn't exist", database, tablename))
	}

	start := time.Now()
	conn := getConn()
	defer conn.Close()

	fieldstypes, err := redigo.Strings(conn.Do("HGETALL", fmt.Sprintf(REDISQL_FIELDS, database, tablename)))
	end := time.Now()
	usetime := end.Sub(start).Seconds()
	if err != nil {
		return nil, usetime, errors.New("ERROR 1045: " + err.Error())
	}

	var res [][]string
	for i := 0; i < len(fieldstypes); i += 2 {
		var tmpres []string
		tmpres = append(tmpres, fieldstypes[i])
		tmpres = append(tmpres, fieldstypes[i+1])
		res = append(res, tmpres)
	}
	end = time.Now()
	usetime = end.Sub(start).Seconds()
	return res, usetime, nil
}

func getTableNumber() (int, error) {
	conn := getConn()
	defer conn.Close()

	num, err := redigo.Int(conn.Do("HGET", REDISQL_DATABASES, database))
	if err != nil {
		return 0, err
	}

	return num, nil
}

// table opertion
func existsTable(tablename string) bool {
	conn := getConn()
	defer conn.Close()

	exists, err := redigo.Bool(conn.Do("HEXISTS", fmt.Sprintf(REDISQL_TABLES, database), tablename))
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	return exists
}

func getCount(tablename string) (int, error) {
	conn := getConn()
	defer conn.Close()

	return redigo.Int(conn.Do("HGET", fmt.Sprintf(REDISQL_COUNT, database), tablename))
}

func getNextId(tablename string) (int, error) {
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

// field opertion
func existsField(tablename, fieldname string) bool {
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
	conn := getConn()
	defer conn.Close()

	return redigo.String(conn.Do("HGET", fmt.Sprintf(REDISQL_FIELDS, database, tablename), fieldname))
}

// index opertion
func existsIndex(tablename, indexname string) bool {
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

// return res when res int array1 and array2
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

// return res when res in array1 or array2
func union(array1, array2 []string) []string {
	var arrRes []string
	arrRes = append(arrRes, array1...)
	arrRes = append(arrRes, array2...)
	return arrRes
}

// return res when res in array1 but not in array2
func outer(array1, array2 []string) []string {
	var arrRes []string
	for _, v1 := range array1 {
		if inarray(array2, v1) == false {
			arrRes = append(arrRes, v1)
		}
	}
	return arrRes
}

// judge is sub in array
func inarray(array []string, sub string) bool {
	for _, v := range array {
		if sub == v {
			return true
		}
	}
	return false
}

// create stack
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

func Compare(oldsign, newsign string) int {
	switch oldsign {
	case "#":
		switch newsign {
		case "(":
			return REDISQL_PRIORITY_LESS
		case "AND":
			return REDISQL_PRIORITY_LESS
		case "OR":
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
		case "OR":
			return REDISQL_PRIORITY_LESS
		case "AND":
			return REDISQL_PRIORITY_LESS
		default:
			return REDISQL_PRIORITY_ERROR
		}
		break
	case ")":
		switch newsign {
		case "OR":
			return REDISQL_PRIORITY_GREATER
		case "AND":
			return REDISQL_PRIORITY_GREATER
		default:
			return REDISQL_PRIORITY_ERROR
		}
		break
	case "AND":
		switch newsign {
		case "OR":
			return REDISQL_PRIORITY_GREATER
		case "AND":
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
	case "OR":
		switch newsign {
		case "OR":
			return REDISQL_PRIORITY_GREATER
		case "AND":
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
