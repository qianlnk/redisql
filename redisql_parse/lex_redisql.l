%{
#include <stdio.h>
#include <stdlib.h>
#include "y.tab.h"
#include "parse_redisql.h"
#define DEBUG 1

%}

digit ([0-9])
character ([a-zA-Z])
integer ({digit}+)

%%

USE |
use {
	printf("lex use %s\n", yytext);
	return USE;
}

SHOW |
show {
	return SHOW;
}

DATABASES |
databases {
	return DATABASES;
}

TABLES |
tables {
	return TABLES;
}

INDEX |
index {
	return INDEX;
}

FROM |
from {
	return FROM;
}

DESC |
desc {
	return DESC;
}

CREATE |
create {
	return CREATE;
}

DATABASE |
database {
	return DATABASE;
}

TABLE |
table {
	return TABLE;
}

NUMBER |
number {
	return NUMBER;
}

STRING |
string {
	return STRING;
}

DATE |
date {
	return DATE;
}

ON |
on {
	return ON;
}

INSERT |
insert {
	return INSERT;
}

INTO |
into {
	return INTO;
}

VALUES |
values {
	return VALUES;
}

SELECT |
select {
	return SELECT;
}

AS |
as {
	return AS;
}

WHERE |
where {
	return WHERE;
}

AND |
and {
	printf("lex and %s\n", yytext);
	return AND;
}

OR |
or {
	printf("lex or %s\n", yytext);
	return OR;
}

LIKE |
like {
	return LIKE;
}

TOP |
top {
	return TOP;
}

LIMIT |
limit {
	return LIMIT;
}

[a-zA-Z][a-zA-Z0-9_]* {
	printf("lex name %s\n", yytext);
	yylval.strVal = strdup(yytext);
	return NAME;
}

[0-9]+ {	
	printf("lex intval %s\n", yytext);
		yylval.nVal = atoi(yytext);
		return INTVAL;
	}

[0-9]+"."[0-9]* {
	printf("lex floatval %s\n", yytext);
		yylval.fVal = atof(yytext);
		return FLOATVAL;
	}

'[^'\n]*' {
	printf("lex stringval %s\n", yytext);
		yylval.strVal = strdup(yytext);
		return STRINGVAL;		
	}

[ \t\r\n]+ {
		/*empty*/;
	}


[;] {	
		return yytext[0];
	}

[(),.] {
		return yytext[0];
	}

"=" |
"<>" |
"<" |
">" |
"<=" |
">=" {
		yylval.strVal = strdup(yytext);
		return COMPARISON;		
	}

"+" |
"-" |
"*" |
"/"	{
		return yytext[0];		
	}

%%

int yywrap(void)
{
  printf("done\n");
  return 1;
}