### mysql 中的索引

#### 索引概述
```text
索引是帮助 MySQL 高效获取数据的数据结构（有序）。
在数据之外，数据库系统还维护着满足特定查找算法的数据结构，这些数据结构以某种方式引用（指向）数据，这样就可以在这些数据结构上实现高级查询算法，这种数据结构就是索引。

优点：
    提高数据检索效率，降低数据库的IO成本
    通过索引列对数据进行排序，降低数据排序的成本，降低CPU的消耗

缺点：
    索引列也是要占用空间的
    索引大大提高了查询效率，但降低了更新的速度，比如 INSERT、UPDATE、DELETE
```

#### 索引结构
- 索引的基本结构
    - **mysql的索引结构是在存储引擎层实现的，不同的存储引擎有不同的结构，主要包含以下结构.**
    - **平常如果没有特指索引的结构，那么默认指的是b+tree结构**

 索引结构  | 描述
- | :-: | :-: 
索引结构 | 描述
B+Tree  | 最常见的索引类型，大部分引擎都支持B+树索引
Hash | 底层数据结构是用哈希表实现，只有精确匹配索引列的查询才有效，不支持范围查询
R-Tree(空间索引) | 空间索引是 MyISAM 引擎的一个特殊索引类型，主要用于地理空间数据类型，通常使用较少
Full-Text(全文索引) | 是一种通过建立倒排索引，快速匹配文档的方式，类似于 Lucene, Solr, ES

-  不同的存储引擎对索引的支持

索引 | InnoDB | MyISAM | Memory
- | :-: | :-: | :-: | :-:
B+Tree索引 | 支持 | 支持 | 支持
Hash索引 | 不支持 | 不支持 | 支持
R-Tree索引 | 不支持 | 支持 | 不支持
Full-text | 5.6版本后支持 | 支持 | 不支持

##### 二叉树索引结构

![image](./image/%E4%BA%8C%E5%8F%89%E6%A0%91%E7%BB%93%E6%9E%84.png)

```text
二叉树的缺点：
    顺序插入时，会形成一个单链表，查询性能大大降低，数据量比较大的情况下，树的高度比较高，索引速度比较慢.
```
- 红黑树的索引结构

![iamge](./image/%E7%BA%A2%E9%BB%91%E6%A0%91%E7%BB%93%E6%9E%84.png)

```text
红黑树也存在大数据量情况下，层级较深，检索速度慢的问题。为了解决上述问题，可以使用 B-Tree 结构。
```

- B-Tree索引结构(多路平衡查找树)

![image](./image/B-Tree%E7%BB%93%E6%9E%84.png)

- 标准B+Tree 结构

![image](./image/%E6%A0%87%E5%87%86b%2Btree%E7%BB%93%E6%9E%84.PNG)

- B+Tree 索引结构(优化之后的)

![image](./image/b%2Btree%E7%B4%A2%E5%BC%95%E7%BB%93%E6%9E%84.png)

```text
与 B-Tree 的区别：
    所有的数据都会出现在叶子节点
    叶子节点形成一个单向链表

MySQL 索引数据结构对经典的 B+Tree 进行了优化。在原 B+Tree 的基础上，增加一个指向相邻叶子节点的链表指针，就形成了带有顺序指针的 B+Tree，提高区间访问的性能。

```

##### hash 索引结构

```text
哈希索引就是采用一定的hash算法，将键值换算成新的hash值，映射到对应的槽位上，然后存储在hash表中。
如果两个（或多个）键值，映射到一个相同的槽位上，他们就产生了hash冲突（也称为hash碰撞），可以通过链表来解决。
```

- hash 索引原理图

![image](./image/Hash%E7%B4%A2%E5%BC%95%E5%8E%9F%E7%90%86%E5%9B%BE.png)

