#ifndef _PARSE_REDISQL_H_
#define _PARSE_REDISQL_H_

#include <stdio.h>
#include <stdlib.h>

#ifdef __CPLUSPLUS
extern "C" {
#endif

int redisql_parse(const char * sql);

enum eActionType {
	REDISQL_USE = 0,			//use databasename;
	REDISQL_SHOW_DATABASES,		//show databases; 
	REDISQL_SHOW_TABLES, 		//show tables; 
	REDISQL_SHOW_INDEX,			//show index from tablename;
	REDISQL_DESC,				//desc tablename;
	REDISQL_CREATE_DATABASE,	//create database databasename; 
	REDISQL_CREATE_TABLE,		//create table tbname(field1...); 
	REDISQL_CREATE_INDEX,		//create index indexname on tablename(fieldname);
	REDISQL_INSERT,				//insert into tablename(field1...) values(value1...);
	REDISQL_SELECT,				//select field1... from table1... where case1... limit start end
	REDISQL_UPDATE,				//update tablename set field1=value1 where case1...
	REDISQL_DELECT,				//delete from tablename where case1...
	REDISQL_DROP_DATABASE,		//drop database databasename
	REDISQL_DROP_TABLE, 		//drop table tablename
	REDISQL_EXIT,				//exit
	REDISQL_HELP				//help
};

//field type node
typedef struct tag_FieldType{
	char *pcField;
	char *pcType;
	struct tag_FieldType *pstNextField;
}FieldType;

//field value node
typedef struct tag_FieldValue{
	char *pcField;
	union{
		int nValue;
		char *pcValue;
		double fValue;
	}uValue;
	struct tag_FieldValue *pstNextField;
}FieldValue;

//limit
typedef struct tag_Limit{
	int nStart;
	int nEnd;
}Limit;

//parse sql as a struct
typedef struct tag_SqlNode{
	int nType;
	char *pcName;				//databasename, tablename, indexname
	FieldType *pstFieldType;
	FieldValue *pstFieldValue;
	char *pcWhere;
	int nTop;
	Limit stLimit;
}SqlNode;

//golble list
extern SqlNode g_stSql;

//free
void freeSqlNode();
//destory tree
void destorySqlNode();

#ifdef __CPLUSPLUS
}
#endif
#endif//_PARSE_REDISQL_H_