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
		Name []string `json:"myname"`
		Age  []string `json:"myage"`
		City []string `json:"mycity"`
	}
	var testres test
	res, err := FROM("student a, score b").FIELDS("a.name myname, a.age myage, b.math, b.english, b.chinese").WHERE("a.id = b.sid and ((a.age>=22)or(b.english >80))").SELECT()
	if err != nil {
		fmt.Println(res, err)
	}
	json.Unmarshal(res, &testres)
	fmt.Println(testres)
}
