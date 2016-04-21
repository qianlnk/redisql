#include <stdio.h>
#include <string.h>
#include "parse_redisql.h"

int main()
{
	char sql[1000];
	while(1)
	{
		printf("redisql> ");
		gets(sql);
		redisql_parse(sql);
		showSql();
		destorySqlNode();
	}
	return 0;
}
