### Elasticsearch

>  **简介:**  Elasticsearch是一个高度可扩展的、开源的、基于 Lucene 的全文搜索和分析引擎。它允许您快速，近实时地存储，搜索和分析大量数据，并支持多租户。 Elasticsearch也使用Java开发并使用 Lucene 作为其核心来实现所有索引和搜索的功能，但是它的目的是通过简单的 RESTful API 来隐藏 Lucene 的复杂性，从而让全文搜索变得简单。

> ElasticSearch是面向文档型的数据库，一条数据在这里就是一个文档。

**es和关系型数据库概念对比**


ES | 关系型数据库
| :---: | :---: |
| 索引（Index）| 数据库（DataBase）|
| 类型（Type）| 表（Table）|
| 映射（mapping）| 表结构（Schema）}
| 文档（Document）| 行（Row）|
| 字段（Field）| 列（Column）|
| 反向索引 | 正向索引 |
| DSL查询 | SQL查询 |

**概念介绍**
```text
Segment：段，Lucence中存储是按段来进行存储，每个段相当于一个数据集。

Commit Point：提交点，记录着Lucence中所有段的集合。

Lucene Index：Lucene索引，由一堆Segment段集合和commit point组成。
```
#### 1. ES 原理

##### 1.1 存储过程
```text
1) 存储文档经过词法分析得到一系列的词(Term)

2) 通过一系列词来创建形成词典和反向索引表

3) 将索引进行存储并以文件的方式落盘
```

##### 1.2 查询过程
```text
a) 用户输入查询语句。

b) 对查询语句经过词法分析得到一系列词(Term) 。

c) 通过语法分析得到一个查询树。

d) 通过索引存储将索引读入到内存。

e) 利用查询树搜索索引，从而得到每个词(Term) 的文档链表，对文档链表进行交、差、并得到结果文档。

f) 将搜索到的结果文档对查询的相关性进行排序。

g) 返回查询结果给用户
```

##### 1.3 ES 数据写入
**写数据的基本过程**
```text
(1) 客户端选择一个ES节点发送写请求，ES节点接收请求变为协调节点。

(2) 协调节点判断写请求中如果没有指定文档id，则自动生成一个doc_id。协调节点对doc_id进行哈希取值，判断出文档应存储在哪个切片中。协调节点找到存储切片的对应节点位置，将请求转发给对应的node节点。

(3) Node节点的primary shard处理请求，并将数据同步到replica shard

(4) 协调节点发现所有的primary shard和所有的replica shard都处理完之后，就返回结果给客户端
```
**写数据的基本原理**
```text
(1) 数据先写入内存 buffer，然后每隔 1s，将数据 refresh 到操作系统缓存（os cache），生成新的segment。（os cache 中存储的数据能被搜索到）

(2) 写入 os cache 中的translog数据，默认每隔 5 秒刷一次到磁盘中去，如果translog 大到一定程度，或者默认每隔 30mins，会触发 commit 操作，将缓冲区的数据都 flush 到 segment file 磁盘文件中。
```

##### 1.4 ES 数据读取
```text
(1) 客户端给任意一个节点发送请求，该节点变为协调节点

(2) 协调节点根据doc_id，进行哈希取值，判断出文档存储在哪个切片上。

(3) 协调节点将请求转发到对应的节点上，然后使用随机轮询算法（round-robin）,在切片和副本切片中随机选择一个，以使读请求负载均衡

(4) 接收请求的节点返回文档数据给协调节点，协调节点再返回数据给客户端。
```

##### 1.5 ES 数据检索
```text
ES读数据是通过doc_id来进行查询，先根据doc_id判断出文档存储在哪个切片上，再从切片上把数据读取过来。

(1) 客户端给任意一个节点发送请求，该节点变为协调节点

(2)  协调节点根据doc_id，进行哈希取值，判断出文档存储在哪个切片上。

(3) 协调节点将请求转发到对应的节点上，然后使用随机轮询算法（round-robin）,在切片和副本切片中随机选择一个，以使读请求负载均衡

(4) 接收请求的节点返回文档数据给协调节点，协调节点再返回数据给客户端。
```

