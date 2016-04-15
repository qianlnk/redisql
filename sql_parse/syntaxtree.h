#include "stdlib.h"
#include "stdio.h"

#ifndef SYNTAXTREE_H
#define SYNTAXTREE_H

#ifdef __CPLUSPLUS
extern "C" {
#endif

enum sql_node_type{BRANCH=0, KEYWORD, OPERATOR, ID, INT_TYPE_C, FLOAT_TYPE_C, STRING_TYPE_C, BOUND_SYM};

enum sql_action_type {SAT_SELECT=0, SAT_INSERT, SAT_UPDATE, SAT_DELETE, SAT_CREATE_TABLE, SAT_QUIT, SAT_DROP_TABLE, SAT_HELP};

struct stnode
{
	//结点类型

	int m_nType;

	//结点名称

	char * m_strName;

	//结点值

	union
	{
		int intval;

		char * strval;

		double floatval;
	} m_Val;

	struct stnode * m_pChildList;

	struct stnode * m_pBrotherList;
};

//话法树

extern struct stnode * syntax_tree_ptr;

//SQL语句动作类型

extern enum sql_action_type sql_action;

//分配结点

struct stnode * malloc_node();

//释放结点

void free_node(struct stnode * p);

//添加孩子

int append_child(struct stnode * pParent, struct stnode * pChild);

//打印树结点

void print_tree_node(struct stnode * pParent);

//打印树

void print_syntax_tree(struct stnode * pParent);

//销毁语法树

void destroy_syntax_tree(struct stnode * pParent);

//得到孩子

struct stnode * get_child(struct stnode * pParent, int nOrderNo/*从1开始*/);

//得到孩子数量

int get_child_cnt(struct stnode * pParent);

#ifdef __CPLUSPLUS
}
#endif

#endif //SYNTAXTREE_H


