# Kubernetes设计架构

- [Kubernetes 中文文档](https://www.kubernetes.org.cn/)
- [kubernetes yaml generate](https://k8syaml.com/)

> Kubernetes集群包含有节点代理kubelet和Master组件(APIs, scheduler, etc)，一切都基于分布式的存储系统。下面这张图是Kubernetes的架构图。

![image](./image/k8s01.jpg)

## 1. Kubernetes节点架构

Kubernetes主要由以下几个核心组件组成：

- etcd保存了整个集群的状态；
- apiserver提供了资源操作的唯一入口，并提供认证、授权、访问控制、API注册和发现等机制；
- controller manager负责维护集群的状态，比如故障检测、自动扩展、滚动更新等；
- scheduler负责资源的调度，按照预定的调度策略将Pod调度到相应的机器上；
- kubelet负责维护容器的生命周期，同时也负责Volume（CVI）和网络（CNI）的管理；
- Container runtime负责镜像管理以及Pod和容器的真正运行（CRI）；
- kube-proxy负责为Service提供cluster内部的服务发现和负载均衡；

除了核心组件，还有一些推荐的Add-ons：

- kube-dns负责为整个集群提供DNS服务
- Ingress Controller为服务提供外网入口
- Heapster提供资源监控
- Dashboard提供GUI
- Federation提供跨可用区的集群
- Fluentd-elasticsearch提供集群日志采集、存储与查询

### 1.1 Master 节点
![image](./image/k8s-master.png)

### 1.2 Node 节点
![image](./image/k8s-node.png)

### 1.3 分层架构
![iamge](./image/k8s%E5%88%86%E5%B1%82.jpg)
```text
核心层：Kubernetes最核心的功能，对外提供API构建高层的应用，对内提供插件式应用执行环境

应用层：部署（无状态应用、有状态应用、批处理任务、集群应用等）和路由（服务发现、DNS解析等）

管理层：系统度量（如基础设施、容器和网络的度量），自动化（如自动扩展、动态Provision等）以及策略管理（RBAC、Quota、PSP、NetworkPolicy等）

接口层：kubectl命令行工具、客户端SDK以及集群联邦

生态系统：在接口层之上的庞大容器集群管理调度的生态系统，可以划分为两个范畴

Kubernetes外部：日志、监控、配置管理、CI、CD、Workflow、FaaS、OTS应用、ChatOps等

Kubernetes内部：CRI、CNI、CVI、镜像仓库、Cloud Provider、集群自身的配置和管理等
```

## 2. 基础组件

### 2.1 kubelet
```text
kubelet负责管理pods和它们上面的容器，images镜像、volumes、etc。
```

### 2.2 kube-proxy
```text
每一个节点也运行一个简单的网络代理和负载均衡（详见services FAQ )（PS:官方 英文）。 正如Kubernetes API里面定义的这些服务（详见the services doc）（PS:官方 英文）也可以在各种终端中以轮询的方式做一些简单的TCP和UDP传输。

服务端点目前是通过DNS或者环境变量( Docker-links-compatible 和 Kubernetes{FOO}_SERVICE_HOST 及 {FOO}_SERVICE_PORT 变量都支持)。这些变量由服务代理所管理的端口来解析。

```

### 2.3 Kubernetes控制面板
```text
Kubernetes控制面板可以分为多个部分。目前它们都运行在一个master 节点，然而为了达到高可用性，这需要改变。不同部分一起协作提供一个统一的关于集群的视图。
```

### 2.4 etcd
```text
所有master的持续状态都存在etcd的一个实例中。这可以很好地存储配置数据。因为有watch(观察者)的支持，各部件协调中的改变可以很快被察觉。
```

### 2.5 Kubernetes API Server
```text
API服务提供Kubernetes API （PS:官方 英文）的服务。这个服务试图通过把所有或者大部分的业务逻辑放到不两只的部件中从而使其具有CRUD特性。它主要处理REST操作，在etcd中验证更新这些对象（并最终存储）。
```

### 2.6 Scheduler
```text
调度器把未调度的pod通过binding api绑定到节点上。调度器是可插拔的，并且我们期待支持多集群的调度，未来甚至希望可以支持用户自定义的调度器。
```

### 2.7 Kubernetes控制管理服务器
```text
所有其它的集群级别的功能目前都是由控制管理器所负责。例如，端点对象是被端点控制器来创建和更新。这些最终可以被分隔成不同的部件来让它们独自的可插拔
```

## 3. Kubernetes 核心组件

### 3.1 pod

#### 3.1.1 pod 概念

>Pod 是 k8s 系统中可以创建和管理的最小单元，是资源对象模型中由用户创建或部署的最 小资源对象模型，也是在 k8s 上运行容器化应用的资源对象，其他的资源对象都是用来支 撑或者扩展 Pod 对象功能的，比如控制器对象是用来管控 Pod 对象的，Service 或者 Ingress 资源对象是用来暴露 Pod 引用对象的，PersistentVolume 资源对象是用来为 Pod 提供存储等等，k8s 不会直接处理容器，而是 Pod，Pod 是由一个或多个 container 组成 Pod 是 Kubernetes 的最重要概念，每一个 Pod 都有一个特殊的被称为”根容器“的 Pause 容器。Pause 容器对应的镜 像属于 Kubernetes 平台的一部分，除了 Pause 容器，每个 Pod 还包含一个或多个紧密相关的用户业务容器


- (1) Pod vs 应用 每个 Pod 都是应用的一个实例，有专用的 IP
- (2) Pod vs 容器 一个 Pod 可以有多个容器，彼此间共享网络和存储资源，每个 Pod 中有一个 Pause 容器保 存所有的容器状态， 通过管理 pause 容器，达到管理 pod 中所有容器的效果
- (3) Pod vs 节点 同一个 Pod 中的容器总会被调度到相同 Node 节点，不同节点间 Pod 的通信基于虚拟二层网 络技术实现
- (4) Pod vs Pod 普通的 Pod 和静态 Pod

#### 3.1.2 Pod 特性
- (1) 资源共享
> 一个 Pod 里的多个容器可以共享存储和网络，可以看作一个逻辑的主机。共享的如 namespace,cgroups 或者其他的隔离资源。
多个容器共享同一 network namespace，由此在一个 Pod 里的多个容器共享 Pod 的 IP 和 端口 namespace，所以一个 Pod 内的多个容器之间可以通过 localhost 来进行通信,所需要 注意的是不同容器要注意不要有端口冲突即可。不同的 Pod 有不同的 IP,不同 Pod 内的多 个容器之前通信，不可以使用 IPC（如果没有特殊指定的话）通信，通常情况下使用 Pod 的 IP 进行通信。
一个 Pod 里的多个容器可以共享存储卷，这个存储卷会被定义为 Pod 的一部分，并且可 以挂载到该 Pod 里的所有容器的文件系统上。

- (2) 生命周期短暂
> Pod 属于生命周期比较短暂的组件，比如，当 Pod 所在节点发生故障，那么该节点上的 Pod 会被调度到其他节点，但需要注意的是，被重新调度的 Pod 是一个全新的 Pod,跟之前的 Pod 没有半毛钱关系

- (3) 平坦的网络
> K8s 集群中的所有 Pod 都在同一个共享网络地址空间中，也就是说每个 Pod 都可以通过其 他 Pod 的 IP 地址来实现访问。

#### 3.1.3 Pod 的分类
- (1) 普通 Pod 普通 Pod 一旦被创建，就会被放入到 etcd 中存储，随后会被 Kubernetes Master 调度到某 个具体的 Node 上并进行绑定，随后该 Pod 对应的 Node 上的 kubelet 进程实例化成一组相 关的 Docker 容器并启动起来。在默认情 况下，当 Pod 里某个容器停止时，Kubernetes 会 自动检测到这个问题并且重新启动这个 Pod 里某所有容器， 如果 Pod 所在的 Node 宕机， 则会将这个 Node 上的所有 Pod 重新调度到其它节点上。

- (2) 静态 Pod 静态 Pod 是由 kubelet 进行管理的仅存在于特定 Node 上的 Pod,它们不能通过 API Server 进行管理，无法与 ReplicationController、Deployment 或 DaemonSet 进行关联，并且 kubelet 也无法对它们进行健康检查。

#### 3.1.4 Pod 容器的状态
> 一旦 Pod 被调度到节点上，kubelet 便开始使用容器引擎（通常是 docker）创建容器。容器有三种可能的状态：Waiting / Running / Terminated：

- Waiting： 容器的初始状态。处于 Waiting 状态的容器，仍然有对应的操作在执行，例如：拉取镜像、应用 Secrets等。
- Running： 容器处于正常运行的状态。容器进入 Running 状态之后，如果指定了 postStart hook，该钩子将被执行。
- Terminated： 容器处于结束运行的状态。容器进入 Terminated 状态之前，如果指定了 preStop hook，该钩子将被执行

#### 3.1.5 重启策略
定义 Pod 或工作负载时，可以指定 restartPolicy，可选的值有：

- Always （默认值）
- OnFailure
- Never
- restartPolicy 将作用于 Pod 中的所有容器。kubelete 将在五分钟内，按照递延的时间间隔（10s, 20s, 40s ......）尝试重启已退出的容器，并在十分钟后再次启动这个循环，直到容器成功启动，或者 Pod 被删除。

#### 3.1.6 容器组的存活期
通常，如果没有人或者控制器删除 Pod，Pod 不会自己消失。只有一种例外，那就是 Pod 处于 Scucceeded 或 Failed 的 phase，并超过了垃圾回收的时长（在 kubernetes master 中通过 terminated-pod-gc-threshold 参数指定），kubelet 自动将其删除。

### 3.2 Deployment 控制器

> Deployment为Pod和ReplicaSet提供了一个声明式定义(declarative)方法，用来替代以前的ReplicationController来方便的管理应用。典型的应用场景包括：
- 定义Deployment来创建Pod和ReplicaSet
- 滚动升级和回滚应用
- 扩容和缩容
- 暂停和继续Deployment

#### 3.2.1 deployment 基本操作
```sh
# 查看pod
[root@scliang-aliyun k8s-yaml]# kubectl get po
NAME                     READY   STATUS    RESTARTS   AGE
nginx-6799fc88d8-mrbcw   1/1     Running   0          4d17h

# 查看deployment
[root@scliang-aliyun k8s-yaml]# kubectl get deployment
NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   1/1     1            1           4d17h

# 修改deployment,并且替换，将deployment写入文件修改，也可以自己手写
[root@scliang-aliyun k8s-yaml]# kubectl get deployment -o yaml > nginx-deployment.yaml


# 修改副本数参数replicas: 3
# replace 
[root@scliang-aliyun k8s-yaml]# kubectl replace -f nginx-deployment.yaml 
deployment.apps/nginx replaced

# 更改变动成功
[root@scliang-aliyun k8s-yaml]# kubectl get po
NAME                     READY   STATUS    RESTARTS   AGE
nginx-6799fc88d8-mrbcw   1/1     Running   0          4d17h
nginx-6799fc88d8-sbpwg   1/1     Running   0          3m26s
nginx-6799fc88d8-tg6m2   1/1     Running   0          3m26s


# 替换镜像
[root@scliang-aliyun k8s-yaml]# kubectl set image deployment  nginx nginx=nginx:1.15.3 --record
deployment.apps/nginx image updated

# 查看更新过程，进行滚动更新
# 新的rs副本数增加，老的副本数量减少，最终老的rs副本数量减为0，更新完成
[root@scliang-aliyun k8s-yaml]# kubectl rollout status deployment nginx
Waiting for deployment "nginx" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "nginx" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "nginx" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "nginx" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "nginx" rollout to finish: 1 old replicas are pending termination...
deployment "nginx" successfully rolled out
```

#### 3.2.2 deployment 的回滚
```sh
# 查看更新的版本信息
[root@scliang-aliyun k8s-yaml]# kubectl rollout history deployment 
deployment.apps/nginx 
REVISION  CHANGE-CAUSE
1         <none>
2         kubectl set image deployment nginx nginx=nginx:1.15.3 --record=true


# 回滚到上一个版本
[root@scliang-aliyun k8s-yaml]# kubectl rollout undo deployment nginx
deployment.apps/nginx rolled back

# 回滚成功了，nginx默认版本
[root@scliang-aliyun k8s-yaml]# kubectl get deployment -o yaml | grep image
        - image: nginx
          imagePullPolicy: Always


# 回滚到指定版本
[root@scliang-aliyun k8s-yaml]# kubectl rollout history deployment 
deployment.apps/nginx 
REVISION  CHANGE-CAUSE
2         kubectl set image deployment nginx nginx=nginx:1.15.3 --record=true
3         <none>
4         kubectl set image deployment nginx nginx=nginx:1.15.2 --record=true
5         kubectl set image deployment nginx nginx=nginx:1.15.1 --record=true

# 查看指定版本的信息
[root@scliang-aliyun k8s-yaml]# kubectl rollout history deployment nginx --revision=4
deployment.apps/nginx with revision #4
Pod Template:
  Labels:	app=nginx
	pod-template-hash=6765bcf49c
  Annotations:	kubernetes.io/change-cause: kubectl set image deployment nginx nginx=nginx:1.15.2 --record=true
  Containers:
   nginx:
    Image:	nginx:1.15.2 # 改版本信息更改了镜像
    Port:	<none>
    Host Port:	<none>
    Environment:	<none>
    Mounts:	<none>
  Volumes:	<none>

# 回滚指定版本
[root@scliang-aliyun k8s-yaml]# kubectl rollout undo deployment nginx --to-revision=4
deployment.apps/nginx rolled back

# 查看回滚结果
[root@scliang-aliyun k8s-yaml]# kubectl get deployment -o yaml | grep image
      kubernetes.io/change-cause: kubectl set image deployment nginx nginx=nginx:1.15.2
        - image: nginx:1.15.2
          imagePullPolicy: Always
```

#### 3.2.3 deployment 扩容
```sh
# 扩容副本数量为4
[root@scliang-aliyun k8s-yaml]# kubectl scale --replicas=4 deployment nginx
deployment.apps/nginx scaled

# 查看结果
[root@scliang-aliyun k8s-yaml]# kubectl get po
NAME                     READY   STATUS    RESTARTS   AGE
nginx-6765bcf49c-4h95m   1/1     Running   0          12s
nginx-6765bcf49c-52wch   1/1     Running   0          81s
nginx-6765bcf49c-f8fb5   1/1     Running   0          99s
nginx-6765bcf49c-jqwg7   1/1     Running   0          116s

```

### 3.3 StatefulSet 控制器

- StatefulSet 概述
```text
StatefulSet 顾名思义，用于管理 Stateful（有状态）的应用程序。

StatefulSet 管理 Pod 时，确保其 Pod 有一个按顺序增长的 ID。

与 Deployment 相似，StatefulSet 基于一个 Pod 模板管理其 Pod。与 Deployment 最大的不同在于 StatefulSet 始终将一系列不变的名字分配给其 Pod。这些 Pod 从同一个模板创建，但是并不能相互替换：每个 Pod 都对应一个特有的持久化存储标识。

同其他所有控制器一样，StatefulSet 也使用相同的模式运作：用户在 StatefulSet 中定义自己期望的结果，StatefulSet 控制器执行需要的操作，以使得该结果被达成。
```

- StatefulSet 使用场景
```text
对于有如下要求的应用程序，StatefulSet 非常适用：

1) 稳定、唯一的网络标识（dnsname）
2) 每个Pod始终对应各自的存储路径（PersistantVolumeClaimTemplate）
3) 按顺序地增加副本、减少副本，并在减少副本时执行清理
按顺序自动地执行滚动更新
4) 如果一个应用程序不需要稳定的网络标识，或者不需要按顺序部署、删除、增加副本，您应该考虑使用 Deployment 这类无状态（stateless）的控制器。
```

- StatefulSet 的限制
```text
1) Pod 的存储要么由 storage class 对应的 PersistentVolume Provisioner (opens new window)提供，要么由集群管理员事先创建
2) 删除或 scale down 一个 StatefulSet 将不会删除其对应的数据卷。这样做的考虑是数据安全
3) 删除 StatefulSet 时，将无法保证 Pod 的终止是正常的。如果要按顺序 gracefully 终止 StatefulSet 中的 Pod，可以在删除 StatefulSet 前将其 scale down 到 0
4) 当使用默认的 Pod Management Policy (OrderedReady) 进行滚动更新时，可能进入一个错误状态，并需要人工介入才能修复
```

- 从上面的应用场景可以发现，StatefulSet由以下几个部分组成：
```text
 用于定义网络标志（DNS domain）的Headless Service
 用于创建PersistentVolumes的volumeClaimTemplates
 定义具体应用的StatefulSet
```

- StatefulSet中每个Pod的DNS格式为statefulSetName-{0..N-1}.serviceName.namespace.svc.cluster.local，其中
```text
1) serviceName为Headless Service的名字
2) 0..N-1为Pod所在的序号，从0开始到N-1
3) statefulSetName为StatefulSet的名字
4) namespace为服务所在的namespace，Headless Servic和StatefulSet必须在相同的namespace
5) .cluster.local为Cluster Domain
```

- 创建一个statefulset

```yaml
# 启动一个nginx的statefulset
apiVersion: v1
kind: Service
metadata:
  name:  nginx-svc
  labels:
    app: nginx-svc
spec:
  ports:
  - name: web
    port: 80
  clusterIP: None
  selector:
    app: nginx
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
spec:
  selector:
    matchLabels:
      app: nginx
  serviceName: "nginx-svc"
  replicas: 2
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80
          name: web
```


#### 3.3.1 StatefulSet 基本操作
```sh
[root@scliang-aliyun k8s-yaml]# kubectl get sts
NAME   READY   AGE
web    2/2     4m47s    # stateful ready 

# 对StatefulSet进行扩容
[root@scliang-aliyun k8s-yaml]# kubectl scale --replicas=3 sts web
statefulset.apps/web scaled


[root@scliang-aliyun k8s-yaml]# kubectl get sts
NAME   READY   AGE
web    2/3     5m44s

# 新创建一个pod，名称固定$steteful_name-{num}
[root@scliang-aliyun k8s-yaml]# kubectl get po
NAME    READY   STATUS              RESTARTS   AGE
web-0   1/1     Running             0          5m50s
web-1   1/1     Running             0          5m33s
web-2   0/1     ContainerCreating   0          8s

```

#### 3.3.3 StatefulSet 的更新策略

- On Delete
```text
OnDelete 策略实现了 StatefulSet 的遗留版本（kuberentes 1.6及以前的版本）的行为。如果 StatefulSet 的 .spec.updateStrategy.type 字段被设置为 OnDelete，当您修改 .spec.template 的内容时，StatefulSet Controller 将不会自动更新其 Pod。您必须手工删除 Pod，此时 StatefulSet Controller 在重新创建 Pod 时，使用修改过的 .spec.template 的内容创建新 Pod。
```

- Rolling Updates

```text
.spec.updateStrategy.type 字段的默认值是 RollingUpdate，该策略为 StatefulSet 实现了 Pod 的自动滚动更新。在用户更新 StatefulSet 的 .spec.tempalte 字段时，StatefulSet Controller 将自动地删除并重建 StatefulSet 中的每一个 Pod。处理顺序如下：

从序号最大的 Pod 开始，逐个删除和更新每一个 Pod，直到序号最小的 Pod 被更新

当正在更新的 Pod 达到了 Running 和 Ready 的状态之后，才继续更新其前序 Pod
```

- Partitions
```text
通过指定 .spec.updateStrategy.rollingUpdate.partition 字段，可以分片（partitioned）执行RollingUpdate 更新策略。当更新 StatefulSet 的 .spec.template 时：

序号大于或等于 .spec.updateStrategy.rollingUpdate.partition 的 Pod 将被删除重建
序号小于 .spec.updateStrategy.rollingUpdate.partition 的 Pod 将不会更新，及时手工删除该 Pod，kubernetes 也会使用前一个版本的 .spec.template 重建该 Pod
如果 .spec.updateStrategy.rollingUpdate.partition 大于 .spec.replicas，更新 .spec.tempalte 将不会影响到任何 Pod
```

### 3.4 DaemonSet 控制器
- DaemonSet 概述
```text
DaemonSet 控制器确保所有（或一部分）的节点都运行了一个指定的 Pod 副本。

1) 每当向集群中添加一个节点时，指定的 Pod 副本也将添加到该节点上
2) 当节点从集群中移除时，Pod 也就被垃圾回收了
3) 删除一个 DaemonSet 可以清理所有由其创建的 Pod
```



### 3.5 Service
```text
Kubernetes 中 Service 是一个 API 对象，通过 kubectl + YAML 或者 Kuboard，定义一个 Service，可以将符合 Service 指定条件的 Pod 作为可通过网络访问的服务提供给服务调用者。

```
Service 是 Kubernetes 中的一种服务发现机制：
>
>1. Pod 有自己的 IP 地址
>2. Service 被赋予一个唯一的 dns name
>3. Service 通过 label selector 选定一组 Pod
>4. Service 实现负载均衡，可将请求均衡分发到选定这一组 Pod 中



### 3.6 Ingress
> Ingress 公开从集群外部到集群内服务的 HTTP 和 HTTPS 路由。 流量路由由 Ingress 资源上定义的规则控制。
> Ingress 是对集群中服务的外部访问进行管理的 API 对象，典型的访问方式是 HTTP。
>
> Ingress 可以提供负载均衡、SSL 终结和基于名称的虚拟托管。

```text
Ingress 需要指定 apiVersion、kind、 metadata和 spec 字段。 Ingress 对象的命名必须是合法的 DNS 子域名名称。 关于如何使用配置文件，请参见部署应用、 配置容器、 管理资源。 Ingress 经常使用注解（annotations）来配置一些选项，具体取决于 Ingress 控制器，例如重写目标注解。 不同的 Ingress 控制器支持不同的注解。 查看你所选的 Ingress 控制器的文档，以了解其支持哪些注解。

Ingress 规约 提供了配置负载均衡器或者代理服务器所需的所有信息。 最重要的是，其中包含与所有传入请求匹配的规则列表。 Ingress 资源仅支持用于转发 HTTP(S) 流量的规则。
```

#### 3.6.1 Ingress 规则
```text
每个 HTTP 规则都包含以下信息：

可选的 host。在此示例中，未指定 host，因此该规则适用于通过指定 IP 地址的所有入站 HTTP 通信。 如果提供了 host（例如 foo.bar.com），则 rules 适用于该 host。

路径列表 paths（例如，/testpath）,每个路径都有一个由 serviceName 和 servicePort 定义的关联后端。 在负载均衡器将流量定向到引用的服务之前，主机和路径都必须匹配传入请求的内容。

backend（后端）是 Service 文档中所述的服务和端口名称的组合。 与规则的 host 和 path 匹配的对 Ingress 的 HTTP（和 HTTPS ）请求将发送到列出的 backend。
通常在 Ingress 控制器中会配置 defaultBackend（默认后端），以服务于无法与规约中 path 匹配的所有请求。
```

- 默认后端
```
没有设置规则的 Ingress 将所有流量发送到同一个默认后端，而 .spec.defaultBackend 则是在这种情况下处理请求的那个默认后端。 defaultBackend 通常是 Ingress 控制器的配置选项，而非在 Ingress 资源中指定。 如果未设置任何的 .spec.rules，那么必须指定 .spec.defaultBackend。 如果未设置 defaultBackend，那么如何处理所有与规则不匹配的流量将交由 Ingress 控制器决定（请参考你的 Ingress 控制器的文档以了解它是如何处理那些流量的）。

如果没有 hosts 或 paths 与 Ingress 对象中的 HTTP 请求匹配，则流量将被路由到默认后端。
```

- 资源后端
```text
Resource 后端是一个引用，指向同一命名空间中的另一个 Kubernetes 资源，将其作为 Ingress 对象。 Resource 后端与 Service 后端是互斥的，在二者均被设置时会无法通过合法性检查。 Resource 后端的一种常见用法是将所有入站数据导向带有静态资产的对象存储后端。
```

- 路径类型
```text
Ingress 中的每个路径都需要有对应的路径类型（Path Type）。未明确设置 pathType 的路径无法通过合法性检查。当前支持的路径类型有三种：

ImplementationSpecific：对于这种路径类型，匹配方法取决于 IngressClass。 具体实现可以将其作为单独的 pathType 处理或者与 Prefix 或 Exact 类型作相同处理。

Exact：精确匹配 URL 路径，且区分大小写。

Prefix：基于以 / 分隔的 URL 路径前缀匹配。匹配区分大小写，并且对路径中的元素逐个完成。 路径元素指的是由 / 分隔符分隔的路径中的标签列表。 如果每个 p 都是请求路径 p 的元素前缀，则请求与路径 p 匹配。
```

- 主机名通配符
```text
主机名可以是精确匹配（例如“foo.bar.com”）或者使用通配符来匹配 （例如“*.foo.com”）。 精确匹配要求 HTTP host 头部字段与 host 字段值完全匹配。 通配符匹配则要求 HTTP host 头部字段与通配符规则中的后缀部分相同。
```

- Ingress 类
```text
Ingress 可以由不同的控制器实现，通常使用不同的配置。 每个 Ingress 应当指定一个类，也就是一个对 IngressClass 资源的引用。 IngressClass 资源包含额外的配置，其中包括应当实现该类的控制器名称。
```
- 创建一个ingress
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: myingress # 指定ingress的名称
spec:
  rules:
  - host: test.com    # 解析的域名
    http:
      paths:
      - pathType: Prefix  # 路径规则
        path: "/"       # 路径
        backend:
          service:
            name: nginx   # 解析到哪一个service
            port: 
              number: 80          # 对应的端口

```

### 3.7 
