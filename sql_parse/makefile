task: sql_parse.a syntaxtree.a
	mv sql_parse.a ../lib/sql_parse.a
	mv syntaxtree.a ../lib/syntaxtree.a
sql_parse.a: lex.yy.o y.tab.o
	ar -crv sql_parse.a lex.yy.o y.tab.o
lex.yy.o: lex.yy.c y.tab.h
	cc -c lex.yy.c
y.tab.o: y.tab.c
	cc -c y.tab.c
syntaxtree.a: syntaxtree.o
	ar -crv syntaxtree.a syntaxtree.o
syntaxtree.o: syntaxtree.c
	cc -c syntaxtree.c
lex.yy.c:sql_scan.l
	flex sql_scan.l
y.tab.c y.tab.h:sql_syntax.y
	bison -d -y -v -t sql_syntax.y
clean:
	rm lex.yy.c lex.yy.o y.tab.h y.tab.c y.tab.o syntaxtree.o