- hash 索引的特点
```text
特点：
    Hash索引只能用于对等比较（=、in），不支持范围查询（betwwn、>、<、…）
    无法利用索引完成排序操作
    查询效率高，通常只需要一次检索就可以了，效率通常要高于 B+Tree 索引

存储引擎支持：
    Memory
    InnoDB: 具有自适应hash功能，hash索引是存储引擎根据 B+Tree 索引在指定条件下自动构建的

```

#### 索引分类

分类 | 含义 | 特点 | 关键字
- | :-: | :-: | :-: | 
主键索引 | 针对于表中主键创建的索引 | 默认自动创建，只能有一个 | PRIMARY
唯一索引 | 避免同一个表中某数据列中的值重复 | 可以有多个 | UNIQUE
常规索引 | 快速定位特定数据 | 可以有多个 | -
全文索引 | 全文索引查找的是文本中的关键词，而不是比较索引中的值 | 可以有多个 | FULLTEXT

- 在 InnoDB 存储引擎中，根据索引的存储形式，又可以分为以下两种：

分类 | 含义 | 特点
- | :-: | :-: | :-: | 
聚集索引(Clustered Index) | 将数据存储与索引放一块，索引结构的叶子节点保存了行数据 | 必须有，而且只有一个
二级索引(Secondary Index) | 将数据与索引分开存储，索引结构的叶子节点关联的是对应的主键 | -

- 狙击索引,二级索引

![image](./image/%E7%B4%A2%E5%BC%95%E5%88%86%E7%B1%BB.png)

![image](./image/%E7%B4%A2%E5%BC%95%E5%88%86%E7%B1%BB%E5%8E%9F%E7%90%86.png)

```sql
-- 1) 根据name的索引查出聚集索引id
-- 2) 之后需要回表根据id查出所需要的内容
select * from name="arm"
```

- 聚集索引选取规则：
```text
1) 如果存在主键，主键索引就是聚集索引
2) 如果不存在主键，将使用第一个唯一(UNIQUE)索引作为聚集索引
3) 如果表没有主键或没有合适的唯一索引，则 InnoDB 会自动生成一个 rowid 作为隐藏的聚集索引
```


#### 索引语法

```sql
-- 1) 创建索引
CREATE [ UNIQUE | FULLTEXT ] INDEX $index_name ON $table_name ($index_col_name, ...);

-- 2) 查看索引
SHOW INDEX FROM $table_name;

-- 3) 删除索引
DROP INDEX $index_name ON $table_name;

-- 4) case

-- name字段为姓名字段，该字段的值可能会重复，为该字段创建索引
create index idx_user_name on tb_user(name);

-- phone手机号字段的值非空，且唯一，为该字段创建唯一索引
create unique index idx_user_phone on tb_user (phone);

-- 为profession, age, status创建联合索引
create index idx_user_pro_age_stat on tb_user(profession, age, status);

-- 为email建立合适的索引来提升查询效率
create index idx_user_email on tb_user(email);

-- 删除索引
drop index idx_user_email on tb_user;

```

#### sql性能分析
- 查看全局的操作频率
```sql
-- 查看当前数据库的 INSERT, UPDATE, DELETE, SELECT 访问频次：
-- SHOW GLOBAL STATUS LIKE 'Com_______'; 或者 SHOW SESSION STATUS LIKE 'Com_______';
mysql> show global status like 'Com_______';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| Com_binlog    | 0     |
| Com_commit    | 3     |
| Com_delete    | 0     |
| Com_import    | 0     |
| Com_insert    | 0     |
| Com_repair    | 0     |
| Com_revoke    | 0     |
| Com_select    | 1452  |
| Com_signal    | 0     |
| Com_update    | 0     |
| Com_xa_end    | 0     |
+---------------+-------+
11 rows in set (0.01 sec)

```

- 慢查询日志
```sql
-- 慢查询日志记录了所有执行时间超过指定参数（long_query_time，单位：秒，默认10秒）的所有SQL语句的日志。
-- MySQL的慢查询日志默认没有开启，需要在MySQL的配置文件（/etc/my.cnf）中配置如下信息：

-- 开启慢查询日志开关
slow_query_log=1
--  设置慢查询日志的时间为2秒，SQL语句执行时间超过2秒，就会视为慢查询，记录慢查询日志
long_query_time=2
```

