package redisql

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSelect(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	ChangeDatabase("lnkgift")
	type test struct {
		Id      []string `json:"myid"`
		Name    []string `json:"myname"`
		Age     []string `json:"myage"`
		Class   []string `json:"myclass", "class"`
		Date    []string
		Math    []string `json:"math"`
		English []string `json: "english"`
		Chinese []string `json: "chinese"`
	}
	var testres test
	//res, err := FROM("student a").FIELDS("a.id myid, a.name myname, a.age myage, a.class").WHERE("a.name like '%zhenjia' or (a.age <= 25 and a.class = 'jisuanji1002')").LIMIT(0, 2).SELECT()
	res, err := FROM("student a, score b").FIELDS("a.id myid, a.name myname, a.age myage, a.class myclass, a.date, b.math, b.english, b.chinese").WHERE("a.id = b.sid and ((a.date>='20160322040100')and(b.english >80))").SELECT()
	if err != nil {
		fmt.Println(res, err)
	}
	json.Unmarshal(res, &testres)
	fmt.Println(testres)
}
