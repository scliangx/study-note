## ETCD 

> 一个用于存储分布式系统中最关键的数据的仓库，它是分布式的、可靠的键值对仓库。
>
> etcd 是一个一致的分布式键值存储。在分布式系统中主要用作单独的协调服务。并且旨在保存可以完全放入内存的少量数据

### 1. ETCD 的组成

> etcd架构图

![image](F:\GoProjects\src\MyPractiseNotes\etcd\image\etcd.png)

**etcd 的特点**

```text
简单：
    - 易使用：基于HTTP+JSON的API让你用curl就可以轻松使用；
    - 易部署：使用Go语言编写，跨平台，部署和维护简单。

可靠：
    - 强一致：使用Raft算法充分保证了分布式系统数据的强一致性；
    - 高可用：具有容错能力，假设集群有n个节点，当有(n-1)/2节点发送故障，依然能提供服务；
    - 持久化：数据更新后，会通过WAL格式数据持久化到磁盘，支持Snapshot快照。

快速：每个实例每秒支持一千次写操作，极限写性能可达10K QPS。

安全：可选SSL客户认证机制
```



**etcd主要分为四个部分：**

> 1. HTTP Server： 用于处理用户发送的API请求以及其它etcd节点的同步与心跳信息请求。
> 2. Store：用于处理etcd支持的各类功能的事务，包括数据索引、节点状态变更、监控与反馈、事件处理与执行等等，是etcd对用户提供的大多数API功能的具体实现。
> 3. Raft：Raft强一致性算法的具体实现，是etcd的核心。
> 4. WAL：Write Ahead Log（预写式日志），是etcd的数据存储方式。除了在内存中存有所有数据的状态以及节点的索引以外，etcd就通过WAL进行持久化存储。WAL中，所有的数据提交前都会事先记录日志。Snapshot是为了防止数据过多而进行的状态快照；Entry表示存储的具体日志内容。

```text
通常，一个用户的请求发送过来，会经由HTTP Server转发给Store进行具体的事务处理，如果涉及到节点的修改，则交给Raft模块进行状态的变更、日志的记录，然后再同步给别的etcd节点以确认数据提交，最后进行数据的提交，再次同步。
```

### 2. 主从数据同步

```text
1. client连接follower或者leader，如果连接的是follower则，follower会把client的请求(写请求，读请求则自身就可以直接处理)转发到leader

2. leader接收到client的请求，将该请求转换成entry，写入到自己的日志中，得到在日志中的index，会将该entry发送给所有的follower(实际上是批量的entries)

3. follower接收到leader的AppendEntriesRPC请求之后，会将leader传过来的批量entries写入到文件中（通常并没有立即刷新到磁盘），然后向leader回复OK,leader收到过半的OK回复之后，就认为可以提交了，然后应用到leader自己的状态机中，leader更新commitIndex，应用完毕后回复客户端

4. 在下一次leader发给follower的心跳中，会将leader的commitIndex传递给follower，follower发现commitIndex更新了则也将commitIndex之前的日志都进行提交和应用到状态机中
```

> 在leader收到数据操作的请求，先不着急更新本地数据（数据是持久化在磁盘上的），而是生成对应的log，然后把生成log的请求广播给所有的follower，每个follower在收到请求之后听从leader的命令，也写入log，然后返回success回去。eader收到过半的OK回复之后，就认为可以提交了，**进行二次提交**，正式写入数据（持久化），然后再告诉follower，他们也持久化

### 3. 选举过程

```text
1. 集群初始化时候，每个节点都是Follower角色；都维护一个随机的timer，如果timer时间到了还没有收到leader的消息，自己就会变成candidate，竞选leader，

2. 当Follower在一定时间内没有收到来自主节点的心跳，会将自己角色改变为Candidate，并发起一次选主投票；当收到包括自己在内超过半数节点赞成后，选举成功；当收到票数不足半数选举失败，或者选举超时。若本轮未选出主节点，将进行下一轮选举（出现这种情况，是由于多个节点同时选举，所有节点均为获得过半选票）。

3. 集群中存在至多1个有效的主节点，通过心跳与其他节点同步数据

4. Candidate节点收到来自主节点的信息后，会立即终止选举过程，进入Follower角色。为了避免陷入选主失败循环，每个节点未收到心跳发起选举的时间是一定范围内的随机值，这样能够避免2个节点同时发起选主
```

