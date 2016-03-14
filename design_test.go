package redisql

import (
	"fmt"
	"testing"
)

func TestInsert(*testing.T) {
	Connect("192.168.99.100", "6379", "", "tcp", 5, 120)
	SelectDB(1)
	var testRedis Redisql
	//err := testRedis.TABLE("task").FIELDS("name", "level", "count").TYPES("string", "int", "int").CREATE()
	err := testRedis.TABLE("task").FIELDS("name", "level").UNIQUE()
	//err := testRedis.INTO("task").FIELDS("name", "level", "count").VALUES("hyrra cache", 1, 2).INSERT()
	fmt.Println(err)
}
