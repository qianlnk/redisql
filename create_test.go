package redisql

import (
	"fmt"
	"testing"
)

func TestDatabase(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	err := CreateDatabase("lnkgift")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func TestCreate(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	ChangeDatabase("lnkgift")
	err := TABLE("log").FIELDS("userid", "operate", "detail", "data").TYPES("int", "string", "string, datatime").CREATE()
	fmt.Println(err)
}

func TestIndex(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	ChangeDatabase("lnkgift")
	err := TABLE("log").FIELDS("userid").INDEX()
	err = TABLE("log").FIELDS("operate").INDEX()
	err = TABLE("log").FIELDS("detail").INDEX()
	err = TABLE("log").FIELDS("data").INDEX()
	fmt.Println(err)
}
