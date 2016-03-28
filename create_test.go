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
	err := TABLE("student").FIELDS("sid", "name", "class", "age", "date").TYPES("number", "string", "string, number, date").CREATE()
	fmt.Println(err)
}

func TestIndex(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	ChangeDatabase("lnkgift")
	//err := TABLE("student").FIELDS("sid").INDEX()
	//err = TABLE("student").FIELDS("date").INDEX()
	err := TABLE("student").FIELDS("age").INDEX()
	err = TABLE("student").FIELDS("name").INDEX()
	err = TABLE("student").FIELDS("class").INDEX()
	fmt.Println(err)
}
