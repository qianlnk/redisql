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
	//	err := TABLE("student").FIELDS("xuehao", "name", "class", "age", "date").TYPES("number", "string", "string, number, date").CREATE()
	//	fmt.Println(err)
	//	err = TABLE("score").FIELDS("sid", "math", "english", "chinese").TYPES("number", "number", "number, number").CREATE()
	//	fmt.Println(err)
	err := TABLE("family").FIELDS("sid", "member", "relationship", "phone", "job", "age").TYPES("number, string", "string", "string, string, number").CREATE()
	fmt.Println(err)
}

func TestIndex(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	ChangeDatabase("lnkgift")
	err := TABLE("student").FIELDS("xuehao").INDEX()
	err = TABLE("student").FIELDS("name").INDEX()
	err = TABLE("student").FIELDS("class").INDEX()
	err = TABLE("student").FIELDS("age").INDEX()
	err = TABLE("student").FIELDS("date").INDEX()
	err = TABLE("score").FIELDS("sid").INDEX()
	err = TABLE("score").FIELDS("math").INDEX()
	err = TABLE("score").FIELDS("english").INDEX()
	err = TABLE("score").FIELDS("chinese").INDEX()
	err = TABLE("student").FIELDS("id").INDEX()
	err = TABLE("score").FIELDS("id").INDEX()
	fmt.Println(err)
}
