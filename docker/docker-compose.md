### docker-compose

#### 1. 什么是docker-compose
```text
Compose 是用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，您可以使用 YML 文件来配置应用程序需要的所有服务。然后，使用一个命令，就可以从 YML 文件配置中创建并启动所有服务。

Compose 使用的三个步骤：

    1) 使用 Dockerfile 定义应用程序的环境。
    2) 使用 docker-compose.yml 定义构成应用程序的服务，这样它们可以在隔离环境中一起运行。
    3) 最后，执行 docker-compose up 命令来启动并运行整个应用程序。
```
#### 2. Docker-compose 的安装
**linux**
```sh
# Linux 上我们可以从 Github 上下载它的二进制包来使用，最新发行的版本地址：https://github.com/docker/compose/releases

# 1. 下载二进制包
# 不同的版本更换v2.2.2即可
sudo curl -L "https://github.com/docker/compose/releases/download/v2.2.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# 2. 添加执行权限
sudo chmod +x /usr/local/bin/docker-compose

# 3. 创建软连接,如果第一步直接下载到/usr/bin/docker-compose 目录下的，则可以省略该步骤
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

# 4. 检查安装是否成功
docker-compose version
docker-compose version 1.25.1, build a82fef07
```
**macos & windowns**
```sh
# Mac和Windows 的 Docker 桌面版和 Docker Toolbox 已经包括 Compose 和其他 Docker 应用程序,，所以 Mac 用户不需要单独安装 docker-compose
```

#### 3. docker-compose 的基本使用

##### 3.1 docker-compose.yaml

> 创建一个docker-compose.yml文件
```yml
version: '3'
services:
  mysql:
    image: "mysql:8.0"  # 指定镜像
    ports:
     - "6379:6379"  # 端口映射
```

> 启动&停止docker-compose
```sh
# 下述操作，如果文件名为docker-compose.yml，那可以不用指定-f $file_name 参数

# 1. 直接执行
docker-compose up

# 2. 后台执行
docker-compose up -d

# 3. 如果文件名称不叫docker-compose，这可以指定文件名
docker-compose -f $file_name up

# 4. 停止容器
docker-compose -f $file_name down # 终止compose，彻底终止，包括compose的资源也会被清楚

docker-compose  -f $file_name stop # 停止当前compose，可以通过start再次启动，回到stop之前的状态

# 5. 重启compose
docker-compose -f $file_name restart # 知识重启compose

# 也可以使用down之后在up
```

**docker-compose 常见参数选项**

```sh
# 1. 指定本 yml 依从的 compose 哪个版本制定的。
version 
```
```sh
# 2. 添加或删除容器拥有的宿主机的内核功能。
cap_add:
  - ALL # 开启全部权限
cap_drop:
  - SYS_PTRACE # 关闭 ptrace权限
```

```sh
# 3. cgroup_parent 为容器指定父 cgroup 组，意味着将继承该组的资源限制
cgroup_parent: m-executor-abcd
```
```sh
# 4. command 覆盖容器启动的默认命令
command: ["bash", "echo hello"]
```