![image](./image/%E6%85%A2%E6%9F%A5%E8%AF%A2.png)


```sql
-- 使用修改配置文件的方式不生效
mysql> set global slow_query_log='ON';
Query OK, 0 rows affected (0.00 sec)

-- 查看慢查询日志开关状态：
mysql> show variables like 'slow_query_log';
+----------------+-------+
| Variable_name  | Value |
+----------------+-------+
| slow_query_log | ON    |
+----------------+-------+
1 row in set (0.00 sec)

-- 设置慢查询log的位置
mysql> set global slow_query_log_file='/var/lib/mysql/localhost-slow.log';
Query OK, 0 rows affected (0.00 sec)

-- root@iZ2ze58f53sxjm9z7mgn5xZ ~]# ls /var/lib/mysql/localhost-slow.log
-- /var/lib/mysql/localhost-slow.log

```

- profile
```sql
-- show profile 能在做SQL优化时帮我们了解时间都耗费在哪里。通过 have_profiling 参数，能看到当前 MySQL 是否支持 profile 操作：

-- 查看是否支持 profile
mysql> SELECT @@have_profiling;
+------------------+
| @@have_profiling |
+------------------+
| YES              |
+------------------+
1 row in set, 1 warning (0.00 sec)


-- profiling 默认关闭，可以通过set语句在session/global级别开启 profiling：
-- 查询是否打开profile
mysql> select @@profiling;
+-------------+
| @@profiling |
+-------------+
|           0 |
+-------------+
1 row in set, 1 warning (0.00 sec)

-- 设置profiling
mysql> set profiling = 1;
Query OK, 0 rows affected, 1 warning (0.00 sec)

-- 设置完成查看已经打开
mysql> select @@profiling;
+-------------+
| @@profiling |
+-------------+
|           1 |
+-------------+
1 row in set, 1 warning (0.00 sec)


-- 使用 查询每一条sql的耗时时长
mysql> show profiles;
+----------+------------+----------------------------+
| Query_ID | Duration   | Query                      |
+----------+------------+----------------------------+
|        1 | 0.00016250 | select @@profiling         |
|        2 | 0.00429275 | select * from gorm_db.user |
+----------+------------+----------------------------+
2 rows in set, 1 warning (0.00 sec)


-- 查看指定query_id的SQL语句各个阶段的耗时：
mysql> show profile for query 2;
+--------------------------------+----------+
| Status                         | Duration |
+--------------------------------+----------+
| starting                       | 0.000072 |
| Executing hook on transaction  | 0.000007 |
| starting                       | 0.000008 |
| checking permissions           | 0.000007 |
| Opening tables                 | 0.003542 |
| init                           | 0.000013 |
| System lock                    | 0.000009 |
| optimizing                     | 0.000004 |
| statistics                     | 0.000011 |
| preparing                      | 0.000014 |
| executing                      | 0.000548 |
| end                            | 0.000009 |
| query end                      | 0.000004 |
| waiting for handler commit     | 0.000008 |
| closing tables                 | 0.000007 |
| freeing items                  | 0.000019 |
| cleaning up                    | 0.000012 |
+--------------------------------+----------+
17 rows in set, 1 warning (0.00 sec)

-- 查看指定query_id的SQL语句CPU的使用情况
mysql> show profile cpu for query 2;
+--------------------------------+----------+----------+------------+
| Status                         | Duration | CPU_user | CPU_system |
+--------------------------------+----------+----------+------------+
| starting                       | 0.000072 | 0.000000 |   0.000071 |
| Executing hook on transaction  | 0.000007 | 0.000000 |   0.000006 |
| starting                       | 0.000008 | 0.000000 |   0.000008 |
| checking permissions           | 0.000007 | 0.000000 |   0.000006 |
| Opening tables                 | 0.003542 | 0.001278 |   0.002237 |
| init                           | 0.000013 | 0.000008 |   0.000005 |
| System lock                    | 0.000009 | 0.000005 |   0.000003 |
| optimizing                     | 0.000004 | 0.000003 |   0.000001 |
| statistics                     | 0.000011 | 0.000006 |   0.000004 |
| preparing                      | 0.000014 | 0.000009 |   0.000006 |
| executing                      | 0.000548 | 0.000103 |   0.000064 |
| end                            | 0.000009 | 0.000005 |   0.000004 |
| query end                      | 0.000004 | 0.000003 |   0.000002 |
| waiting for handler commit     | 0.000008 | 0.000005 |   0.000003 |
| closing tables                 | 0.000007 | 0.000004 |   0.000002 |
| freeing items                  | 0.000019 | 0.000012 |   0.000008 |
| cleaning up                    | 0.000012 | 0.000007 |   0.000004 |
+--------------------------------+----------+----------+------------+
17 rows in set, 1 warning (0.00 sec)

```

