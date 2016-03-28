package redisql

import (
	"fmt"
	"testing"
)

func TestInsert(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	ChangeDatabase("lnkgift")
	err := INTO("student").FIELDS("sid", "name", "class", "age", "date").VALUES(1, "xiezj", "jisuanji1001", 26, "2016-03-21 04:00:00").INSERT()
	err = INTO("student").FIELDS("sid", "name", "class", "age", "date").VALUES(2, "zhuojf", "jisuanji1002", 25, "2016-03-22 04:01:00").INSERT()
	fmt.Println(err)
}

//insert into user(id, name, age, pwd) values(1, 'xzj', 26, '123456')
//select id, name, age, pwd from uer where ...
//delete from user where ...
//update user set age = 27 where name = 'xzj'

//alter table user add index(index_name) (name)
