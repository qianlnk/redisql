redisql
----
##support Query
* Exact Match Query
* Range Query
* Simple query conditions
* Count
* Limit
* Other Comming soon

##how to use
###connect to redis
```golang
redisql.Connect("127.0.0.1", "6379", "", "tcp", 5, 120)
```
###select redis db default 0
```golang
redisql.Select(0)
```
###change database if you had created it
```golang
redisql.ChangeDatabase("databasename")
```
###or create it
```golang
redisql.CreateDatabase("databasename")
```
###create table
```golang
err := redisql.TABLE("user").FIELDS("name, age, city").TYPES("string", "int, string").CREATE()
```
###create index
```golang
err := redisql.TABLE("user").FIELDS("name").INDEX()
err := redisql.TABLE("user").FIELDS("age, city").INDEX()
```
###insert new data
```golang
err := INTO("user").FIELDS("name, age", "city").VALUES("qianlnk", 25, "sh").INSERT()
```
###drop
comming soon
#delete
comming soon
#select
```golang
type test struct {
	Id    []string `json:"myid"`
	Name  []string `json:"myname"`
	Age   []string `json:"myage"`
	Class []string `json:"class, myclass"`
}
var testres test
res, err := FROM("student a").FIELDS("a.id myid, a.name myname, a.age myage, a.class").WHERE("a.name like '%zhenjia' or (a.age <= 25 and a.class = 'jisuanji1002')").LIMIT(0, 2).RESELECT()
if err != nil {
	fmt.Println(res, err)
}
json.Unmarshal(res, &testres)
fmt.Println(testres)
```
