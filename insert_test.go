package redisql

import (
	"fmt"
	"testing"
)

func TestInsert(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	ChangeDatabase("lnkgift")
	err := INTO("student").FIELDS("xuehao", "name", "class", "age", "date").VALUES(3100301210, "xiezhenjia", "jisuanji1002", 26, "2016-03-24 04:00:00").INSERT()
	err = INTO("student").FIELDS("xuehao", "name", "class", "age", "date").VALUES(3100301211, "zhuojianfei", "jisuanji1002", 25, "2016-03-22 04:01:00").INSERT()
	err = INTO("student").FIELDS("xuehao", "name", "class", "age", "date").VALUES(3100301213, "yangjunhui", "jisuanji1002", 24, "2016-03-23 04:01:00").INSERT()
	err = INTO("student").FIELDS("xuehao", "name", "class", "age", "date").VALUES(3100301214, "yuyong", "jisuanji1002", 25, "2016-03-20 04:01:00").INSERT()
	err = INTO("student").FIELDS("xuehao", "name", "class", "age", "date").VALUES(3100301215, "zhengzhiqiang", "jisuanji1002", 26, "2016-03-21 04:01:00").INSERT()
	err = INTO("student").FIELDS("xuehao", "name", "class", "age", "date").VALUES(3100301216, "zhanglei", "jisuanji1002", 23, "2016-03-22 04:01:00").INSERT()
	err = INTO("score").FIELDS("sid", "math", "english", "chinese").VALUES(1, 98, 83, 100).INSERT()
	err = INTO("score").FIELDS("sid", "math", "english", "chinese").VALUES(2, 78, 81, 89).INSERT()
	err = INTO("score").FIELDS("sid", "math", "english", "chinese").VALUES(3, 94, 70, 76).INSERT()
	err = INTO("score").FIELDS("sid", "math", "english", "chinese").VALUES(4, 70, 85, 88).INSERT()
	err = INTO("score").FIELDS("sid", "math", "english", "chinese").VALUES(5, 68, 90, 88).INSERT()
	err = INTO("score").FIELDS("sid", "math", "english", "chinese").VALUES(6, 95, 83, 93).INSERT()
	err = INTO("student").FIELDS("xuehao", "name", "class", "age", "date").VALUES(3100301217, "wenqicheng", "jisuanji1002", 22, "2016-03-23 04:01:00").INSERT()
	err = INTO("score").FIELDS("sid", "math", "english", "chinese").VALUES(7, 98, 73, 89).INSERT()
	fmt.Println(err)
}

//insert into user(id, name, age, pwd) values(1, 'xzj', 26, '123456')
//select id, name, age, pwd from uer where ...
//delete from user where ...
//update user set age = 27 where name = 'xzj'

//alter table user add index(index_name) (name)
