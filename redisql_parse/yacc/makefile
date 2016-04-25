cflags = -g -fpic -c
ldflags = -shared -fpic -O2
libparse.so: parse.o y.tab.o parse_redisql.o
	gcc $(ldflags) y.tab.o parse_redisql.o -o libparse.so
lex.yy.o: lex.yy.c y.tab.h
	gcc -c lex.yy.c
y.tab.o: y.tab.c
	gcc -c y.tab.c
parse_redisql.o: parse_redisql.c
	gcc -c parse_redisql.c
y.tab.c y.tab.h:yacc_redisql.y lex.yy.c
	bison -d -y -v -t yacc_redisql.y
lex.yy.c:lex_redisql.l
	flex lex_redisql.l
clean:
	rm lex.yy.c y.tab.h y.tab.c parse.o y.tab.o parse_redisql.o
