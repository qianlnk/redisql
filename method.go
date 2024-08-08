package redisql

import (
	"strings"
	"time"

	"github.com/qianlnk/redisql/redisql_parse"
)

const (
	REDISQL_USE             = iota //use databasename;
	REDISQL_SHOW_DATABASES         //show databases;
	REDISQL_SHOW_TABLES            //show tables;
	REDISQL_SHOW_INDEX             //show index from tablename;
	REDISQL_DESC                   //desc tablename;
	REDISQL_CREATE_DATABASE        //create database databasename;
	REDISQL_CREATE_TABLE           //create table tbname(field1,type1...);
	REDISQL_CREATE_INDEX           //create index indexname on tablename(fieldname);
	REDISQL_INSERT                 //insert into tablename(field1...) values(value1...);
	REDISQL_SELECT                 //select field1... from table1... where case1... limit start end
	REDISQL_UPDATE                 //update tablename set field1=value1 where case1...
	REDISQL_DELETE                 //delete from tablename where case1...
	REDISQL_DROP_DATABASE          //drop database databasename
	REDISQL_DROP_TABLE             //drop table tablename
	REDISQL_EXIT                   //exit
	REDISQL_HELP                   //help
	REDISQL_EMPTY
)

const (
	QUERY_OK_0 = "Query OK, 0 rows affected"
	QUERY_OK_1 = "Query OK, 1 row affected"
)

type QueryRes struct {
	Type      int
	ParseTime float64
	QueryTime float64
	Result    interface{}
}

func Query(sql string) (*QueryRes, error) {
	res := new(QueryRes)

	parseStart := time.Now()

	sqlParse := redisql_parse.GetSql(sql)

	res.ParseTime = time.Now().Sub(parseStart).Seconds()

	res.Type = sqlParse.Type

	queryStart := time.Now()
	switch sqlParse.Type {
	case REDISQL_USE:
		err := ChangeDatabase(sqlParse.DatabaseName)
		if err != nil {
			return nil, err
		}
		res.QueryTime = time.Now().Sub(queryStart).Seconds()
		res.Result = "Database changed"
		break
	case REDISQL_SHOW_DATABASES:
		dbs, usetime, err := GetDatabases()
		if err != nil {
			return nil, err
		}
		res.QueryTime = usetime
		res.Result = dbs
		break
	case REDISQL_SHOW_TABLES:
		tbs, usetime, err := GetTables()
		if err != nil {
			return nil, err
		}
		res.QueryTime = usetime
		res.Result = tbs
		break
	case REDISQL_SHOW_INDEX:
		indexs, err := getIndexs(sqlParse.TableName)
		if err != nil {
			return nil, err
		}
		res.QueryTime = time.Now().Sub(queryStart).Seconds()
		res.Result = indexs
		break
	case REDISQL_DESC:
		tbInfo, usetime, err := GetTableInfo(sqlParse.TableName)
		if err != nil {
			return nil, err
		}
		res.QueryTime = usetime
		res.Result = tbInfo
		break
	case REDISQL_CREATE_DATABASE:
		err := CreateDatabase(sqlParse.DatabaseName)
		if err != nil {
			return nil, err
		}
		res.QueryTime = time.Now().Sub(queryStart).Seconds()
		res.Result = QUERY_OK_1
		break
	case REDISQL_CREATE_TABLE:
		tmpTable := new(Table)
		tmpTable.Name = sqlParse.TableName
		for _, v := range sqlParse.FieldTypes {
			tmpTable.Fields = append(tmpTable.Fields, v.Field)
			tmpTable.Types = append(tmpTable.Types, v.Type)
		}
		err := tmpTable.CREATE()
		if err != nil {
			return nil, err
		}
		res.QueryTime = time.Now().Sub(queryStart).Seconds()
		res.Result = QUERY_OK_0
		break
	case REDISQL_CREATE_INDEX:
		tmpTable := new(Table)
		tmpTable.Name = sqlParse.TableName
		for _, v := range sqlParse.FieldTypes {
			tmpTable.Fields = append(tmpTable.Fields, v.Field)
		}
		err := tmpTable.INDEX()
		if err != nil {
			return nil, err
		}
		res.QueryTime = time.Now().Sub(queryStart).Seconds()
		res.Result = QUERY_OK_0
		break
	case REDISQL_INSERT:
		tmpInsert := new(Insert)
		tmpInsert.Into = sqlParse.TableName
		for _, v := range sqlParse.FieldValues {
			tmpInsert.Fields = append(tmpInsert.Fields, v.Field)
			tmpInsert.Values = append(tmpInsert.Values, v.Value)
		}
		err := tmpInsert.INSERT()
		if err != nil {
			return nil, err
		}
		res.QueryTime = time.Now().Sub(queryStart).Seconds()
		res.Result = QUERY_OK_1
		break
	case REDISQL_SELECT:
		tmpSelect := new(Select)
		for _, f := range sqlParse.FieldAliases {
			var tmpf SltField
			tmpf.TableAlias = f.TableAlias
			tmpf.Name = f.Field
			tmpf.Alias = f.Alias
			tmpSelect.Fields = append(tmpSelect.Fields, tmpf)
		}

		tmpFrom := make(SltTable)
		for _, t := range sqlParse.TableAliases {
			tmpFrom[t.Alias] = t.Table
		}
		tmpSelect.Froms = tmpFrom

		tmpSelect.Where = strings.Fields(sqlParse.Where)

		tmpSelect.Top = sqlParse.Top

		datas, err := tmpSelect.LIMIT(sqlParse.Limit.Start, sqlParse.Limit.End).SELECT()
		if err != nil {
			return nil, err
		}
		res.QueryTime = time.Now().Sub(queryStart).Seconds()
		res.Result = datas
		break
	case REDISQL_UPDATE:
		break
	case REDISQL_DELETE:
		break
	case REDISQL_DROP_DATABASE:
		break
	case REDISQL_DROP_TABLE:
		break
	case REDISQL_EXIT:
		break
	case REDISQL_HELP:
		break
	case REDISQL_EMPTY:
		break
	default:
	}

	return res, nil
}

func ToArray(reply interface{}) [][]string {
	var res [][]string
	switch reply.(type) {
	case []string:
		datas := reply.([]string)
		for _, data := range datas {
			var dt []string
			dt = append(dt, data)
			res = append(res, dt)
		}
		break
	case map[string][]string:
		datas := reply.(map[string][]string)
		for k, v := range datas {
			var dt []string
			var vs string
			for i, tmpv := range v {
				if i > 0 {
					vs += ","
				}
				vs += tmpv
			}
			dt = append(dt, k)
			dt = append(dt, vs)
			res = append(res, dt)
		}
		break
	case [][]string:
		res = reply.([][]string)
		break
	}

	return res
}
