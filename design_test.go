package redisql

import (
	"fmt"
	"testing"
)

func Testtmp(*testing.T) {
	Connect("192.168.99.100", "6379", "", "tcp", 5, 120)
	SelectDB(1)
	var testRedis Redisql
	//err := testRedis.TABLE("task").FIELDS("name", "level", "count").TYPES("string", "int", "int").CREATE()
	err := testRedis.TABLE("task").FIELDS("name", "level").UNIQUE()
	//err := testRedis.INTO("task").FIELDS("name", "level", "count").VALUES("hyrra cache", 1, 2).INSERT()
	fmt.Println(err)
}

//insert into user(id, name, age, pwd) values(1, 'xzj', 26, '123456')
//select id, name, age, pwd from uer where ...
//delete from user where ...
//update user set age = 27 where name = 'xzj'

//alter table user add index(index_name) (name)
