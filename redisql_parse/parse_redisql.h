#ifndef _PARSE_REDISQL_H_
#define _PARSE_REDISQL_H_

#include <stdio.h>
#include <stdlib.h>

#ifdef __CPLUSPLUS
extern "C" {
#endif

int redisql_parse(const char * sql);

enum eNodeType {
	KEYWORD = 0,
	ID,
	OPERATOR,
	INT_C,		//int float date all count as number
	FLOAT_C,
	DATE_C,
	STRING_C,
	BRANCH,
	BOUND_SYM
};

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

//parse as an node every word.
typedef struct tag_stNode{
	int nType;		//enum eNodeType
	char *pcName;	//name
	union{
		int nVal;
		char *pcVal;
		double fVal;
	}uVal;
	struct tag_stNode *pstChild;
	struct tag_stNode *pstBrother;
}stNode;

//golble list
extern stNode *gpstTree;
//golble action type
extern enum eActionType geAction;

//malloc
stNode * mallocNode();
//free
void freeNode(stNode *pst);
//destory tree
void destoryTree(stNode *pstTree);
//add node
int appendNode(stNode *pstParent, stNode *pstChild);

#ifdef __CPLUSPLUS
}
#endif
#endif//_PARSE_REDISQL_H_