##### 1.6 ES 关键字检索
```text
ES检索关键词是ES最常使用的做法，通过关键词，将包含关键词的文档全部搜索出来。

(1) 客户端向任意一个节点发送请求，该节点变为协调节点

(2) 协调节点将搜索请求转到所有的shard上

(3) 每个shard将自身的检索结果（搜索到的doc_id和分数）,返回给协调节点。

(4) 协调节点根据检索结果进行相关性排序，产出最终的结果。再把doc_id发送给各个节点，拉取文档数据，最终返回给客户端。
```

##### 1.7 ES 数据删除
```text
删除操作，是在commit 的时候会生成一个.del文件，里面将doc标识为deleted状态，搜索的时候根据.del文件就知道这个 doc 是否被删除了
```

#### 2. docker-compose启动es
```yaml
version: '3'
services:
  elasticsearch_n0:
    image: elasticsearch:6.6.2
    container_name: elasticsearch_n0
    privileged: true
    environment:
      - cluster.name=elasticsearch-cluster
      - node.name=node0
      - node.master=true
      - node.data=true
      - bootstrap.memory_lock=true
      - http.cors.enabled=true
      - http.cors.allow-origin=*
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
      - "discovery.zen.ping.unicast.hosts=elasticsearch_n0,elasticsearch_n1"
      - "discovery.zen.minimum_master_nodes=2"
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - ./data/node0:/usr/share/elasticsearch/data
      - ./logs/node0:/usr/share/elasticsearch/logs
    ports:
      - 9200:9200
  elasticsearch_n1:
    image: elasticsearch:6.6.2
    container_name: elasticsearch_n1
    privileged: true
    environment:
      - cluster.name=elasticsearch-cluster
      - node.name=node1
      - node.master=true
      - node.data=true
      - bootstrap.memory_lock=true
      - http.cors.enabled=true
      - http.cors.allow-origin=*
      - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
      - "discovery.zen.ping.unicast.hosts=elasticsearch_n0,elasticsearch_n1"
      - "discovery.zen.minimum_master_nodes=2"
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - ./data/node1:/usr/share/elasticsearch/data
      - ./logs/node1:/usr/share/elasticsearch/logs
    ports:
      - 9201:9200
  kibana:
    image: docker.elastic.co/kibana/kibana:6.6.2
    container_name: kibana
    environment:
      - SERVER_NAME=kibana
      - ELASTICSEARCH_URL=http://elasticsearch_n0:9200
      - XPACK_MONITORING_ENABLED=true
    ports:
      - 5601:5601
    depends_on:
      - elasticsearch_n0

```

##### 2.1 启动,查看docker-compose
```sh
# 启动docker-compose
>$ docker-compose -f $docker_compose_file up -d

# 查看启动情况
>$ docker ps

CONTAINER ID   IMAGE                                   COMMAND                  CREATED        STATUS        PORTS                              NAMES
3a35d3d6c466   docker.elastic.co/kibana/kibana:6.6.2   "/usr/local/bin/kiba…"   15 hours ago   Up 15 hours   0.0.0.0:5601->5601/tcp             kibana
818ee8121acd   elasticsearch:6.6.2                     "/usr/local/bin/dock…"   15 hours ago   Up 15 hours   9300/tcp, 0.0.0.0:9201->9200/tcp   elasticsearch_n1
94765e5405dd   elasticsearch:6.6.2                     "/usr/local/bin/dock…"   15 hours ago   Up 15 hours   0.0.0.0:9200->9200/tcp, 9300/tcp   elasticsearch_n0


# 如果启动失败可以查看启动日志
>$ docker logs -f $container_name
```

#### 3. 倒排索引
> **正排索引:** 是以文档对象的唯一 ID 作为索引，以文档内容作为记录的结构。
> **倒排索引:** Inverted index，指的是将文档内容中的单词作为索引，将包含该词的文档 ID 作为记录的结构

