package redisql_parse

/*
#cgo CPPFLAGS: -I .
#cgo LDFLAGS: -L . -lparse
#include "yacc/parse_redisql.h"
*/
import "C"
import (
	"fmt"
)

type FieldType struct {
	Field string
	Type  string
}

type FieldValue struct {
	Type  int
	Value interface{}
}

type FieldAlias struct {
	TableAlias string
	Field      string
	Alias      string
}

type TableAlias struct {
	Table string
	Alias string
}

type Limit struct {
	Start int
	End   int
}

type RedisqlNode struct {
	Type           int
	DatabaseName   string
	TableName      string
	IndexName      string
	FieldTypeNum   int
	FieldTypes     []FieldType
	FieldValueNum  int
	FieldValues    []FieldValue
	FieldAliaseNum int
	FieldAliases   []FieldAlias
	TableAliaseNum int
	TableAliases   []TableAlias
	Where          string
	Top            int
	Limit          Limit
}

func GetSql(sql string) *RedisqlNode {
	var res *RedisqlNode
	C.redisql_parse(C.CString(sql))
	fmt.Println(C.getType())
	res.Type = C.int(C.getType())
	switch res.Type {
	case C.REDISQL_USE:
		res.DatabaseName = C.GoString(C.getDatabaseName())
		break
	default:
		break
	}

	fmt.Println(res)
	return res
}
