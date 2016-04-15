#include "sql_parse.h"
#include "syntaxtree.h"
#include <stdio.h>
int main()
{
	char *sql = "select name, age, class FROM student WHERE name = '1231asdad';";
	printf("1\n");
    sql_parse(sql);
	printf("2\n");
	if (!syntax_tree_ptr)
	{
		printf("syntax_tree_ptr is null\n");
	}
	
	printf("sql_action = %d\n", sql_action);
	print_syntax_tree(syntax_tree_ptr);
	return 0;
}
