package main

import (
	"bufio"
	"fmt"
	//redigo "github.com/garyburd/redigo/redis"
	"os"
	"qianno.xie/redisql"
	"strconv"
	"strings"
)

const (
	WELCOME_INFO = "Welcome to the REDISQL monitor. Command end with '\\n'.\n" +
		"server version: 1.0.0\n\n" +
		"author: qianno.xie\n" +
		"github: github.com/qianlnk\n" +
		"email:  qianlnk@163.com\n\n" +
		"Type 'help' or '\\h' for help.\n\n"
)

func show(fields []string, datas [][]string, usetime float64) {
	maxlen := make(map[int]int)
	for i, data := range datas {
		for j, dt := range data {
			if i == 0 {
				if len(dt) > len(fields[j]) {
					maxlen[j] = len(dt)
				} else {
					maxlen[j] = len(fields[j])
				}
			} else {
				if len(dt) > maxlen[j] {
					maxlen[j] = len(dt)
				}
			}
		}
	}
	line := "+"
	for _, v := range maxlen {
		for i := 0; i <= v; i++ {
			line += "-"
		}
		line += "+"
	}
	if line == "+" {
		fmt.Printf("Empty set (%-.2f sec)\n\n", usetime)
		return
	}
	fmt.Println(line)
	fmt.Printf("|")
	for i, f := range fields {
		format := "%-" + strconv.Itoa(maxlen[i]+1) + "s|"
		fmt.Printf(format, f)
	}
	fmt.Printf("\n")
	fmt.Println(line)
	count := 0
	for _, data := range datas {
		fmt.Printf("|")
		for i, dt := range data {
			format := "%-" + strconv.Itoa(maxlen[i]+1) + "s|"
			fmt.Printf(format, dt)
		}
		fmt.Printf("\n")
		count++
	}
	fmt.Println(line)
	fmt.Printf("%d rows int set (%-.2f sec)\n\n", count, usetime)
}

func ui(res *redisql.QueryRes) {
	switch res.Type {
	case redisql.REDISQL_USE:
		fmt.Printf("%s (p: %.2f sec, q: %.2f sec)\n\n", res.Result, res.ParseTime, res.QueryTime)
		break
	case redisql.REDISQL_SHOW_DATABASES:
		var fields []string
		fields = append(fields, "Databases")
		var datas [][]string
		dbs := res.Result.([]string)
		for _, db := range dbs {
			var tmpdb []string
			tmpdb = append(tmpdb, db)
			datas = append(datas, tmpdb)
		}
		show(fields, datas, res.ParseTime)
	}
}

func main() {
	redisql.Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	redisql.Selectdb(0)
	fmt.Printf(WELCOME_INFO)
	database := "redisql"

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s> ", database)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.Trim(cmd, "\n")
		if cmd == "" {
			continue
		}
		res, err := redisql.Query(cmd)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ui(res)
	}

}
