#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "parse_redisql.h"

//golble list
stNode *gpstTree;
//golble action type
enum eActionType geAction;

//malloc
stNode * mallocNode()
{
	stNode *pstNew = (struct tag_stNode *)malloc(sizeof(struct tag_stNode));

	if (NULL == pstNew)
	{
		return NULL;
	}

	memset(pstNew, 0, sizeof(struct tag_stNode));

	return pstNew;
}
//free
void freeNode(stNode *pst)
{
	if (NULL != pst)
	{
		free(pst);
	}
}

//destory tree
void destoryTree(stNode *pstTree)
{
	if (NULL == pstTree)
	{
		return;
	}

	destoryTree(pstTree->pstChild);
	destoryTree(pstTree->pstBrother);
	free(pstTree);
}

//add node
int appendNode(stNode *pstParent, stNode *pstChild)
{
	return 0;
}