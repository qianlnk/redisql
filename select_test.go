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
	res, err := FROM("user a").FIELDS("a.name myname, a.age myage, a.city mycity").WHERE("(a.age=24)and(a.name like '%qian%')").SELECT()
	if err != nil {
		fmt.Println(res, err)
	}
	json.Unmarshal(res, &testres)
	fmt.Println(testres)
}
