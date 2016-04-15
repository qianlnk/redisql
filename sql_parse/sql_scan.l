%{
#include "y.tab.h"
#include "stdio.h"
#include "stdlib.h"
#include "syntaxtree.h"

//#define PRINT_DEBUG

int create_keyword_node(char * pszKeyWordName);

int create_name_node(char * pszName);

int create_bound_node(char * pszName);

int create_int_const_node(int nVal);

int create_float_const_node(double fVal);

int create_string_const_node(char * strVal);

int create_operator_node(char * strOperator);

%}

%%

CREATE {
		if(!create_keyword_node("CREATE"))
		{
			return 0;
		}
		
		return CREATE;
	}

TABLE {
		if(!create_keyword_node("TABLE"))
		{
			return 0;
		}
		
		return TABLE;
	}

QUIT {
		if(!create_keyword_node("QUIT"))
		{
			return 0;
		}
		
		return QUIT;		
	}

INT    {
		if(!create_keyword_node("INT"))
		{
			return 0;
		}

		return INT;
	}

FLOAT {
		if(!create_keyword_node("FLOAT"))
		{
			return 0;
		}

		return FLOAT;
	}

STRING {
		if(!create_keyword_node("STRING"))
		{
			return 0;
		}

		return STRING;
	}
INSERT {
		if(!create_keyword_node("INSERT"))
		{
			return 0;
		}

		return INSERT;		
	}

INTO {
		if(!create_keyword_node("INTO"))
		{
			return 0;
		}

		return INTO;		
	}

VALUES {
		if(!create_keyword_node("VALUES"))
		{
			return 0;
		}

		return VALUES;		
	}
NULL {
		if(!create_keyword_node("NULLX"))
		{
			return 0;
		}

		return NULLX;
		
	}

DELETE {
		if(!create_keyword_node("DELETE"))
		{
			return 0;
		}

		return DELETE;		
	}

FROM {
		if(!create_keyword_node("FROM"))
		{
			return 0;
		}

		return FROM;		
	}
DROP {
		if(!create_keyword_node("DROP"))
		{
			return 0;
		}

		return DROP;		
	}

HELP {
		if(!create_keyword_node("HELP"))
		{
			return 0;
		}

		return HELP;
		
	}

AND {
		if(!create_keyword_node("AND"))
		{
			return 0;
		}

		return AND;		
	}

OR {
		if(!create_keyword_node("OR"))
		{
			return 0;
		}

		return OR;		
	}

NOT {
		if(!create_keyword_node("NOT"))
		{
			return 0;
		}

		return NOT;		
	}

WHERE {
		if(!create_keyword_node("WHERE"))
		{
			return 0;
		}

		return WHERE;		
	}

UPDATE {
		if(!create_keyword_node("UPDATE"))
		{
			return 0;
		}

		return UPDATE;		
	}

SET 	{
		if(!create_keyword_node("SET"))
		{
			return 0;
		}

		return SET;
	}

SELECT 	{
		if(!create_keyword_node("SELECT"))
		{
			return 0;
		}

		return SELECT;		
	}


[a-zA-Z][a-zA-Z0-9_]* {
				if(!create_name_node(yytext))
				{
					return 0;
				}

				return NAME;
		      }

[0-9]+ {
		if(!create_int_const_node(atoi(yytext)))
		{
			return 0;
		}	

		return INTVAL;
	}

[0-9]+"."[0-9]* {

		if(!create_float_const_node(atof(yytext)))
		{
			return 0;
		}

		return FLOATVAL;
	}

'[^'\n]*' {
		if(!create_string_const_node(yytext))
		{
			return 0;
		}

		return STRINGVAL;		
	}

[ \t\r\n]+ {
		/*忽略*/;
	}


[;] {
		create_bound_node(yytext);
		
		return 0;
	}

[(),.] {
		if(!create_bound_node(yytext))
		{
			return 0;
		}	
	
		return yytext[0];
	}

"=" |
"<>" |
"<" |
">" |
"<=" |
">=" {
		if(!create_operator_node(yytext))
		{
			return 0;
		}

		return COMPARISON;		
	}

"+" |
"-" |
"*" |
"/"	{
		if(!create_operator_node(yytext))
		{
			return 0;
		}

		return yytext[0];		
	}

. 	{
		printf("非法SQL字符串:%s\n", yytext);
		
		return 0;
	}

%%

//创建关键字结点