> 根据倒排索引的概念，可以用一个 Map 来简单描述这个结构。这个 Map 的 Key 的即是分词后的单词，这里的单词称为 Term，这一系列的 Term 组成了倒排索引的第一个部分 —— **Term Dictionary(索引表**，可简称为 Dictionary)。
> 倒排索引的另一部分为 **Postings List（记录表）** ，也对应上述 Map 结构的 Value 部分集合。记录表由所有的 Term 对应的数据（Postings） 组成，它不仅仅为文档 id 信息，可能包含以下信息

```text
(1) 词条（term）：索引里面最小的存储和查询单元，对于英文来说是一个词，对于中文来说一般指分词后的一个词。

(2) 词典（Term Dictionary）：也叫字典，是词条的组合。搜索引擎的通常索引单位是单词，单词词典是文档集合中出现过的所有单词构成的字符串集合，单词词典内每条索引项记载单词本身的一些信息以及指向倒排所有的指针。
(3) 倒排表（Post list）：一个文档通常由多个词组成，倒排表记录的是某个词在哪些文档里出现过及出现的位置。每个记录称为一个倒排项（Posting），倒排表记录的不单单是文档编号，还记录了词频等信息。

(4) 倒排文件（Inverted File）：所有单词的倒排列表往往顺序地存储在磁盘的某个文件里，这个文件被称之为倒排文件，倒排文件是存储倒排索引的物理文件。
```
##### 3.1 Term Dictionary 实现
> Terms Dictionary 通过 .tim 后缀文件存储，其内部采用 NodeBlock 对 Term 进行压缩前缀存储，处理过程会将相同前缀的的 Term 压缩为一个 NodeBlock，NodeBlock 会存储公共前缀，然后将每个 Term 的后缀以及对应 Term 的 Posting 关联信息处理为一个 Entry 保存到 Block


**假设为公共前缀为 a 的 Term 集合，内部部分 Term 的又包含了相同前缀 ab，这时这部分 Term 就会处理为一个嵌套的 Block**
```text
termA: {a,$termB(prefix=ab),ac,ar}
termB: {abc,abd.aba...}

-- termA 为公共前缀a的term集合，其中一个term又包含了前缀为ab的term的集合，这个term集合处理为一个嵌套block
```

###### 3.1.1 数据查找
> Terms Dictionary 是按 NodeBlock 存储在.tim 文件上。当文档数量越来越多的时，Dictionary 中的 Term 也会越来越多，那查询效率必然也会逐渐变低
> Lucene 采用了 **FST(Finite State Transducer（有限状态转换器）)** 这个数据结构来实现这个索引
- 给定一个 Input 可以得到一个 output，相当于 HashMap
- 共享前缀、后缀节省空间，FST 的内存消耗要比 HashMap 少很多
- 词查找复杂度为 O(len(str))　　
- 构建后不可变更

> FST 是通过 Dictionary 的每个 NodeBlock 的前缀构成，所以通过 FST 只可以直接找到这个 NodeBlock 在 .tim 文件上具体的 File Pointer, 然后还需要在 NodeBlock 中遍历 Entry 匹配后缀进行查找

```text
(1) 快速试错，即是在 FST 上找不到可以直接跳出不需要遍历整个 Dictionary，类似于 BloomFilter。

(2) 快速定位 Block 的位置，通过 FST 是可以直接计算出 Block 的在文件中位置。

(3) FST 也是一个 Automaton(自动状态机)。这是正则表达式的一种实现方式，所以 FST 能提供正则表达式的能力。通过 FST 能够极大的提高近似查询的性能，包括通配符查询、SpanQuery、PrefixQuery 等
```


##### 3.2 Posting List 实现
> PostingList 包含文档 id、词频、位置等多个信息，这些数据之间本身是相对独立的，因此 Lucene 将 Postings List 被拆成三个文件存储：

- .doc 后缀文件：记录 Postings 的 docId 信息和 Term 的词频

- .pay 后缀文件：记录 Payload 信息和偏移量信息

- .pos 后缀文件：记录位置信息

> 基本所有的查询都会用 .doc 文件获取文档 id，且一般的查询仅需要用到 .doc 文件就足够了，只有对于近似查询等位置相关的查询则需要用位置相关数据。

