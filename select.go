package redisql

import (
	"errors"
	"strings"
)

//how to use
//eg: FIELDS("a.name, a.age, a.city, b.userid, b.operate, b.detail, b.data")
//	  .FROM("user, log").ALIAS("a, b")
//	  .WHERE(a.id = b.userid).AND(a.age = 24).OR(b.operate = "login").SELECT()
type Inner struct {
	Inner string
	Alias string
	On    string
	And   string
	Or    string
}

type Left struct {
	Left  string
	Alias string
	On    string
	And   string
	Or    string
}

type Right struct {
	Right string
	Alias string
	On    string
	And   string
	Or    string
}

type Select struct {
	Froms  map[string]string
	Fields map[string]string
	Where  string
	And    string
	Or     string
}

func FROM(tables ...string) *Select {
	tmptables := make(map[string]string)
	for _, t := range tables {
		tabs := strings.Split(t, ",")
		for _, tt := range tabs {
			tas := strings.Split(tt, " ")
			if len(tas) <= 0 || len(tas) > 2 {
				panic(errors.New("tables wrong."))
			}

			if len(tas) == 2 {
				tmptables[tas[0]] = tas[1]
			} else {
				tmptables[tas[0]] = tas[0]
			}
		}
	}

	return &Select{
		Froms: tmptables,
	}
}

func (slt *Select) FIELDS(fields ...string) *Select {
	tmpfields := make(map[string]string)
	for _, f := range fields {
		fs := strings.Split(f, ",")
		for _, ff := range fs {
			fas := strings.Split(ff, " ")
			if len(fas) <= 0 || len(fas) > 2 {
				panic(errors.New("fields wrong."))
			}

			if len(fas) == 2 {
				tmpfields[fas[0]] = fas[1]
			} else {
				tmpfields[fas[0]] = fas[0]
			}
		}
	}

	return &Select{
		Froms:  slt.Froms,
		Fields: tmpfields,
	}
}

func (slt *Select) SELECT() error {
	return nil
}
