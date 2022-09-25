### MongoDB

> MongoDB 是一个基于分布式文件存储的数据库。由 C++ 语言编写。旨在为 WEB 应用提供可扩展的高性能数据存储解决方案。
>
> MongoDB 是一个介于关系数据库和非关系数据库之间的产品，是非关系数据库当中功能最丰富，最像关系数据库的。

**特点**
- MongoDB 是一个面向文档存储的数据库，操作起来比较简单和容易。
- 如果负载的增加（需要更多的存储空间和更强的处理能力） ，它可以分布在计算机网络中的其他节点上这就是所谓的分片。
- MongoDB 支持丰富的查询表达式。查询指令使用JSON形式的标记，可轻易查询文档中内嵌的对象及数组。
- MongoDB 使用update()命令可以实现替换完成的文档（数据）或者一些指定的数据字段。
- MongoDB 中的Map/reduce主要是用来对数据进行批量处理和聚合操作。
- Map和Reduce，Map函数调用emit(key,value)遍历集合中所有的记录，将key与value传给Reduce函数进行处理。
- Map函数和Reduce函数是使用Javascript编写的，并可以通过db.runCommand或mapreduce命令来执行MapReduce操作。
- GridFS是MongoDB中的一个内置功能，可以用于存放大量小文件。
- MongoDB 允许在服务端执行脚本，可以用Javascript编写某个函数，直接在服务端执行，也可以把函数的定义存储在服务端，下次直接调用即可。
- MongoDB 支持各种编程语言:RUBY，PYTHON，JAVA，C++，PHP，C#等多种语言

**mongodb概念解析**
| SQL术语/概念 | MongoDB术语/概念 | 解释/说明 |
| :---: | :---: | :---: | 
| database | database | 数据库 |
| table | collection | 数据库表/集合 |
| row | document | 数据记录行/文档 |
| column | field | 数据字段/域 |
| index | index | 索引 |
| table | joins | 表连接,MongoDB不支持 |
| primary key | primary key | 主键,MongoDB自动将_id字段设置为主键 |


### 基于docker启动mongo
```sh
# --auth 需要密码才能访问容器服务
>$ docker run -itd --name mongo -p 27017:27017 mongo --auth

>$ docker exec -it mongo mongo admin
# 创建一个名为 admin，密码为 admin的用户
>$  db.createUser({ user:'admin',pwd:'admin',roles:[ { role:'userAdminAnyDatabase', db: 'admin'}]});
# 尝试使用上面创建的用户信息进行连接
>$ db.auth('admin', 'admin')

# 查看用户
> show users
{
        "_id" : "admin.admin",
        "userId" : UUID("0a705d89-af5e-4075-9d89-536c696ab9ef"),
        "user" : "admin",
        "db" : "admin",
        "roles" : [
                {
                        "role" : "userAdminAnyDatabase",
                        "db" : "admin"
                }
        ],
        "mechanisms" : [
                "SCRAM-SHA-1",
                "SCRAM-SHA-256"
        ]
}
```

### mongoDB 常用操作
**database**
```sh
# 创建database
> use mydb
switched to db mydb
> db
mydb

# 查看所有database
> show dbs
admin   0.000GB
config  0.000GB
local   0.000GB

# 删除database
> db.dropDatabase()
{ "ok" : 1 }
```

**collection**
```shell
# 创建集合
# name: 要创建的集合名称
# options: 可选参数, 指定有关内存大小及索引的选项
db.createCollection(name, options)

> use mydb
switched to db mydb
# 直接创建
> db.createCollection("myCollection")
{ "ok" : 1 }

# 指定参数创建
> db.createCollection("myCollection1", { capped : true, autoIndexId : true, size : 
...    6142800, max : 10000 } )
{
        "note" : "The autoIndexId option is deprecated and will be removed in a future release",
        "ok" : 1
}

# 查看collection
> show collections
myCollection

# 在mongodb中，不需要创建collection，插入数据的时候会自动创建
> db.web.insert({"name":"web"})
WriteResult({ "nInserted" : 1 })

> show collections
myCollection
myCollection1
web

# 删除collection
# db.collection.drop()
> db.myCollection1.drop()
true
> show collections
myCollection
web

```

