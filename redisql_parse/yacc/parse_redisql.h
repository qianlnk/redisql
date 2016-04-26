#ifndef _PARSE_REDISQL_H_
#define _PARSE_REDISQL_H_

#include <stdio.h>
#include <stdlib.h>

#define MALLOC_ERR 	1
#define TYPE_ERR 	2

#ifdef __CPLUSPLUS
extern "C" {
#endif

int redisql_parse(const char * sql);

enum eFieldType{
	REDISQL_INT = 0,
	REDISQL_STRING,
	REDISQL_FLOAT
};

enum eActionType {
	REDISQL_USE = 0,			//use databasename;
	REDISQL_SHOW_DATABASES,		//show databases; 
	REDISQL_SHOW_TABLES, 		//show tables; 
	REDISQL_SHOW_INDEX,			//show index from tablename;
	REDISQL_DESC,				//desc tablename;
	REDISQL_CREATE_DATABASE,	//create database databasename; 
	REDISQL_CREATE_TABLE,		//create table tbname(field1,type1...); 
	REDISQL_CREATE_INDEX,		//create index indexname on tablename(fieldname);
	REDISQL_INSERT,				//insert into tablename(field1...) values(value1...);
	REDISQL_SELECT,				//select field1... from table1... where case1... limit start end
	REDISQL_UPDATE,				//update tablename set field1=value1 where case1...
	REDISQL_DELETE,				//delete from tablename where case1...
	REDISQL_DROP_DATABASE,		//drop database databasename
	REDISQL_DROP_TABLE, 		//drop table tablename
	REDISQL_EXIT,				//exit
	REDISQL_HELP,				//help
	REDISQL_EMPTY
};

//field type node
typedef struct tag_FieldType{
	char *pcField;
	char *pcType;
	struct tag_FieldType *pstNextField;
}FieldType;

//value
union Value{
	int nValue;
	char *pcValue;
	double fValue;
};

//field value node
typedef struct tag_FieldValue{
	int nFieldType;
	union Value uValue;
	struct tag_FieldValue *pstNextField;
}FieldValue;

//limit
typedef struct tag_Limit{
	int nStart;
	int nEnd;
}Limit;

typedef struct tag_FieldAlias{
	char *pcTableAlias;
	char *pcField;
	char *pcAlias;
	struct tag_FieldAlias *pstNextField;
}FieldAlias;

typedef struct tag_TableAlias{
	char *pcTable;
	char *pcAlias;
	struct tag_TableAlias *pstNextTable;
}TableAlias;

//parse sql as a struct
typedef struct tag_SqlNode{
	int nType;					//action type
	char *pcDatabaseName;		//database name
	char *pcTableName;			//table name
	char *pcIndexName;			//index name
	int nFieldTypeNum;			//
	FieldType *pstFieldType;		//create table's fields and types
	int nFieldValueNum;
	FieldValue *pstFieldValue;	//insert
	int nFieldAliasNum;
	FieldAlias *pstFieldAlias;	//select
	int nTableAliasNum;
	TableAlias *pstFrom;			//select from
	char *pcWhere;				//parse every word split with " ".
	int nTop;					//select top
	Limit stLimit;				//select limit
}SqlNode;

//golble list
extern SqlNode g_stSql;

void init();
void setType(int nType);
int setDatabaseName(const char * pcDatabaseName);
int setTableName(const char * pcTableName);
int setIndexName(const char * pcIndexName);
int addFieldType(const char * pcField, const char * pcType);
int addFieldValue(int nFieldType, union Value uValue);
int addFieldAlias(const char * pcTableAlias, const char * pcField, const char * pcAlias);
int addTableAlias(const char * pcTable, const char * pcAlias);
int setWhere(const char * pcWhere);
void setTop(int nTop);
void setLimit(int nStart, int nEnd);

int getType();
char * getDatabaseName();
char * getTableName();
char * getIndexName();
int getFieldTypeNum();	
int getFieldValueNum();
int getFieldAliasNum();
int getTableAliasNum();
char * getFieldType(int sn);	//field and type slpit with " "
char * getFieldValue(int sn);
char * getFieldAlias(int sn);
char * getTableAlias(int sn);
char * getWhere();
int getTop();
char * getLimit();

FieldType * mallocFieldType();
FieldValue * mallocFieldValue();
FieldAlias * mallocFieldAlias();
TableAlias * mallocTableAlias();

void destoryFieldType(FieldType *pst);
void destoryFieldValue(FieldValue *pst);
void destoryFieldAlias(FieldAlias *pst);
void destoryTableAlias(TableAlias *pst);

//destory
void destorySqlNode();

void showSql();
#ifdef __CPLUSPLUS
}
#endif
#endif//_PARSE_REDISQL_H_