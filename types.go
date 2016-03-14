package redisql

import (
	"github.com/garyburd/redigo/redis"
)

const (
	REDIS_KEY_TABLES_MAX_ID = "sys_tables_max_id"    //hset sys_tables_max_id tablename maxid
	REDIS_KEY_INDEXS        = "sys_indexs.table:%s"  //hset sys_indexs.table:tablename indexname fields
	REDIS_KEY_FIELDS        = "sys_fields.table:%s"  //hset sys_fields.table:tablename field type
	REDIS_KEY_UNIQUES       = "sys_uniques.table:%s" //sadd sys_uniques.table:tablename field
	SQL_CREATE              = "create"
	SQL_INSERT              = "insert"
	SQL_UNIQUE              = "unique"
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

type Redisql struct {
	Database int
	Table    string
	Fields   []string
	Types    []string
	Values   []interface{}
}

var (
	DB       RedisConnect
	selectdb int
)
