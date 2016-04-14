package main

import (
	"bufio"
	"fmt"
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
const (
	REDISQL_USE       = "use"
	REDISQL_SHOW      = "show"
	REDISQL_DATABASES = "databases"
	REDISQL_TABLES    = "tables"
	REDISQL_DESC      = "desc"
	REDISQL_CREATE    = "create"
	REDISQL_DATABASE  = "database"
	REDISQL_TABLE     = "table"
	REDISQL_INDEX     = "index"
	REDISQL_ON        = "on"
	REDISQL_SELECT    = "select"
	REDISQL_FROM      = "from"
	REDISQL_WHERE     = "where"
	REDISQL_TOP       = "top"
	REDISQL_LIMIT     = "limit"
	REDISQL_INSERT    = "insert"
	REDISQL_INTO      = "into"
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

func main() {
	redisql.Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	redisql.Selectdb(0)
	fmt.Printf(WELCOME_INFO)
	database := "redisql"

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s> ", database)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.Replace(cmd, "\n", "", -1)
		cmd = strings.Replace(cmd, "\r", "", -1)
		cmd = strings.Replace(cmd, ";", "", -1)
		specialChar1 := []string{"!=", ">=", "<="}
		specialChar2 := []string{"=", ">", "<", "(", ")"}
		specialChar3 := []string{"! =", ">  =", "<  ="}
		for _, c := range specialChar1 {
			cmd = strings.Replace(cmd, c, " "+c+" ", -1)
		}
		for _, c := range specialChar2 {
			cmd = strings.Replace(cmd, c, " "+c+" ", -1)
		}
		for i, c := range specialChar3 {
			cmd = strings.Replace(cmd, c, specialChar1[i], -1)
		}
		cmd = strings.ToLower(cmd)
		cmds := strings.Fields(cmd)
		if len(cmds) == 0 {
			continue
		}
		switch cmds[0] {
		case REDISQL_USE:
			if len(cmds) != 2 {
				fmt.Println("cmd err.")
				break
			}
			err := redisql.ChangeDatabase(cmds[1])
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			database = cmds[1]
			break
		case REDISQL_SHOW:
			if len(cmds) < 2 {
				fmt.Println("cmd err.")
				break
			}
			switch cmds[1] {
			case REDISQL_DATABASES:
				if len(cmds) != 2 {
					fmt.Println("cmd err.")
					break
				}
				dbs, usetime, err := redisql.GetDatabases()
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				var fields []string
				var datas [][]string
				fields = append(fields, "Database")
				for _, db := range dbs {
					var tmpdb []string
					tmpdb = append(tmpdb, db)
					datas = append(datas, tmpdb)
				}
				show(fields, datas, usetime)
				break
			case REDISQL_TABLES:
				if len(cmds) != 2 {
					fmt.Println("cmd err.")
					break
				}
				tbs, usetime, err := redisql.GetTables()
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				var fields []string
				var datas [][]string
				fields = append(fields, "Tables_in_"+database)
				for _, tb := range tbs {
					var tmptb []string
					tmptb = append(tmptb, tb)
					datas = append(datas, tmptb)
				}
				show(fields, datas, usetime)
				break
			case REDISQL_INDEX:
				if len(cmds) != 4 {
					fmt.Println("cmd err.")
					break
				} else {
					if cmds[2] != REDISQL_FROM {
						fmt.Println("cmd err.")
						break
					}
				}
				break
			default:
				fmt.Println("cmd err.")
			}
			break
		case REDISQL_DESC:
			if len(cmds) != 2 {
				fmt.Println("cmd err.")
				break
			}
			datas, usetime, err := redisql.GetTableInfo(cmds[1])
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			var fields []string
			fields = append(fields, "Field")
			fields = append(fields, "Type")
			show(fields, datas, usetime)
			break
		case REDISQL_CREATE:
			if len(cmds) < 3 {
				fmt.Println("cmd err.")
				break
			}
			switch cmds[1] {
			case REDISQL_DATABASE:
				if len(cmds) != 3 {
					fmt.Println("cmd err.")
					break
				}
				err := redisql.CreateDatabase(cmds[2])
				if err != nil {
					fmt.Println(err.Error())
				}
				break
			case REDISQL_TABLE:
				if len(cmds) < 6 {
					fmt.Println("cmd err.")
					break
				}
				if cmds[3] != "(" || cmds[len(cmds)-1] != ")" {
					fmt.Println("cmd err.")
					break
				}
				var fields string
				var types string
				for i, f := range cmds {
					if i < 4 || i == len(cmds)-1 {
						continue
					}
					if i%2 == 0 {
						if i > 4 {
							fields += ","
						}
						fields += f
					} else {
						types = types + f + " "
					}

				}
				fmt.Println("table:", cmds[2], "fields:", fields)
				err := redisql.TABLE(cmds[2]).FIELDS(fields).TYPES(types).CREATE()
				if err != nil {
					fmt.Println(err.Error())
				}
				break
			case REDISQL_INDEX:
				if len(cmds) != 8 {
					fmt.Println("cmd err.")
					break
				}
				if cmds[3] != REDISQL_ON {
					fmt.Printf("expect 'on' befer '%s'.\n", cmd[3])
					break
				}
				if cmds[5] != "(" || cmds[len(cmds)-1] != ")" {
					fmt.Println("cmd err.")
					break
				}
				err := redisql.TABLE(cmds[4]).FIELDS(cmds[6]).INDEX()
				if err != nil {
					fmt.Println(err.Error())
				}
				break
			default:
				fmt.Printf("unknow cmd %s.\n", cmds[1])
			}
			break
		case REDISQL_SELECT:
			var fields string
			var fromflag int
			for i, v := range cmds {
				if v == REDISQL_SELECT {
					continue
				}
				if v == REDISQL_FROM {
					break
				}
				fields = fields + v + " "
				fromflag = i
			}
			var from string
			var whereflag int
			for i, v := range cmds {
				if i <= fromflag+1 {
					continue
				}
				if v == REDISQL_WHERE {
					break
				}
				from = from + v + " "
				whereflag = i
			}
			var where string
			var limitflag int
			for i, v := range cmds {
				if i <= whereflag+1 {
					continue
				}
				if v == REDISQL_WHERE {
					break
				}
				where = where + v + " "
				limitflag = limitflag - limitflag + i
			}
			fmt.Println("from:", from, "fields:", fields, "where:", where)
			res, err := redisql.FROM(from).FIELDS(fields).WHERE(where).SELECT()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(res)
			break
		default:
			fmt.Printf("unknow cmd %s.\n", cmds[0])
		}
	}

}
