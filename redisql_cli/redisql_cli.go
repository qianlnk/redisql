package main

import (
	"bufio"
	"fmt"
	"os"
	"qianno.xie/redisql"
	"strings"
)

const (
	REDISQL_CHANGE   = "change"
	REDISQL_DATABASE = "database"
	REDISQL_CREATE   = "create"
	REDISQL_TABLE    = "table"
	REDISQL_INDEX    = "index"
	REDISQL_ON       = "on"
	REDISQL_SELECT   = "select"
	REDISQL_FROM     = "from"
	REDISQL_WHERE    = "where"
	REDISQL_TOP      = "top"
	REDISQL_LIMIT    = "limit"
	REDISQL_INSERT   = "insert"
	REDISQL_INTO     = "into"
)

func main() {
	redisql.Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
	redisql.Selectdb(0)
	fmt.Println("redisql 1.0 127.0.0.1")
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
		fmt.Print(cmd)
		cmds := strings.Fields(cmd)
		if len(cmds) == 0 {
			continue
		}
		switch cmds[0] {
		case REDISQL_CHANGE:
			if len(cmds) != 3 {
				fmt.Println("cmd err.")
				break
			}
			if cmds[1] != REDISQL_DATABASE {
				fmt.Printf("unknow cmd %s.\n", cmds[1])
				break
			}
			redisql.ChangeDatabase(cmds[2])
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
				for i, f := range cmds {
					if i < 4 || i == len(cmds)-1 {
						continue
					}
					fields += f + " "
				}
				err := redisql.TABLE(cmds[2]).FIELDS(fields).CREATE()
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
				fromflag = fromflag - fromflag + i
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
				whereflag = whereflag - whereflag + i
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
