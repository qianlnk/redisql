package redisql

import (
	"fmt"
	"testing"
)

func TestSelect(*testing.T) {
	Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	Selectdb(0)
	ChangeDatabase("lnkgift")
	err := FROM("user a, log b").FIELDS("a.name myname, a.age myage, a.city mycity, b.userid, b.operate", "b.detail detail, b.data").SELECT()
	fmt.Println(err)
}
