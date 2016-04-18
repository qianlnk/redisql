sql_parse: redisql_parse.o y.tab.o syntaxtree.o
	gcc redisql_parse.o y.tab.o syntaxtree.o -o sql_parse
redisql_parse.o: redisql_parse.c
	gcc -c redisql_parse.c
lex.yy.o: lex.yy.c y.tab.h
	gcc -c lex.yy.c
y.tab.o: y.tab.c
	gcc -c y.tab.c
syntaxtree.o: syntaxtree.c
	gcc -c syntaxtree.c
y.tab.c y.tab.h:sql_syntax.y lex.yy.c
	bison -d -y -v -t sql_syntax.y
lex.yy.c:sql_scan.l
	flex sql_scan.l
clean:
	rm lex.yy.c y.tab.h y.tab.c redisql_parse.o y.tab.o syntaxtree.o sql_parse y.output *.a
