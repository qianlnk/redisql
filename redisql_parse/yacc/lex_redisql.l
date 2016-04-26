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

INT |
int {
	return INT;
}

FLOAT |
float {

}

DOUBLE |
double {
	return DOUBLE;
}

CHAR |
char {
	return CHAR;
} 

VARCHAR |
varchar {
	return VARCHAR;
}

TEXT |
text {
	return TEXT;
} 

DATE |
date {
	return DATE;
} 

DATETIME |
datetime {
	return DATETIME;
} 

NOT |
not {
	return NOT;
}

NULL |
null {
	return TOKEN_NULL;
}

UNIQUE |
unique {
	return UNIQUE;
} 

PRIMARY |
primary {
	return PRIMARY;
} 

KEY |
key {
	return KEY;
} 

DEFAULT |
default {
	return DEFAULT;
} 

AUTO_INCREMENT |
auto_increment {
	return AUTO_INCREMENT;
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

DISTINCT |
distinct {
	return DISTINCT;
}

COUNT |
count {
	return COUNT;
}

SUM |
sum {
	return SUM;
}

AVG |
avg {
	return AVG;
}

MIN |
min {
	return MIN;
}

MAX |
max {
	return MAX;
}

JOIN |
join {
	return JOIN;
}

CROSS |
cross {
	return CROSS;
}

INNER |
inner {
	return INNER;
}

LEFT |
left {
	return LEFT;
}

RIGHT |
right {
	return RIGHT;
}

CONCAT |
concat {
	return CONCAT;
}

FULL |
full {
	return FULL;
}

NATURAL |
natural {
	return NATURAL;
}

OUTER |
outer {
	return OUTER;
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
	return AND;
}

OR |
or {
	return OR;
}

LIKE |
like {
	yylval.strVal = strdup("LIKE");
	return COMPARISON;
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
	yylval.strVal = strdup(yytext);
	return NAME;
}

[0-9]+ {	
		yylval.nVal = atoi(yytext);
		return INTVAL;
	}

[0-9]+"."[0-9]* {
		yylval.fVal = atof(yytext);
		return FLOATVAL;
	}

'[^'\n]*' {
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
  return 1;
}