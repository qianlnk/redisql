package main

import (
	"fmt"
	"qianno.xie/redisql"
	"time"
)

func main() {
	redisql.Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	redisql.Selectdb(1)
	redisql.Query("create database std")
	redisql.Query("use std")
	redisql.Query("create table student(name varchar(32), class varchar(32), age int(11), joindate datetime)")
	redisql.Query("create table score(sid int(11), math int(11), english int(11), chinese int(11))")

	insertStart := time.Now()
	for i := 0; i < 100000; i++ {
		_, err := redisql.Query(fmt.Sprintf("insert into student(name, class, age, joindate) values('name%d', 'class%d', %d, '2016-04-27')", i, i, i))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("insert student 100000, use time ", time.Now().Sub(insertStart).Seconds())

	scoreStart := time.Now()
	for i := 0; i < 100000; i++ {
		_, err := redisql.Query(fmt.Sprintf("insert into score(sid, math, english, chinese) values(%d, %d, %d, %d)", i, i, i, i))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("insert score 100000, use time ", time.Now().Sub(scoreStart).Seconds())
}