- explain 执行计划
**explain 或者 DESC 命令获取 MySQL 如何执行 SELECT 语句的信息，包括在 SELECT 语句执行过程中表如何连接和连接的顺序。**

```sql
--  直接在select语句之前加上关键字 explain / desc
explain SELECT 字段列表 FROM 表名 HWERE 条件;

mysql> explain select * from gorm_db.user;
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------+
| id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------+
|  1 | SIMPLE      | user  | NULL       | ALL  | NULL          | NULL | NULL    | NULL |    1 |   100.00 | NULL  |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------+
1 row in set, 1 warning (0.00 sec)

```

explain 各字段含义：
```text
id：select 查询的序列号，表示查询中执行 select 子句或者操作表的顺序（id相同，执行顺序从上到下；id不同，值越大越先执行）

select_type：表示 SELECT 的类型，常见取值有 SIMPLE（简单表，即不适用表连接或者子查询）、PRIMARY（主查询，即外层的查询）、UNION（UNION中的第二个或者后面的查询语句）、SUBQUERY（SELECT/WHERE之后包含了子查询）等

type：表示连接类型，性能由好到差的连接类型为 NULL、system、const、
eq_ref、ref、range、index、all

possible_key：可能应用在这张表上的索引，一个或多个

Key：实际使用的索引，如果为 NULL，则没有使用索引

Key_len：表示索引中使用的字节数，该值为索引字段最大可能长度，并非实际使用长度，在不损失精确性的前提下，长度越短越好

rows：MySQL认为必须要执行的行数，在InnoDB引擎的表中，是一个估计值，可能并不总是准确的

filtered：表示返回结果的行数占需读取行数的百分比，filtered的值越大越好
```

#### 索引使用原则

**最左前缀法则**
```text
如果索引关联了多列（联合索引），要遵守最左前缀法则，最左前缀法则指的是查询从索引的最左列开始，并且不跳过索引中的列。
如果跳跃某一列，索引将部分失效（后面的字段索引失效）。

联合索引中，出现范围查询（<, >），范围查询右侧的列索引失效。可以用>=或者<=来规避索引失效问题。
```

```text
1) 在索引列上进行运算操作，索引将失效。如：explain select * from tb_user where substring(phone, 10, 2) = '15';
2) 字符串类型字段使用时，不加引号，索引将失效。如：explain select * from tb_user where phone = 17799990015;，此处phone的值没有加引号
3) 模糊查询中，如果仅仅是尾部模糊匹配，索引不会是失效；如果是头部模糊匹配，索引失效。如：explain select * from tb_user where profession like '%工程';，前后都有 % 也会失效。
4) 用 or 分割开的条件，如果 or 其中一个条件的列没有索引，那么涉及的索引都不会被用到。
5) 如果 MySQL 评估使用索引比全表更慢，则不使用索引。

```

