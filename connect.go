package redisql

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

func Connect(server, port, password, protocol string, idleMax, idleTimeput int) {
	fmt.Println("redis server connect...")
	if DB.server == "" {
		if server != "" {
			DB.server = server
		} else {
			DB.server = "127.0.0.1"
		}

		if port != "" {
			DB.port = port
		} else {
			DB.port = "6379"
		}

		DB.password = password

		if protocol != "" {
			DB.protocol = protocol
		} else {
			DB.protocol = "tcp"
		}

		if idleMax > 0 {
			DB.idleMax = idleMax
		} else {
			DB.idleMax = 5
		}

		if idleTimeput > 0 {
			DB.idleTimeout = idleTimeput
		} else {
			DB.idleTimeout = 120
		}
		DB.pool = &redis.Pool{
			MaxIdle: DB.idleMax,
			//MaxActive:   0,
			IdleTimeout: time.Duration(DB.idleTimeout) * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial(DB.protocol, DB.server+":"+DB.port, redis.DialDatabase(0), redis.DialPassword(DB.password))
				if err != nil {
					fmt.Println("dial err", err)
					return nil, err
				}

				if _, err = c.Do("PING"); err != nil {
					c.Close()
					fmt.Println("ping", err)
					return nil, err
				}

				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if _, err := c.Do("PING"); err != nil {
					fmt.Println("ping", err)
					return err
				}
				return nil
			},
		}
	}
}

func Selectdb(db int) {
	fmt.Println("select db start...")
	redisdb = db
}

func ChangeDatabase(db string) {
	fmt.Println("change database start...")
	if len(strings.Trim(db, " ")) <= 0 {
		panic(errors.New(fmt.Sprintf("can not change database to ''.")))
	}

	if existsDatabase(db) == false {
		panic(fmt.Sprintf("no database named '%s', please call func 'CreateDatabase'.", db))
	}

	database = db
}

func GetDbInfo() (int, string) {
	return redisdb, database
}
