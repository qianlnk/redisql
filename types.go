package redisql

import (
	"github.com/garyburd/redigo/redis"
)

//data type design
//number include INT FLOAT DOUBLE
//string include CHAR、VARCHAR、BINARY、VARBINARY、BLOB、TEXT、ENUM
//date include DATA DATATIME
//ignore enum or set
//number and date will support range query
//string will support fuzzy search
const (
	REDISQL_TYPE_NUMBER = "number"
	REDISQL_TYPE_STRING = "string"
	REDISQL_TYPE_DATE   = "date"
)

//
const (
	REDISQL_PRIORITY_GREATER = 1
	REDISQL_PRIORITY_EQUAL   = 0
	REDISQL_PRIORITY_LESS    = -1
	REDISQL_PRIORITY_ERROR   = -2
)

//redis store key design
const (
	REDISQL_DATABASES     = "databases"                     //hset databases lnkgift 1
	REDISQL_CONDITION_SN  = "%s.condition.sn"               //query condition sn
	REDISQL_TABLES        = "%s.tables"                     //hset lnkgift.tables user 1
	REDISQL_COUNT         = "%s.tables.count"               //hset lnkgift.tables.count user 1
	REDISQL_FIELDS        = "%s.%s.fields"                  //hset lnkgift.user.fields name string
	REDISQL_UNIQUE        = "%s.%s.unique"                  //sadd lnkgift.user.unique name
	REDISQL_INDEXS        = "%s.%s.indexs"                  //hset lnkgift.user.indexs index_name [name]
	REDISQL_DATAS         = "%s.%s.data.%s"                 //hset lnkgift.user.data.1 name qianno
	REDISQL_INDEX_DATAS   = "%s.%s.index.%s"                //string: sadd lnkgift.user.index.name.qianno 1 or number,date: zadd lnkgift.user.index.age 26 1
	REDISQL_TMP_CONDITION = "%s.%s.tmp.condition.%s"        //sadd lnkgift.user.tmp.condition.1 1 2 3 4
	REDISQL_TMP_DESCARTES = "%s.tmp.condition.descartes.%s" //hset lnkgift.tmp.condition.descartes.1 a_1_b_2_c_1_
)

type RedisConnect struct {
	server      string
	port        string
	password    string
	protocol    string
	idleMax     int
	idleTimeout int
	pool        *redis.Pool
}

var (
	DB RedisConnect
	//As redis is a single thread
	redisdb  int
	database string
)
