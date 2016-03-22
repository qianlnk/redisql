package redisql

import (
	"github.com/garyburd/redigo/redis"
)

const (
	REDISQL_DATABASES   = "databases"       //hset databases lnkgift 1
	REDISQL_TABLES      = "%s.tables"       //hset lnkgift.tables user 1
	REDISQL_COUNT       = "%s.tables.count" //hset lnkgift.tables.count user 1
	REDISQL_FIELDS      = "%s.%s.fields"    //hset lnkgift.user.fields name string
	REDISQL_UNIQUE      = "%s.%s.unique"    //sadd lnkgift.user.unique name
	REDISQL_INDEXS      = "%s.%s.indexs"    //hset lnkgift.user.indexs index_name [name]
	REDISQL_DATAS       = "%s.%s.data.%s"   //hset lnkgift.user.1 name qianno
	REDISQL_INDEX_DATAS = "%s.%s.index.%s"  //sadd lnkgift.user.name.qianno 1
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
