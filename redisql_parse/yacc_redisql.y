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
	char *strDateType;
	char *strDatabase;
	char *strTable;
	char *strColumn;
	char *strIndex;
	char *strOption;
	char *strVal;
	int nVal;
	double fVal;
	FieldAlias stFieldAlias;
}

%token <key> USE SHOW DATABASES TABLES INDEX FROM DESC CREATE DATABASE TABLE ON INSERT INTO VALUES SELECT DISTINCT AS WHERE AND OR LIKE TOP LIMIT HELP EXIT NOT TOKEN_NULL UNIQUE PRIMARY KEY DEFAULT AUTO_INCREMENT COUNT SUM AVG MIN MAX JOIN CROSS INNER LEFT RIGHT FULL NATURAL OUTER CONCAT
%token <strDateType> INT FLOAT DOUBLE CHAR VARCHAR TEXT DATE DATETIME 
%token <strVal> NAME STRINGVAL COMPARISON '(' ')' ',' '.' '+' '-' '*' '/' ';'
%token <nVal> INTVAL
%token <fVal> FLOATVAL
%type <strVal> bool_term bool_op  

%type <strOption> use desc 
%type <strOption> show show_databases show_tables show_index
%type <strOption> create create_database create_table create_index 
%type <strOption> insert_into
%type <strOption> select
/*%type <strOption> drop drop_database drop_table drop_index
%type <strOption> delete_from
*/
%type <strDatabase> database_name
%type <strTable> table_name
%type <strColumn> column_name_type_list column_name_type column_name column_type opt_column_name_list column_name_list opt_constraint_list constraint_list constraint
%type <strVal> value_list value
%type <strIndex> index_name
%type <strVal> opt_semicolon opt_distinct expression_list opt_alias column_name_or_star  function_name table_list table_def default_join join opt_outer opt_join_condition join_condition condition
%type <stFieldAlias> column_reference expression mulexp primary term  

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
	| insert_into
	{
		setType(REDISQL_INSERT);
	}
	| select
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
	CREATE TABLE table_name '(' column_name_type_list ')' opt_semicolon
	{
		$$ = $3;
	}
	;

column_name_type_list:
	column_name_type
	{

	}
	| column_name_type_list ',' column_name_type
	{

	}
	;

column_name_type:
	column_name column_type opt_constraint_list
	{
		addFieldType($1, $2);
	}
	;

column_type:
	INT
	{
		$$ = "NUMBER";
	}
	| FLOAT
	{
		$$ = "NUMBER";
	}
	| DOUBLE
	{
		$$ = "NUMBER";
	}
	| CHAR
	{
		$$ = "STRING";
	}
	| VARCHAR
	{
		$$ = "STRING";
	}
	| TEXT
	{
		$$ = "STRING";
	}
	| DATE
	{
		$$ = "DATE";
	}
	| DATETIME
	{
		$$ = "DATE";
	}
	| column_type '(' INTVAL ')'
	{
		$$ = $1;
	}
	;

opt_constraint_list:
	constraint_list {}
	|/*empty*/ 		{}
	;

constraint_list:
	constraint 						{}
	| constraint constraint_list 	{}
	;

constraint:
	NOT TOKEN_NULL		{}
	| UNIQUE			{}
	| PRIMARY KEY		{}
	| DEFAULT value 	{}
	| AUTO_INCREMENT	{}
	;

create_index:
	CREATE INDEX index_name ON table_name '(' column_name_list ')' opt_semicolon
	{
		setIndexName($3);
		setTableName($5);
	}
	;

insert_into:
	INSERT INTO table_name opt_column_name_list VALUES '(' value_list ')' opt_semicolon
	{
		setTableName($3);
	}
	;

opt_column_name_list:
	'(' column_name_list ')'
	{

	}
	|/*empty*/
	{

	}
	;

column_name_list:
	column_name
	{
		addFieldType($1, "");
	}
	| column_name_list ',' column_name
	{
		addFieldType($3, "");
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
		char pc[1000] = {'\0'};
		union Value uVal;
		uVal.nValue = $1;
		addFieldValue(REDISQL_INT, uVal);
		sprintf(pc, "%d", $1);
		$$ = pc;
	}
	| FLOATVAL
	{
		char pc[1000] = {'\0'};
		union Value uVal;
		uVal.fValue = $1;
		addFieldValue(REDISQL_FLOAT, uVal);
		sprintf(pc, "%f", $1);
		$$ = pc;
	}
	| STRINGVAL
	{
		union Value uVal;
		uVal.pcValue = $1;
		addFieldValue(REDISQL_STRING, uVal);
		$$ = $1;
	}
	;

select:
	SELECT opt_distinct expression_list FROM table_list opt_where_condition opt_semicolon
	{
		setType(REDISQL_SELECT);
	}
	;

opt_distinct:
	DISTINCT 	{}
	|/*empty*/	{}

