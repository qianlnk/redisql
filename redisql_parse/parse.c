#include <stdio.h>
#include <string.h>
#include "parse_redisql.h"

int main()
{
	char *sql = "show index from student;";
	redisql_parse(sql);
	showSql();
	destorySqlNode();
	return 0;
}
