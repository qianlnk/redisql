%{
#include <stdio.h>
#include <stdlib.h>
#include "parse_redisql.h"
#define DEBUG 1
int createKeywordNode(char *pcKeywordName);
%}

digit ([0-9])
character ([a-zA-Z])
integer ({digit}+)

%%

//keyword regular
USE |
use {

}

%%

int createKeywordNode(char *pcKeywordName)
{
	stNode *pstNewNode = NULL;
	char *pcName = NULL;

	#if DEBUG
		printf("createKeywordNode: yytext = %s\n", yytext);
	#endif

	if (NULL == pcKeywordName)
	{
		printf("createKeywordNode: pcKeywordName is null.\n");
		return 1;
	}

	pstNewNode = mallocNode();
	if (NULL == pstNewNode) 
	{
		printf("createKeywordNode: malloc node failed.\n");
	}

	pcName = (char *)malloc(strlen(pcKeywordName) + 1);
	if (NULL == pcName)
	{
		printf("createKeywordNode: malloc string failed.\n");
	}

	strcpy(pcName, pcKeywordName);

	pstNewNode->nType = KEYWORD;
	pstNewNode->pcName = pcName;


}