expression_list:
	expression opt_alias
	{
	printf("a\n");
		FieldAlias st;
		st = $1;
		printf("st %s %s %s\n", st.pcTableAlias, st.pcField, $2);
		addFieldAlias(st.pcTableAlias, st.pcField, $2);
	}
	| expression_list ',' expression opt_alias
	{
	printf("b\n");
		FieldAlias st;
		st = $3;
		addFieldAlias(st.pcTableAlias, st.pcField, $4);
	}
	;

opt_alias:
	AS NAME
	{
		$$ = $2;
	}
	| NAME
	{
		$$ = $1;
	}
	|/*empty*/
	{
		$$ = "";
	}
	;

expression:
	expression '+' mulexp
	| expression '-' mulexp
	| mulexp
	{
	printf("c\n");
		$$ = $1;
	}
	;

mulexp: 
	mulexp '*' primary 	
	{}
	| mulexp '/' primary 
	{}
	| mulexp CONCAT primary 
	{}
	| primary 					
	{ 
	printf("d\n");
		$$ = $1; 
	}
	;

primary:
	'(' expression ')' 	
	{ 
	printf("e\n");
		$$ = $2; 
	}
	| '-' primary 
	{ 
		$$ = $2; 
	}
	| term 
	{ 
	printf("f\n");
		$$ = $1; 
	} 
	;

term: 
	value
	{
		printf("value:%s\n", $1);
		FieldAlias st;
		st.pcField = $1;
		$$ = st;
	}
	| TOKEN_NULL
	{

	}
	| column_reference
	{
		printf("column_reference\n");
		$$ = $1;
		printf("column_reference end\n");
	}
	| function_name '(' expression ')'
	{

	}
	;

column_reference:
	column_name_or_star
	{
	printf("g\n");
		FieldAlias stTmp;
		stTmp.pcTableAlias = "";
		stTmp.pcField = $1;
		stTmp.pcAlias = $1;
		$$ = stTmp;
	}
	| table_name '.' column_name_or_star
	{
	printf("h\n");
		FieldAlias stTmp;
		stTmp.pcTableAlias = $1;
		stTmp.pcField = $3;
		stTmp.pcAlias = $3;
		$$ = stTmp;
	}
	;

function_name:
	COUNT {}
	| SUM {}
	| AVG {}
	| MIN {}
	| MAX {}
	;

column_name_or_star:
	'*' 
	{ 
		$$ = "*";
	}
	| column_name
	{
		$$ = $1;
	}
	;

table_list:
	table_def 
	{}
	| table_list default_join table_def opt_join_condition
	{}
	| join table_def opt_join_condition
	{}
	;

table_def:
	table_name opt_alias
	{
		addTableAlias($1, $2);
	}
	;

default_join:
	',' 
	| JOIN 			{}
	| CROSS JOIN 	{}
	| INNER JOIN 	{}
	;

join:
	LEFT opt_outer JOIN 	{}
	| RIGHT opt_outer JOIN 	{}
	| FULL opt_outer JOIN 	{}
	| NATURAL JOIN 			{}
	;

opt_outer:
	OUTER 			{}
	| /* empty */	{}
	;

opt_join_condition:
	join_condition
	{

	}
	|/*empty*/
	{

	}
	;

join_condition:
	ON condition
	{

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

column_name:
	NAME
	{
		$$ = $1;
	}
	;

index_name:
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

opt_where_condition:
	WHERE condition
	{
		printf("f %s\n", acWhere);
	}
	| 
	{
	}
	;

condition: 
	bool_term 
	{ 
		printf("bool_term %s\n", $1);
		if (acWhere[0] == '\0')
		{
			strcat(acWhere, $1);
		}
		$$ = $1;
	}
	| condition bool_op bool_term
   	{ 
   		//printf("bool_term bool_op condition  %s %s %s\n", $1, $2, $3);
   		char pc[1000];
   		sprintf(pc, "%s %s %s", $1, $2, $3);
   		strcat(acWhere, " ");
   		strcat(acWhere, $2);
   		strcat(acWhere, " ");
   		strcat(acWhere, $3);
   		$$ = pc;
   	}
	;

bool_term: 
	expression COMPARISON expression 
   	{
   		printf("123\n");
   		FieldAlias st1, st3;
   		char pc[1000] = {'\0'};
   		st1 = $1;
   		st3 = $3;
   		if (strcmp(st1.pcTableAlias, "") != 0)
   		{
   			strcat(pc, st1.pcTableAlias);
   			strcat(pc, ".");
   		}
   		strcat(pc, st1.pcField);
   		strcat(pc, " ");
   		strcat(pc, $2);
   		strcat(pc, " ");
   		strcat(pc, st3.pcField);
   		printf("123 end %s\n", pc);
   		$$ = pc;
   	}
   	| '(' condition ')' 	
	{ 
		printf("456 start %s\n", $2);
		char pc[1000] = {'\0'};
		strcat(pc, "(");
   		strcat(pc, " ");
   		strcat(pc, $2);
   		strcat(pc, " ");
   		strcat(pc, ")");
   		printf("456 end %s\n", pc);
   		$$ = pc;
	}
	;

bool_op
	: AND { $$ = "AND"; } 
	| OR { $$ = "OR"; }
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