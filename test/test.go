package main

import (
	"fmt"

	"github.com/qianlnk/redisql/redisql_parse"
)

func main() {
	res := redisql_parse.GetSql("select a.id, a.name, b.math, b.english from student a, scores b, cause c where a.id = b.sid and(a.name like 'xie%' or b.math > 80);")
	fmt.Println(res)
}