```sh
# 5. container_name 指定自定义容器名称，而不是生成的默认名称
container_name: my_redis
```
```sh
# 6. depends_on 设置依赖关系:
    docker-compose up   # 以依赖性顺序启动服务。在以下示例中，先启动 db 和 redis ，才会启动 web。
    docker-compose up SERVICE  #自动包含 SERVICE 的依赖项。在以下示例中，docker-compose up web 还将创建并启动 db 和 redis
    docker-compose stop   #按依赖关系顺序停止服务。在以下示例中，web 在 db 和 redis 之前停止。
```
```sh
# 7. endpoint_mode：访问集群服务的方式。
    # Docker 集群服务一个对外的虚拟 ip。所有的请求都会通过这个虚拟 ip 到达集群服务内部的机器。
    endpoint_mode: vip  
    # DNS 轮询（DNSRR）。所有的请求会自动轮询获取到集群 ip 列表中的一个 ip 地址。
    endpoint_mode: dnsrr   
```
```sh
# 8. 在服务上设置标签。可以用容器上的 labels（跟 deploy 同级的配置） 覆盖 deploy 下的 labels。
labels

```
```sh
# 9. mode：指定服务提供的模式。
    # 复制服务，复制指定服务到集群的机器上。
    replicated 
    #全局服务，服务将部署至集群的每个节点
    global  
```
```sh
# 10. 暴露端口，但不映射到宿主机，只被连接的服务访问
expose  
```
```sh
# 11. 指定容器运行的镜像
image   
```
```sh
# 12. restart  # 重启策略
    no：是默认的重启策略，在任何情况下都不会重启容器。
    always：容器总是重新启动。
    on-failure：在容器非正常退出时（退出状态非0），才会重启容器。
    unless-stopped：在容器退出时总是重启容器，但是不考虑在Docker守护进程启动时就已经停止了的容器
```
```sh
# 13将主机的数据卷或着文件挂载到容器里。
volumes 
```
**更多的参数选项直接使用docker-compose --help即可查看**
```sh
[root@iZ2ze58f53sxjm9z7mgn5xZ ~]# docker-compose --help
Define and run multi-container applications with Docker.

Usage:
  docker-compose [-f <arg>...] [options] [COMMAND] [ARGS...]
  docker-compose -h|--help

Options:
  -f, --file FILE             Specify an alternate compose file
                              (default: docker-compose.yml)
  -p, --project-name NAME     Specify an alternate project name
                              (default: directory name)
  --verbose                   Show more output
  --log-level LEVEL           Set log level (DEBUG, INFO, WARNING, ERROR, CRITICAL)
  --no-ansi                   Do not print ANSI control characters
  -v, --version               Print version and exit
  -H, --host HOST             Daemon socket to connect to

  --tls                       Use TLS; implied by --tlsverify
  --tlscacert CA_PATH         Trust certs signed only by this CA
  --tlscert CLIENT_CERT_PATH  Path to TLS certificate file
.............
```

#### 4. docker-compose 搭建kafka集群案例
```yml
version: "3.7"
services:
  zookeeper:
    container_name: zk
    image: wurstmeister/zookeeper
    restart: always
    volumes:
      - ./data:/data
    ports:
      - 2181:2181

  kafka1:
    container_name: kafka1
    image: wurstmeister/kafka
    restart: always
    ports:
      - 9092:9092
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://8.141.175.100:9092
      KAFKA_ADVERTISED_HOST_NAME: 8.141.175.100          # 如果设置，则就作为broker 的hostname发往producer、consumers以及其他brokers
#      KAFKA_CREATE_TOPICS: "myTopic:3" #kafka启动后初始化一个有3个partition(分区)0个副本名叫myTopic的topic
      KAFKA_ZOOKEEPER_CONNECT: zk:2181          # zookeeper集群连接地址
      KAFKA_ADVERTISED_PORT: 9092            # 此端口将给与producers、consumers、以及其他brokers，它会在建立连接时用到
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9092
      KAFKA_HEAP_OPTS: "-Xmx256M -Xms128M"
      ALLOW_PLAINTEXT_LISTENER: 'true'
    volumes:
      - ./kafka1-logs:/kafka
    depends_on:
      - zookeeper
  kafka2:
    container_name: kafka2
    image: wurstmeister/kafka
    restart: always
    ports:
      - 9093:9093
    environment:
      KAFKA_BROKER_ID: 2
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://8.141.175.100:9093
      KAFKA_ADVERTISED_HOST_NAME: 8.141.175.100
      # KAFKA_CREATE_TOPICS: "myTopic:3" #kafka启动后初始化一个有3个partition(分区)0个副本名叫myTopic的topic
      KAFKA_ZOOKEEPER_CONNECT: zk:2181
      KAFKA_ADVERTISED_PORT: 9093
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9093
      KAFKA_HEAP_OPTS: "-Xmx256M -Xms128M"
      ALLOW_PLAINTEXT_LISTENER: 'true'
    volumes:
      - ./kafka2-logs:/kafka
    depends_on:
      - zookeeper
  kafka3:
    container_name: kafka3
    image: wurstmeister/kafka
    restart: always
    ports:
      - 9094:9094
    environment:
      KAFKA_BROKER_ID: 3
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://8.141.175.100:9094
      KAFKA_ADVERTISED_HOST_NAME: 8.141.175.100
      # KAFKA_CREATE_TOPICS: "myTopic:3" #kafka启动后初始化一个有3个partition(分区)0个副本名叫myTopic的topic
      KAFKA_ZOOKEEPER_CONNECT: zk:2181
      KAFKA_ADVERTISED_PORT: 9094
      KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9094
      KAFKA_HEAP_OPTS: "-Xmx256M -Xms128M"
      ALLOW_PLAINTEXT_LISTENER: 'true'
    volumes:
      - ./kafka3-logs:/kafka
    depends_on:
      - zookeeper

# 启动：
docker-compose -f $docker-compose-file-name up -d
```

