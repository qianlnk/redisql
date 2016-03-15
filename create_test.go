package redisql

import (
	"fmt"
	"testing"
)

func TestCreate(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	SelectDB(1)
	err := TABLE("task").FIELDS("name", "level", "count").TYPES("string", "int", "int").CREATE()
	fmt.Println(err)
}

func TestIndex(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	SelectDB(1)
	err := TABLE("task").FIELDS("level,count").INDEX()
	fmt.Println(err)
}
