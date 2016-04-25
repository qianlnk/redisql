package main

import (
	"qianno.xie/redisql/redisql_parse"
	"fmt"
)

func main() {
        res := redisql_parse.GetSql("use student")
        fmt.Println(res)
}
