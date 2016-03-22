package redisql

import (
	"fmt"
	"testing"
)

func TestInsert(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	ChangeDatabase("lnkgift")
	err := INTO("log").FIELDS("userid", "operate", "detail", "data").VALUES(1, "login", "login at sh", "2016-03-22 04:00:00").INSERT()
	err = INTO("log").FIELDS("userid", "operate", "detail", "data").VALUES(1, "insert", "self info", "2016-03-22 04:01:00").INSERT()
	fmt.Println(err)
}

//insert into user(id, name, age, pwd) values(1, 'xzj', 26, '123456')
//select id, name, age, pwd from uer where ...
//delete from user where ...
//update user set age = 27 where name = 'xzj'

//alter table user add index(index_name) (name)
