how to use
1、connect to redis
redisql.Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
2、select redis db default 0
redisql.Select(0)
3、change database if you had created it
redisql.ChangeDatabase("databasename")
or create it
redisql.CreateDatabase("databasename")
4、create table
err := redisql.TABLE("user").FIELDS("name, age, city").TYPES("string", "int, string").CREATE()
5、create index
err := redisql.TABLE("user").FIELDS("name").INDEX()
err := redisql.TABLE("user").FIELDS("age, city").INDEX()
6、insert new data
err := INTO("user").FIELDS("name, age", "city").VALUES("qianlnk", 25, "sh").INSERT()
