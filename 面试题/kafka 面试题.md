## Kafka


### 1. 如何解决生产者重复发送消息？
> 要启动kafka的幂等性，无需修改代码，默认为关闭，需要修改配置文件:enable.idempotence=true 同时要求 ack=all 且 retries>1。

**幂等原理：**
>每个producer有一个producer id，服务端会通过这个id关联记录每个producer的状态，每个producer的每条消息会带上一个递增的sequence，服务端会记录每个producer对应的当前最大sequence，producerId + sequence ，如果**新的消息带上的sequence不大于当前的最大sequence就拒绝这条消息**，如果消息落盘会同时更新最大sequence，这个时候重发的消息会被服务端拒掉从而避免消息重复。该配置同样应用于kafka事务中。


1. ack=0，不重试
producer发送消息完，不管结果了，如果发送失败也就丢失了。

2. ack=1，leader crash
producer发送消息完，只等待lead写入成功就返回了，leader crash了，这时follower没来及同步，消息丢失。

3. ack=all / -1
producer发送消息完，等待ollower同步完再返回，如果异常则重试。

4. 失败的offset单独记录
producer发送消息，会自动重试，遇到不可恢复异常会抛出，这时可以捕获异常记录到数据库或缓存，进行单独处理。

### 2. 消费者重复消费数据？

> 原因: 数据消费完没有及时提交offset到broke。
> 消息消费端在消费过程中挂掉没有及时提交offset到broke，另一个消费端启动拿之前记录的offset开始消费，由于offset的滞后性可能会导致新启动的客户端有少量重复消费。

**解决方案：**
1. 取消自动自动提交
每次消费完或者程序退出时手动提交。这可能也没法保证一条重复。

2. 下游做幂等
一般的解决方案是让下游做幂等或者尽量每消费一条消息都记录offset，对于少数严格的场景可能需要把offset或唯一ID,例如订单ID和下游状态更新放在同一个数据库里面做事务来保证精确的一次更新或者在下游数据表里面同时记录消费offset，然后更新下游数据的时候用消费位点做乐观锁拒绝掉旧位点的数据更新。

 