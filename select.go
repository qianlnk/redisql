package redisql

import (
	"errors"
	"fmt"
	"strings"
)

/*
create table user(
id int(11) auto_increment PRIMARY KEY,
name varchar(64) not null default '',
age int(11) not null default 0,
city varchar(32) not null default ''
)

create table log(
id int(11) auto_increment PRIMARY KEY,
userid int(11) not null default 0,
operate varchar(128) not null default '',
detail VARCHAR(128)not null default '',
`date` datetime not null default '1900-01-01 00:00:00'
)
*/

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

type SltTable struct {
	Name  string
	Alias string
}

type SltField struct {
	TableAlias string
	Name       string
	Alias      string
}
type Select struct {
	Froms  []SltTable //query table must have an alias
	Fields []SltField //query field must like tablealias.fieldname, it can have an alias also.
	Where  string
	And    string
	Or     string
}

func FROM(tables ...string) *Select {
	var tmptables []SltTable
	for _, t := range tables {
		tabs := strings.Split(t, ",")
		var tab SltTable
		for _, tt := range tabs {
			tas := strings.Fields(tt)
			if len(tas) != 2 {
				panic(errors.New("tables wrong."))
			} else {
				tab.Name = tas[0]
				tab.Alias = tas[1]
			}
			tmptables = append(tmptables, tab)
		}
	}

	return &Select{
		Froms: tmptables,
	}
}

func (slt *Select) FIELDS(fields ...string) *Select {
	var tmpfields []SltField
	for _, f := range fields {
		fs := strings.Split(f, ",")
		for _, ff := range fs {
			var field SltField
			fas := strings.Fields(ff)
			tabField := strings.Split(fas[0], ".")
			if len(tabField) != 2 {
				panic(errors.New("fields wrong1."))
			} else {
				if len(fas) == 2 {
					field.TableAlias = tabField[0]
					field.Name = tabField[1]
					field.Alias = fas[1]
				} else if len(fas) == 1 { //if no define alias,field's alias is it self
					field.TableAlias = tabField[0]
					field.Name = tabField[1]
					field.Alias = tabField[1]
				} else {
					panic(errors.New("fields wrong2."))
				}
				tmpfields = append(tmpfields, field)
			}
		}
	}

	return &Select{
		Froms:  slt.Froms,
		Fields: tmpfields,
	}
}

func (slt *Select) check() error {
	//check table
	for _, t := range slt.Froms {
		if existsTable(t.Name) == false {
			return errors.New(fmt.Sprintf("table '%s' not exists.", t.Name))
		}
	}
	//check field
	for _, f := range slt.Fields {
		var tmptab string
		for _, t := range slt.Froms {
			if t.Alias == f.TableAlias {
				tmptab = t.Name
			}
		}
		if tmptab == "" {
			return errors.New(fmt.Sprintf("Unknow alias '%s' int table list.", f.TableAlias))
		} else {
			if existsField(tmptab, f.Name) == false {
				return errors.New(fmt.Sprintf("Unknow cloumn '%s.%s' int field list.", f.TableAlias, f.Name))
			}
		}
	}
	return nil
}

func (slt *Select) SELECT() error {
	fmt.Println(slt)
	err := slt.check()
	if err != nil {
		return err
	}
	return nil
}
