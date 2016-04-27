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

func show(fields []string, datas [][]string, parsetime float64, querytime float64) {
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
		fmt.Printf("Empty set (p: %-.2f sec q: %-.2f sec)\n\n", parsetime, querytime)
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
	fmt.Printf("%d rows int set (p: %-.2f sec q: %-.2f sec)\n\n", count, parsetime, querytime)
}

func printResInfo(res *redisql.QueryRes) {
	fmt.Printf("%s (p: %.2f sec, q: %.2f sec)\n\n", res.Result, res.ParseTime, res.QueryTime)
}
func ui(res *redisql.QueryRes) {
	switch res.Type {
	case redisql.REDISQL_USE:
		printResInfo(res)
		break
	case redisql.REDISQL_SHOW_DATABASES:
		var fields []string
		fields = append(fields, "Databases")
		show(fields, redisql.ToArray(res.Result), res.ParseTime, res.QueryTime)
		break
	case redisql.REDISQL_SHOW_TABLES:
		var fields []string
		_, dbname := redisql.GetDbInfo()
		fields = append(fields, fmt.Sprintf("Tables_in_%s", dbname))
		show(fields, redisql.ToArray(res.Result), res.ParseTime, res.QueryTime)
		break
	case redisql.REDISQL_SHOW_INDEX:
		var fields []string
		fields = append(fields, "index_name", "conlumn")
		show(fields, redisql.ToArray(res.Result), res.ParseTime, res.QueryTime)
		break
	case redisql.REDISQL_DESC:
		var fields []string
		fields = append(fields, "Field")
		fields = append(fields, "type")
		show(fields, redisql.ToArray(res.Result), res.ParseTime, res.QueryTime)
		break
	case redisql.REDISQL_CREATE_DATABASE:
		printResInfo(res)
		break
	case redisql.REDISQL_CREATE_TABLE:
		printResInfo(res)
		break
	case redisql.REDISQL_CREATE_INDEX:
		printResInfo(res)
		break
	case redisql.REDISQL_INSERT:
		printResInfo(res)
		break
	case redisql.REDISQL_SELECT:
		var fields []string
		datas := redisql.ToArray(res.Result)
		fields = datas[0]
		datas = append(datas[:0], datas[1:]...)
		show(fields, datas, res.ParseTime, res.QueryTime)
		break
	case redisql.REDISQL_UPDATE:
		break
	case redisql.REDISQL_DELETE:
		break
	case redisql.REDISQL_DROP_DATABASE:
		break
	case redisql.REDISQL_DROP_TABLE:
		break
	case redisql.REDISQL_EXIT:
		break
	case redisql.REDISQL_HELP:
		break
	case redisql.REDISQL_EMPTY:
		break
	default:
		break
	}
}

func main() {
	redisql.Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	redisql.Selectdb(1)
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
