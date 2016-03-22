#how to use
#connect to redis
redisql.Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
#select redis db default 0
redisql.Select(0)
#change database if you had created it
redisql.ChangeDatabase("databasename")
#or create it
redisql.CreateDatabase("databasename")
#create table
err := redisql.TABLE("user").FIELDS("name, age, city").TYPES("string", "int, string").CREATE()
#create index
err := redisql.TABLE("user").FIELDS("name").INDEX()

err := redisql.TABLE("user").FIELDS("age, city").INDEX()
#insert new data
err := INTO("user").FIELDS("name, age", "city").VALUES("qianlnk", 25, "sh").INSERT()