**SQL 提示**
```sql
-- sql 提示是优化数据库的一个重要手段，简单来说，就是在SQL语句中加入一些人为的提示来达到优化操作的目的。

-- 使用索引：
explain select * from tb_user use index(idx_user_pro) where profession="软件工程";

-- 不使用哪个索引：
explain select * from tb_user ignore index(idx_user_pro) where profession="软件工程";

-- 必须使用哪个索引：
explain select * from tb_user force index(idx_user_pro) where profession="软件工程";

-- use 是建议，实际使用哪个索引 MySQL 还会自己权衡运行速度去更改，force就是无论如何都强制使用该索引。
```

**覆盖索引&回表查询**
```text
1) 尽量使用覆盖索引（查询使用了索引，并且需要返回的列，在该索引中已经全部能找到），减少 select *。

explain 中 extra 字段含义：
using index condition：查找使用了索引，但是需要回表查询数据
using where; using index;：查找使用了索引，但是需要的数据都在索引列中能找到，所以不需要回表查询

2) 如果在聚集索引中直接能找到对应的行，则直接返回行数据，只需要一次查询，哪怕是select *；如果在辅助索引中找聚集索引，如select id, name from xxx where name='xxx';，也只需要通过辅助索引(name)查找到对应的id，返回name和name索引对应的id即可，只需要一次查询；如果是通过辅助索引查找其他字段，则需要回表查询，如select id, name, gender from xxx where name='xxx';

3) 所以尽量不要用select *，容易出现回表查询，降低效率，除非有联合索引包含了所有字段
```

**前缀索引**
```text
1) 当字段类型为字符串（varchar, text等）时，有时候需要索引很长的字符串，这会让索引变得很大，查询时，浪费大量的磁盘IO，影响查询效率，此时可以只降字符串的一部分前缀，建立索引，这样可以大大节约索引空间，从而提高索引效率。

语法：create index idx_xxxx on table_name(columnn(n));

2） 前缀长度：可以根据索引的选择性来决定，而选择性是指不重复的索引值（基数）和数据表的记录总数的比值，索引选择性越高则查询效率越高，唯一索引的选择性是1，这是最好的索引选择性，性能也是最好的。
求选择性公式：

select count(distinct email) / count(*) from tb_user;
select count(distinct substring(email, 1, 5)) / count(*) from tb_user;
show index 里面的sub_part可以看到接取的长度
```

**单列索引&联合索引**
```text
1) 单列索引：即一个索引只包含单个列
2) 联合索引：即一个索引包含了多个列
3) 在业务场景中，如果存在多个查询条件，考虑针对于查询字段建立索引时，建议建立联合索引，而非单列索引。
4) case:
    单列索引情况：
    explain select id, phone, name from tb_user where phone = '17799990010' and name = '韩信';
    这句只会用到phone索引字段

```

**注意事项**
```text
多条件联合查询时，MySQL优化器会评估哪个字段的索引效率更高，会选择该索引完成本次查询
```

#### 索引设计原则

- 索引列设计原则

```text
1) 对于数据量较大，且查询比较频繁的表建立索引

2) 针对于常作为查询条件（where）、排序（order by）、分组（group by）操作的字段建立索引

3) 尽量选择区分度高的列作为索引，尽量建立唯一索引，区分度越高，使用索引的效率越高

4) 如果是字符串类型的字段，字段长度较长，可以针对于字段的特点，建立前缀索引

5) 尽量使用联合索引，减少单列索引，查询时，联合索引很多时候可以覆盖索引，节省存储空间，避免回表，提高查询效率

6) 要控制索引的数量，索引并不是多多益善，索引越多，维护索引结构的代价就越大，会影响增删改的效率

7) 如果索引列不能存储NULL值，请在创建表时使用NOT NULL约束它。当优化器知道每列是否包含NULL值时，它可以更好地确定哪个索引最有效地用于查询
```