#### mysql 体系结构图

![image](../mysql.png)

#### mysql 四层结构

#### 连接层
```text
最上层是一些客户端和链接服务,主要完成一些类似于连接处理、 授权认证、及相关的安全方案。
服务器也会为安全接入的每个客户
```

#### 服务层
```text
第二层架构主要完成大多数的核心服务功能，如SQL接口，并完成缓存的查询，SQL 的分析和优化，部分内置函数的执行。
所有跨存储引擎的功能也在这一层实现， 如过程、函数等。
```

#### 存储层
```text
存储引擎真正的负责了MySQL中数据的存储和提取,服务器通过API和存储引擎进行通信。
不同的存储引擎具有不同的功能，这样我们可以根据自己的需要,来选取合适的存储引擎。
```

#### 引擎层
```text
主要是将数据存储在文件系统之上，并完成与存储引擎的交互.
```

#### 存储引擎
```text
存储引擎就是存储数据、建立索引、更新/查询数据等技术的实现方式。存储引擎是基于表而不是基于库的，所以存储引擎也可以被称为表引擎。
Mysql在V5.1之前默认存储引擎是MyISAM；在此之后默认存储引擎是InnoDB
```

- 查看表的存储引擎

```sql
-- 查询当前表创建时使用的存储引擎
mysql> show create table user\G;
*************************** 1. row ***************************
       Table: user
Create Table: CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `sex` varchar(255) DEFAULT NULL,
  `age` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3
1 row in set (0.00 sec)

ERROR: 
No query specified
```

- 查看支持那些存储引擎
```sql
-- mysql支持的所有存储引擎
mysql> show engines;
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
| Engine             | Support | Comment                                                        | Transactions | XA   | Savepoints |
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
| FEDERATED          | NO      | Federated MySQL storage engine                                 | NULL         | NULL | NULL       |
| MEMORY             | YES     | Hash based, stored in memory, useful for temporary tables      | NO           | NO   | NO         |
| InnoDB             | DEFAULT | Supports transactions, row-level locking, and foreign keys     | YES          | YES  | YES        |
| PERFORMANCE_SCHEMA | YES     | Performance Schema                                             | NO           | NO   | NO         |
| MyISAM             | YES     | MyISAM storage engine                                          | NO           | NO   | NO         |
| MRG_MYISAM         | YES     | Collection of identical MyISAM tables                          | NO           | NO   | NO         |
| BLACKHOLE          | YES     | /dev/null storage engine (anything you write to it disappears) | NO           | NO   | NO         |
| CSV                | YES     | CSV storage engine                                             | NO           | NO   | NO         |
| ARCHIVE            | YES     | Archive storage engine                                         | NO           | NO   | NO         |
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
9 rows in set (0.00 sec)

-- mysql 默认使用的存储引擎

mysql> show variables like '%storage_engine%';
+---------------------------------+-----------+
| Variable_name                   | Value     |
+---------------------------------+-----------+
| default_storage_engine          | InnoDB    |
| default_tmp_storage_engine      | InnoDB    |
| disabled_storage_engines        |           |
| internal_tmp_mem_storage_engine | TempTable |
+---------------------------------+-----------+
4 rows in set (0.00 sec)

```

- 指定存储引擎创建表
```sql
-- $table_name: 表名
-- $engine_name: 需要指定的存储引擎名称
create table $table_name (
    ...
)engine = $engine_name

```

#### InnoDB 存储引擎
```text
InnoDB 是一种兼顾高可靠性和高性能的通用存储引擎，在 MySQL 5.1 之后，InnoDB 是默认的 MySQL 引擎。

特点：
  DML 操作遵循 ACID 模型，支持事务
  行级锁，提高并发访问性能
  支持外键约束，保证数据的完整性和正确性

文件：
  $filename.ibd: $filename表示表明, innodb引擎的每张表都会对应这样一个表空间文件，存储该表的表结构（frm、sdi）、数据和索引。

参数：
  innodb_file_per_table，决定多张表共享一个表空间还是每张表对应一个表空间
```

- 查看环境变量
```sql
-- 查看 Mysql 变量
mysql> show variables like 'innodb_file_per_table';
+-----------------------+-------+
| Variable_name         | Value |
+-----------------------+-------+
| innodb_file_per_table | ON    |
+-----------------------+-------+
1 row in set (0.00 sec)
msyql> 

-- 从innodb存储的文件中提取出文件内容
idb2sdi $filename

```
- innodb 逻辑存储结构
![image](../innodb.png)

#### myisam 存储引擎
```text
MyISAM 是 MySQL 早期的默认存储引擎。


特点：
  不支持事务，不支持外键
  支持表锁，不支持行锁
  访问速度快

文件：
  $filename.sdi: 存储表结构信息
  $filename.MYD: 存储数据
  $filename.MYI: 存储索引
```


#### Memory存储引擎

```text
Memory 引擎的表数据是存储在内存中的，受硬件问题、断电问题的影响，只能将这些表作为临时表或缓存使用。


特点：
  存放在内存中，速度快
  hash索引（默认）
文件：
  $filename.sdi: 存储表结构信息
```

#### 存储引擎特点
#### 存储引擎特点
| 特点        | InnoDB    | MyISAM | Memory |
| :--------: | :---------:|:---------:| -----:|
|Harry Potter | Gryffindor| 90 |
|存储限制 | 64TB |	有 |有 |
|事务安全 | 支持 |	-	| - |
|锁机制 | 行锁	| 表锁 | 表锁 |
|B+tree索引 | 支持 | 支持 | 支持 |
|Hash索引 | - | - | 支持 |
|全文索引 | 支持(5.6版本之后) | 支持 | - |
|空间使用 | 高 | 低 | N/A |
|内存使用 | 高 | 低 | 中等 |
|批量插入速度 | 低 | 高 | 高 |
|支持外键 | 支持 | - | - |


#### 存储引擎的选择

**在选择存储引擎时，应该根据应用系统的特点选择合适的存储引擎。对于复杂的应用系统，还可以根据实际情况选择多种存储引擎进行组合。**

- InnoDB: 如果应用对事物的完整性有比较高的要求，在并发条件下要求数据的一致性，数据操作除了插入和查询之外，还包含很多的更新、删除操作，则 InnoDB 是比较合适的选择
- MyISAM: 如果应用是以读操作和插入操作为主，只有很少的更新和删除操作，并且对事务的完整性、并发性要求不高，那这个存储引擎是非常合适的。
- Memory: 将所有数据保存在内存中，访问速度快，通常用于临时表及缓存。Memory 的缺陷是对表的大小有限制，太大的表无法缓存在内存中，而且无法保障数据的安全性
- 电商中的足迹和评论适合使用 MyISAM 引擎，缓存适合使用 Memory 引擎