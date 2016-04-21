%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "parse_redisql.h"

extern int yylex(void);
void yyerror(const char *s);
%}

%union{
	char *key;
	char *strVal;
	int nVal;
	double fVal;
}

%token <key> USE SHOW DATABASES TABLES INDEX FROM DESC CREATE DATABASE TABLE NUMBER STRING DATE ON INSERT INTO VALUES SELECT AS WHERE AND OR LIKE TOP LIMIT HELP EXIT
%token <strVal> NAME STRINGVAL COMPARISON '(' ')' ',' '.' '+' '-' '*' '/' ';'
%token <nVal> INTVAL
%token <fVal> FLOATVAL
%type <strVal> conlumn_def_list  conlumn_name conlumn_type index_name value_list select_conlumn_list select_conlumn select_table_list select_table where_condition table_alias conlumn_alias condition bool_term bool_op expression primary

%type <strVal> use database_name desc table_name
%type <strVal> show show_databases show_tables show_index
%type <strVal> create create_database create_table create_index
/*%type <strVal> insert_into
%type <strVal> select
%type <strVal> drop drop_database drop_table drop_index
%type <strVal> delete_from
*/
%type <strVal> opt_semicolon

%start sql

%%
sql:
	use
	{
		setType(REDISQL_USE);
		setDatabaseName($1);
	}
	| show 
	{
		//no code
	}
	| desc
	{
		setType(REDISQL_DESC);
		setTableName($1);
	}
	| create
	{
		//no code
	}
	| INSERT INTO table_name '(' conlumn_def_list ')' VALUES '(' value_list ')' opt_semicolon
	{
		setType(REDISQL_INSERT);
		setTableName($3);
	}
	| SELECT select_conlumn_list FROM select_table_list where_condition opt_semicolon
	{
		setType(REDISQL_SELECT);
	}
	;

use:
	USE database_name opt_semicolon
	{
		$$ = $2;
	}
	;

show:
	show_databases
	{
		setType(REDISQL_SHOW_DATABASES);
	}
	| show_tables
	{
		setType(REDISQL_SHOW_TABLES);
	}
	| show_index
	{
		setType(REDISQL_SHOW_INDEX);
		setIndexName($1);
	}
	;

show_databases:
	SHOW DATABASES opt_semicolon
	{
		$$ = NULL;
	}
	;

show_tables:
	SHOW TABLES opt_semicolon
	{
		$$ = NULL;
	}
	;

show_index:
	SHOW INDEX FROM table_name opt_semicolon
	{
		$$ = $4;
	}
	;

desc:
	DESC table_name opt_semicolon
	{
		$$ = $2;
	}
	;

create:
	create_database
	{
		setType(REDISQL_CREATE_DATABASE);
		setDatabaseName($1);
	}
	| create_table
	{
		setType(REDISQL_CREATE_TABLE);
		setTableName($1);
	}
	| create_index
	{
		setType(REDISQL_CREATE_INDEX);
	}
	;

create_database:
	CREATE DATABASE database_name opt_semicolon
	{
		$$ = $3;
	}
	;

create_table:
	CREATE TABLE table_name '(' conlumn_def_list ')' opt_semicolon
	{
		$$ = $3;
	}
	;

create_index:
	CREATE INDEX index_name ON table_name '(' conlumn_def_list ')' opt_semicolon
	{
		setIndexName($3);
		setTableName($5);
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

opt_semicolon:
	';'
	{
		$$ = $1;
	}
	|/*empty*/
	{
		$$ = "";
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