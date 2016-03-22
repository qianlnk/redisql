package redisql

import (
	"fmt"
	"testing"
)

func TestDatabase(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Select(0)
	err := CreateDatabase("lnkgift")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func TestCreate(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Select(0)
	ChangeDatabase("lnkgift")
	err := TABLE("task").FIELDS("userid", "name", "level", "count").TYPES("int", "string", "int, int").CREATE()
	fmt.Println(err)
}

func TestIndex(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Select(0)
	ChangeDatabase("lnkgift")
	err := TABLE("user").FIELDS("age,city").INDEX()
	fmt.Println(err)
}
