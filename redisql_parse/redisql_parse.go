package redisql_parse

import (
	"strconv"
	"strings"
	"unsafe"
)

/*
#cgo CPPFLAGS: -I .
#cgo LDFLAGS: -L . -lparse
#include "yacc/parse_redisql.h"
*/
import "C"

type FieldType struct {
	Field string
	Type  string
}

type FieldValue struct {
	Field string
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
	Type         int
	DatabaseName string
	TableName    string
	IndexName    string
	FieldTypes   []FieldType
	FieldValues  []FieldValue
	FieldAliases []FieldAlias
	TableAliases []TableAlias
	Where        string
	Top          int
	Limit        Limit
}

func getType() int {
	return int(C.getType())
}

func getDatabaseName() string {
	cDbName := C.getDatabaseName()
	//defer C.free(unsafe.Pointer(cDbName))
	return C.GoString(cDbName)
}

func getTableName() string {
	cTbName := C.getTableName()
	//defer C.free(unsafe.Pointer(cTbName))
	return C.GoString(cTbName)
}

func getIndexName() string {
	cIdName := C.getIndexName()
	//defer C.free(unsafe.Pointer(cIdName))
	return C.GoString(cIdName)
}

func getFieldTypes() []FieldType {
	var res []FieldType

	count := int(C.getFieldTypeNum())
	valueFlag := int(C.getFieldValueNum())
	if valueFlag != 0 {
		return res
	}
	for i := 0; i < count; i++ {
		var tmpFt FieldType
		ft := C.getFieldType(C.int(i))
		fts := strings.Fields(C.GoString(ft))
		C.free(unsafe.Pointer(ft))
		tmpFt.Field = fts[0]
		if len(fts) >= 2 {
			tmpFt.Type = fts[1]
		}
		res = append(res, tmpFt)
	}

	return res
}

func getFieldValues() []FieldValue {
	var res []FieldValue
	count := int(C.getFieldValueNum())
	sltFlag := int(C.getTableAliasNum())
	if sltFlag != 0 {
		return res
	}
	for i := 0; i < count; i++ {
		var tmpFv FieldValue
		fv := C.getFieldValue(C.int(i))
		fvs := strings.Fields(C.GoString(fv))
		C.free(unsafe.Pointer(fv))
		fvType, _ := strconv.Atoi(fvs[0])
		tmpFv.Field = fvs[1]
		switch fvType {
		case C.REDISQL_INT:
			tmpFv.Value, _ = strconv.Atoi(fvs[2])
			break
		case C.REDISQL_FLOAT:
			tmpFv.Value, _ = strconv.ParseFloat(fvs[2], 64)
			break
		case C.REDISQL_STRING:
			tmpFv.Value = strings.Trim(fvs[2], "'")
		}
		res = append(res, tmpFv)
	}
	return res
}

func getFieldAlias() []FieldAlias {
	var res []FieldAlias
	count := int(C.getFieldAliasNum())
	for i := 0; i < count; i++ {
		var tmpFa FieldAlias
		fa := C.getFieldAlias(C.int(i))
		fas := strings.Fields(C.GoString(fa))
		C.free(unsafe.Pointer(fa))
		tmpFa.TableAlias = fas[0]
		tmpFa.Field = fas[1]
		tmpFa.Alias = fas[2]
		res = append(res, tmpFa)
	}
	return res
}

func getTableAlias() []TableAlias {
	var res []TableAlias
	count := int(C.getTableAliasNum())
	for i := 0; i < count; i++ {
		var tmpTa TableAlias
		ta := C.getTableAlias(C.int(i))
		tas := strings.Fields(C.GoString(ta))
		C.free(unsafe.Pointer(ta))
		tmpTa.Table = tas[0]
		tmpTa.Alias = tas[1]
		res = append(res, tmpTa)
	}
	return res
}

func getWhere() string {
	cWhere := C.getWhere()
	//defer C.free(unsafe.Pointer(cWhere))
	return C.GoString(cWhere)
}

func getTop() int {
	return int(C.getTop())
}

func getLimit() Limit {
	var res Limit
	cLimit := C.getLimit()
	defer C.free(unsafe.Pointer(cLimit))
	limits := strings.Fields(C.GoString(cLimit))
	res.Start, _ = strconv.Atoi(limits[0])
	res.End, _ = strconv.Atoi(limits[1])
	return res
}

func GetSql(sql string) *RedisqlNode {
	C.redisql_parse(C.CString(sql))
	defer C.destorySqlNode()
	//C.showSql()
	return &RedisqlNode{
		Type:         getType(),
		DatabaseName: getDatabaseName(),
		TableName:    getTableName(),
		IndexName:    getIndexName(),
		FieldTypes:   getFieldTypes(),
		FieldValues:  getFieldValues(),
		FieldAliases: getFieldAlias(),
		TableAliases: getTableAlias(),
		Where:        getWhere(),
		Top:          getTop(),
		Limit:        getLimit(),
	}
}