int create_keyword_node(char * pszKeyWordName)
{
	#ifdef PRINT_DEBUG
	printf("%s\n", yytext);
	#endif

	if(!pszKeyWordName)
	{
		printf("KeyWordName == NULL\n");

		return 0;
	}
	
	struct stnode * p = malloc_node();

	char * pBuf = NULL;

	if(!p)
	{
		printf("创建结点失败\n");

		return 0;
	}

	p->m_nType = KEYWORD;

	int nLen = strlen(pszKeyWordName);

	pBuf = 	(char * )malloc(nLen + 1);

	strcpy(pBuf, pszKeyWordName);

	pBuf[nLen] = 0;

	p->m_strName = pBuf;

	yylval.node_ptr = p;

	return 1;
}

//创建NAME

int create_name_node(char * pszName)
{
	#ifdef PRINT_DEBUG
	printf("%s\n", yytext);
	#endif

	if(!pszName)
	{
		return 0;
	}	
	
	struct stnode * p = malloc_node();

	char * pBuf = NULL;

	if(!p)
	{
		printf("创建结点失败\n");

		return 0;
	}

	//名称

	p->m_nType = ID;

	int nLen = strlen(pszName);

	pBuf = 	(char * )malloc(nLen + 1);

	strcpy(pBuf, pszName);

	pBuf[nLen] = 0;

	p->m_strName = pBuf;	

	yylval.node_ptr = p;

	return 1;
}

//创建界限符

int create_bound_node(char * pszName)
{
	#ifdef PRINT_DEBUG
	printf("%s\n", yytext);
	#endif

	if(!pszName)
	{
		return 0;
	}	
	
	struct stnode * p = malloc_node();

	char * pBuf = NULL;

	if(!p)
	{
		printf("创建结点失败\n");

		return 0;
	}

	//名称

	p->m_nType = BOUND_SYM;

	int nLen = strlen(pszName);

	pBuf = 	(char * )malloc(nLen + 1);

	strcpy(pBuf, pszName);

	pBuf[nLen] = 0;

	p->m_strName = pBuf;	

	yylval.node_ptr = p;

	return 1;
}

//创建整常数

int create_int_const_node(int nVal)
{
	#ifdef PRINT_DEBUG
	printf("%s\n", yytext);
	#endif

	struct stnode * p = malloc_node();

	if(!p)
	{
		printf("创建结点失败\n");

		return 0;
	}

	p->m_nType = INT_TYPE_C;

	p->m_Val.intval = nVal;

	yylval.node_ptr = p;

	return 1;
}

//创建小数

int create_float_const_node(double fVal)
{
	#ifdef PRINT_DEBUG
	printf("%s\n", yytext);
	#endif

	struct stnode * p = malloc_node();

	if(!p)
	{
		printf("创建结点失败\n");

		return 0;
	}

	p->m_nType = FLOAT_TYPE_C;

	p->m_Val.floatval = fVal;

	yylval.node_ptr = p;

	return 1;
}

//创建字符串

int create_string_const_node(char * strVal)
{
	#ifdef PRINT_DEBUG
	printf("%s\n", yytext);
	#endif

	struct stnode * p = malloc_node();

	if(!p)
	{
		printf("创建结点失败\n");

		return 0;
	}

	p->m_nType = STRING_TYPE_C;

	char * strTemp = (char *)malloc(strlen(strVal) + 1);

	if(!strTemp)
	{
		printf("create_string_const_node failed\n");

		return 0;
	}

	strcpy(strTemp, strVal);

	strTemp[strlen(strVal)] = 0;

	p->m_Val.strval = strTemp;

	yylval.node_ptr = p;

	return 1;
}

//创建运算符结点

int create_operator_node(char * strOperator)
{
	#ifdef PRINT_DEBUG
	printf("%s\n", yytext);
	#endif	

	if(!strOperator)
	{
		printf("strOperator == NULL\n");

		return 0;
	}	
	
	struct stnode * p = malloc_node();

	char * pBuf = NULL;

	if(!p)
	{
		printf("创建结点失败\n");

		return 0;
	}

	p->m_nType = OPERATOR;

	int nLen = strlen(strOperator);

	pBuf = 	(char * )malloc(nLen + 1);

	strcpy(pBuf, strOperator);

	pBuf[nLen] = 0;

	p->m_strName = pBuf;

	yylval.node_ptr = p;

	return 1;	
}


