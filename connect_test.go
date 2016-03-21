package redisql_test

import (
	//"fmt"
	"qianno.xie/redisql"
	"testing"
)

func TestConnect(*testing.T) {
	redisql.Connect("192.168.99.100", "6379", "", "tcp", 5, 120)
	//	myDB := redisql.getConn()
	//	defer myDB.Close()
	//	res, err := myDB.Do("SET", "test", "test123")
	//	fmt.Print(res, err)
}
