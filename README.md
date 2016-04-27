redisql
----
##support Query Cmds
* use databasename
* show databases
* show tables
* show index from tablename
* create database databasename
* create table tablename(field1 type1, ...);
* create index indexname on tablename(field);
* insert into tablename(field1 ...) values(value1 ...)
* select
* Limit
* Other Comming soon

##how to use
###make
```golang
//build parse sql lib 'libparse.so' 
> cd redisql/redisql_parse/yacc
> make

//copy 'libparse.so' to redisql_cli or your project
//build client, need 'libparse.so'.
> cd redisql/redisql_cli
> go build
```
###client
![](https://github.com/qianlnk/redisql/blob/master/redisql.jpg)
###call by golang
```golang
redisql.Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
redisql.Selectdb(0)
res, err := redisql.Query("show databases")
if err != nil {
	fmt.Println(err)
	return
}
...
```
