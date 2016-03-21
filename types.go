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

	SYS_TABLES_MAX_ID = "sys_tables_max_id"        //hset sys_tables_max_id tablename maxid
	SYS_INDEXS        = "sys_indexs.table:%s"      //hset sys_indexs.table:tablename indexname fields
	SYS_FIELDS        = "sys_fields.table:%s"      //hset sys_fields.table:tablename field type
	SYS_UNIQUES       = "sys_uniques.table:%s"     //sadd sys_uniques.table:tablename field
	USER_TABLE        = "user_table:%s.id:%s"      //hset user_table:tablename.id:id field value
	USER_INDEX        = "user_table:%s.index:%s%s" //sadd user_table:tablename.index:indexname.field:value id

	SQL_CREATE = "create"
	SQL_INSERT = "insert"
	SQL_UNIQUE = "unique"
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
	selectdb int
	database string
)