**插入文档(insert/save)**
```shell
# db.COLLECTION_NAME.insert(document)

> db.web.insert({"name":"web"})
WriteResult({ "nInserted" : 1 })

# 用于向集合插入一个新文档
> db.collection.insertOne(
    <document>,
    { 
        writeConcern: <document>
    }
) 

# 用于向集合插入一个多个文档
db.collection.insertMany(
    [ <document 1> , <document 2>, ... ],
    {
        writeConcern: <document>,
        ordered: <boolean>
        }
) 

# document：要写入的文档
# writeConcern：写入策略，默认为 1，即要求确认写操作，0 是不要求
# ordered：指定是否按顺序写入，默认 true，按顺序写入
```

**更新文档(update)**
```shell
# update() 方法用于更新已存在的文档
db.collection.update(
    <query>, 
    <update>,
    {
        upsert: <boolean>,
        multi: <boolean>, 
        writeConcern: <document>
    }
)

# query : update的查询条件，类似sql update查询内where后面的。
# update : update的对象和一些更新的操作符（如$,$inc...）等，也可以理解为sql update查询内set后面的
# upsert : 可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
# multi : 可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。
# writeConcern :可选，抛出异常的级别。

db.mycol.insertMany([{
    "name":"c++",
    "age":20
},{
    "name":"python",
    "age":25
},{
    "name":"c",
    "age":90
},{
    "name":"java",
    "age":35
}]
)


db.mycol.update({"name":"python"},{$set:{"name":"python3"}})

db.mycol.find()
[
  {"_id": {"$oid": "63305db8e6aef633eca11e8f"},"age": 14,"name": "golang"},
  {"_id": {"$oid": "63305e3be6aef633eca11e90"},"age": 20,"name": "c++"},
  {"_id": {"$oid": "63305e3be6aef633eca11e91"},"age": 25,"name": "python3"},
  {"_id": {"$oid": "63305e3be6aef633eca11e92"},"age": 90,"name": "c"},
  {"_id": {"$oid": "63305e3be6aef633eca11e93"},"age": 35,"name": "java"}
]
```

**删除文档**
```shell
db.collection.remove(
   <query>,
   {
     justOne: <boolean>,
     writeConcern: <document>
   }
)
# query :（可选）删除的文档的条件。
# justOne : （可选）如果设为 true 或 1，则只删除一个文档，如果不设置该参数，或使用默认值 false，则删除所有匹配条件的文档。
# writeConcern :（可选）抛出异常的级别。

db.mycol.remove({"name":"java"})

db.mycol.find()
[
  {"_id": {"$oid": "63305db8e6aef633eca11e8f"},"age": 14,"name": "golang"},
  {"_id": {"$oid": "63305e3be6aef633eca11e90"},"age": 20,"name": "c++"},
  {"_id": {"$oid": "63305e3be6aef633eca11e91"},"age": 25,"name": "python3"},
  {"_id": {"$oid": "63305e3be6aef633eca11e92"},"age": 90,"name": "c"}
]
```

#### 查询文档
```text
db.collection.find(query, projection)

query ：可选，使用查询操作符指定查询条件
projection ：可选，使用投影操作符指定返回的键。查询时返回文档中所有键值， 只需省略该参数即可（默认省略）
```
MongoDB 与 RDBMS Where 语句比较

| 操作 | 格式 | 范例 | RDBMS中的类似语句 |
| :---: | :---: | :---: | :---: |
| 等于 | {key:value} | db.col.find({"name":"tom1"}).pretty() | where name='tom1' |
| 小于 | {key:{$lt:value}} | db.col.find({"likes":{$lt:50}}).pretty() | where likes < 50 |
| 小于或等于 | {key>:{$lte:value}} | db.col.find({"likes":{$lte:50}}).pretty() | where likes <= 50 |
| 大于 | {key:{$gt:value}} | db.col.find({"likes":{$gt:50}}).pretty() | where likes > 50 |
| 大于或等于 | {key:{$gte:value}} | db.col.find({"likes":{$gte:50}}).pretty() | where likes >= 50 |
| 不等于 | {key:{$ne:value}} | db.col.find({"likes":{$ne:50}}).pretty() | where likes != 50 |

**and/or**
```shell
# or
db.mycol.find({
    $or: [
        {"name":"python3"},
        {"age": 90}
    ]
})

# [
#   {"_id": {"$oid": "63305e3be6aef633eca11e91"},"age": 25,"name": "python3"},
#   {"_id": {"$oid": "63305e3be6aef633eca11e92"},"age": 90,"name": "c"}
# ]

# and + or
# age > 30 and name == 'python3' or age=90
db.mycol.find(
    {
        "age": {$gt: 30},
        $or:[
            {"name":"python3"},
            {"age": 90}
        ]}
)
# [{"_id": {"$oid": "63305e3be6aef633eca11e92"},"age": 90,"name": "c"}]
```