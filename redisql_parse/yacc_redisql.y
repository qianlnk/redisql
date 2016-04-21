%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "parse_redisql.h"

extern int yylex(void);
void yyerror(const char *s);
%}

%union{
	char *strVal;
	int nVal;
	double fVal;
}

%token <strVal> USE SHOW DATABASES TABLES INDEX FROM DESC CREATE DATABASE TABLE NUMBER STRING DATE ON INSERT INTO VALUES SELECT AS WHERE AND OR LIKE TOP LIMIT HELP EXIT
%token <strVal> NAME STRINGVAL COMPARISON '(' ')' ',' '.' '+' '-' '*' '/' ';'
%token <nVal> INTVAL
%token <fVal> FLOATVAL
%type <strVal> use_database database_name table_name semicolon_empty conlumn_def_list  conlumn_name conlumn_type index_name value_list select_conlumn_list select_conlumn select_table_list select_table where_condition table_alias conlumn_alias condition bool_term bool_op expression primary

%start sql

%%
sql:
	use_database
	{
		setType(REDISQL_USE);
		setDatabaseName($1);
	}
	| SHOW DATABASES semicolon_empty
	{
		setType(REDISQL_SHOW_DATABASES);
	}
	| SHOW TABLES semicolon_empty
	{
		setType(REDISQL_SHOW_TABLES);
	}
	| SHOW INDEX FROM table_name semicolon_empty
	{
		setType(REDISQL_SHOW_INDEX);
		setTableName($4);
	}
	| DESC table_name semicolon_empty
	{
		setType(REDISQL_DESC);
		setTableName($2);
	}
	| CREATE DATABASE database_name semicolon_empty
	{
		setType(REDISQL_CREATE_DATABASE);
		setDatabaseName($3);
	}
	| CREATE TABLE table_name '(' conlumn_def_list ')' semicolon_empty
	{
		setType(REDISQL_CREATE_TABLE);
		setTableName($3);
	}
	| CREATE INDEX index_name ON table_name '(' conlumn_def_list ')' semicolon_empty
	{
		setType(REDISQL_CREATE_INDEX);
		setIndexName($3);
		setTableName($5);
	}
	| INSERT INTO table_name '(' conlumn_def_list ')' VALUES '(' value_list ')' semicolon_empty
	{
		setType(REDISQL_INSERT);
		setTableName($3);
	}
	| SELECT select_conlumn_list FROM select_table_list where_condition semicolon_empty
	{
		setType(REDISQL_SELECT);
	}
	;

use_database:
	USE database_name semicolon_empty
	{
		$$ = $2;
	}
	;

database_name:
	NAME
	{
		$$ = $1;
	}
	;

table_name:
	NAME
	{
		$$ = $1;
	}
	;

semicolon_empty:
	';'
	{
		$$ = $1;
	}
	|
	{

	}
	;

conlumn_def_list:
	conlumn_def
	{

	}
	| conlumn_def_list ',' conlumn_def
	{

	}
	;

conlumn_def:
	conlumn_name conlumn_type
	{
		printf("%s %s\n", $1, $2);
		addFieldType($1, $2);
	}
	;

conlumn_name:
	NAME
	{
		$$ = $1;
	}
	;

conlumn_type:
	NUMBER
	{
		$$ = "NUMBER";
	}
	| STRING
	{
		$$ = "STRING";
	}
	| DATE
	{
		$$ = "DATE";
	}
	|
	{
		$$ = "";
	}
	;

index_name:
	NAME
	{
		$$ = $1;
	}
	;

value_list:
	value
	{

	}
	| value_list ',' value
	{

	}
	;

value:
	INTVAL
	{
		union Value uVal;
		uVal.nValue = $1;
		addFieldValue(REDISQL_INT, uVal);
	}
	| FLOATVAL
	{
		union Value uVal;
		uVal.fValue = $1;
		addFieldValue(REDISQL_FLOAT, uVal);
	}
	| STRINGVAL
	{
		union Value uVal;
		uVal.pcValue = $1;
		addFieldValue(REDISQL_STRING, uVal);
	}
	;

select_conlumn_list:
	select_conlumn
	{

	}
	| select_conlumn_list ',' select_conlumn
	{

	}
	;

table_alias:
	NAME
	{
		$$ = $1;
	}
	;

conlumn_alias:
	NAME
	{
		$$ = $1;
	}
	;

select_conlumn:
	table_alias '.' conlumn_name
	{
		addFieldAlias($1, $2, $2);
	}
	| table_alias '.' conlumn_name conlumn_alias
	{
		addFieldAlias($1, $2, $3);
	}
	| table_alias '.' conlumn_name AS conlumn_alias
	{
		addFieldAlias($1, $2, $3);
	}
	;

select_table_list:
	select_table
	{

	}
	| select_table_list ',' select_table
	{

	}
	;

select_table:
	table_name
	{
		addTableAlias($1, $1);
	}
	| table_name table_alias
	{
		addTableAlias($1, $2);
	}
	;

where_condition:
	WHERE condition
	{
		printf("f %s\n", $2);
	}
	| 
	{
	}
	;

condition: 
	bool_term 
	{ 
	}
	| bool_term bool_op condition 
   	{ 

   	}
	;

bool_term: 
	expression COMPARISON expression 
   	{

   	}
	| '(' condition ')' 	
	{ 
		$$ = $2; 
	}
	;

bool_op
	: AND { $$ = "AND"; } 
	| OR { $$ = "OR"; }
	;


expression:
	primary
	{

	}
	;

primary:
	'(' expression ')'
	{}
	|
	table_alias '.' conlumn_name
	{

	}
	| value_list
	{

	}
	;
%%
#include "lex.yy.c"
void yyerror(const char * s)
{
	printf("%s\n", s);
}

int redisql_parse(const char * sql)
{
	init();
	if(!sql)
	{
		printf("sql is null\n");
	}
        printf("sql = %s\n", sql);
	int len = strlen(sql);
       printf("2\n");
	YY_BUFFER_STATE state = yy_scan_string(sql);
        printf("3\n");
	yy_switch_to_buffer(state);
        printf("4\n");
	int n = yyparse();
        printf("5\n");
	yy_delete_buffer(state);
        printf("6\n");
	return n;
}