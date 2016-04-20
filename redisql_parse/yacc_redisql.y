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

%token <strVal> USE SHOW DATABASES TABLES INDEX FROM DESC CREATE DATABASE TABLE NUMBER STRING DATE ON INSERT INTO VALUES SELECT WHERE AND OR LIKE TOP LIMIT HELP EXIT
%token <strVal> NAME STRINGVAL COMPARISON '(' ')' ',' '.' '+' '-' '*' '/' ';'
%token <nVal> INTVAL
%token <fVal> FLOATVAL
%type <strVal> use_database database table

%start sql

%%
sql:
	use_database
	{
		printf("3use %s\n", $1);
		setType(REDISQL_USE);
		setDatabaseName($1);
	}
	| SHOW DATABASES ';'
	{
		setType(REDISQL_SHOW_DATABASES);
	}
	| SHOW TABLES ';'
	{
		setType(REDISQL_SHOW_TABLES);
	}
	|SHOW INDEX FROM table ';'
	{
		setType(REDISQL_SHOW_INDEX);
		setTableName($4);
	}
	;

use_database:
	USE database ';'
	{
		printf("2use %s\n", $2);
		$$ = $2;
	}
	;

database:
	NAME
	{
		printf("1use %s\n", $1);
		$$ = $1;
	}
	;

table:
	NAME
	{
		$$ = $1;
	}
%%
#include "lex.yy.c"
void yyerror(const char * s)
{
	printf("%s\n", s);
}

int redisql_parse(const char * sql)
{
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