**.doc 文件存储的是每个 Term 对应的文档 Id 和词频。每个 Term 都包含一对 TermFreqs 和 SkipData 结构**
> 其中 TermFreqs 存放 docId 和词频信息，SkipData 为跳表信息，用于实现 TermFreqs 内部的快速跳转。

###### 3.2.1 TermFreqs

> TermFreqs 存储文档号和对应的词频，它们两是一一对应的两个 int 值。Lucene 为了尽可能的压缩数据，采用的是混合存储 ，由 PackedBlock 和 VIntBlocks 两种结构组成。

**PackedBlock**
> PackedBlock 是采用 PackedInts 结构将一个 int[] 压缩打包成一个紧凑的 Block。它的压缩方式是取数组中最大值所占用的 bit 长度作为一个预算的长度，然后将数组每个元素按这个长度进行截取，以达到压缩的目的

**VIntBlock**
> VIntBlock 是采用 VInt 来压缩 int 值，对于绝大多数语言，int 型都占 4 个字节，不论这个数据是 1、100、1000、还是 1000,000。VInt 采用可变长的字节来表示一个整数。数值较大的数，使用较多的字节来表示，数值较少的数，使用较少的字节来表示。每个字节仅使用第 1 至第 7 位(共 7 bits)存储数据，第 8 位作为标识，表示是否需要继续读取下一个字节

**压缩选择**
> 根据上述两种 Block 的特点，Lucene 会每处理包含 Term 的 128 篇文档，将其对应的 DocId 数组和 TermFreq 数组分别处理为 PackedDocDeltaBlock 和 PackedFreqBlock 的 PackedInt 结构，两者组成一个 PackedBlock，最后不足 128 的文档则采用 VIntBlock 的方式来存储

###### 3.2.2. SkipData
> 搜索中存在将每个 Term 对应的 DocId 集合进行取交集的操作，即判断某个 Term 的 DocId 在另一个 Term 的 TermFreqs 中是否存在。TermFreqs 中每个 Block 中的 DocId 是有序的，可以采用顺序扫描的方式来查询，但是如果 Term 对应的 doc 特别多时搜索效率就会很低，同时由于 Block 的大小是不固定的，我们无法使用二分的方式来进行查询。因此 Lucene 为了减少扫描和比较的次数，采用了 SkipData 这个跳表结构来实现快速跳转

**跳表**
> 跳表是在原有的有序链表上面增加了多级索引，通过索引来实现快速查找。
> 实质就是一种可以进行二分查找的有序链表。


**SkipData 结构**
> 在 TermFreqs 中每生成一个 Block 就会在 SkipData 的第 0 层生成一个节点，然后第 0 层以上每隔 N 个节点生成一个上层节点。
> 每个节点通过 Child 属性关联下层节点，节点内 DocSkip 属性保存 Block 的最大的 DocId 值，DocBlockFP、PosBlockFP、PayBlockFP 则表示 Block 数据对应在 .pay、.pos、.doc 文件的位置

###### 3.2.3 Posting 数据
> Posting List 采用多个文件进行存储，每个Term包含以下内容：
- SkipOffset：用来描述当前 term 信息在 .doc 文件中跳表信息的起始位置。
- DocStartFP：是当前 term 信息在 .doc 文件中的文档 ID 与词频信息的起始位置。
- PosStartFP：是当前 term 信息在 .pos 文件中的起始位置。
- PayStartFP：是当前 term 信息在 .pay 文件中的起始位置

###### 4. 倒排索引查询
- 通过 Term Index 数据（.tip 文件）中的 StartFP 获取指定字段的 FST

- 通过 FST 找到指定 Term 在 Term Dictionary（.tim 文件）可能存在的 Block

- 将对应 Block 加载内存，遍历 Block 中的 Entry，通过后缀（Suffix）判断是否存在指定 Term

- 存在则通过 Entry 的 TermStat 数据中各个文件的 FP 获取 Posting 数据

- 如果需要获取 Term 对应的所有 DocId 则直接遍历 TermFreqs，如果获取指定 DocId 数据则通过 SkipData 快速跳转