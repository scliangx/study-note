# 虚拟机部署k8s集群

> [域名ip查询] https://ip.tool.chinaz.com/raw.githubusercontent.com

```sh
# 关闭防火墙
[root@k8s-master ~]# systemctl stop firewalld
# 禁用开机启动
[root@k8s-master ~]# systemctl disable firewalld

# 禁用交换分区
[root@k8s-master ~]# swapoff -a

# 禁用selinux
[root@k8s-master ~]# setenforce 0 

# 查看状态
[root@k8s-master ~]# getenforce 0      
[root@k8s-master ~]# vim /etc/selinux/config
#永久关闭selinux，在文档最后加下面这句
SELINUX=disabled

# 更改hostname名称
[root@k8s-master ~]# hostnamectl set-hostname $hostname

# 建立IP与主机名的映射
[root@k8s-master ~]# vim /etc/hosts

172.20.73.73 k8s-master
172.20.75.230 k8s-node1
172.20.71.151 k8s-node2

# 时间同步
[root@k8s-master ~]# yum -y install ntp
[root@k8s-master ~]# systemctl start ntpd
[root@k8s-master ~]# systemctl enable ntpd

# 将桥接的IPv4流量传递到iptables的链
[root@k8s-master ~]# touch /etc/sysctl.d/k8s.conf

[root@k8s-master ~]# cat >> /etc/sysctl.d/k8s.conf <<EOF     
net.bridge.bridge-nf-call-ip6tables=1                     
net.bridge.bridge-nf-call-iptables=1                       
net.ipv4.ip_forward=1
vm.swappiness=0
EOF

# 加载系统参数
[root@k8s-master ~]# sysctl --system

# 安装docker，k8s驱动
[root@k8s-master ~]# yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

[root@k8s-master ~]# yum list docker-ce --showduplicates | sort -r

[root@k8s-master ~]# yum -y install docker-ce docker-ce-cli containerd.io

[root@k8s-master ~]# systemctl start docker
# 设置开机启动
[root@k8s-master ~]# systemctl enable docker

[root@k8s-master ~]# vim /etc/docker/daemon.json

{
    "registry-mirrors": ["https://b9pmyelo.mirror.aliyuncs.com"],  
    "exec-opts": ["native.cgroupdriver=systemd"]           
}

[root@k8s-master ~]# systemctl restart docker
[root@k8s-master ~]# docker info |tail -5
[root@k8s-master ~]# docker info | grep -i "Cgroup Driver"

# 使用kubeadm安装k8s及相关工具
[root@k8s-master ~]# cat >/etc/yum.repos.d/kubernetes.repo <<EOF  
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

# 查看可安装版本
[root@k8s-master ~]# yum list --showduplicates | grep  kubeadm

# 安装指定版本
[root@k8s-master ~]# yum -y install kubelet-1.22.6 kubeadm-1.22.6 kubectl-1.22.6

# 开机启动
[root@k8s-master ~]# systemctl enable kubelet   

# 初始化master节点(MASTER)
[root@k8s-master ~]# kubeadm init  --apiserver-advertise-address=8.141.175.100 --image-repository registry.aliyuncs.com/google_containers --kubernetes-version v1.22.6 --service-cidr=10.96.0.0/12 --pod-network-cidr=10.244.0.0/16

[root@k8s-master ~]# mkdir -p $HOME/.kube
[root@k8s-master ~]# cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
[root@k8s-master ~]# chown $(id -u):$(id -g) $HOME/.kube/config
[root@k8s-master ~]# export KUBECONFIG=/etc/kubernetes/admin.conf

[root@k8s-master ~]# scp /etc/kubernetes/admin.conf k8s-node1:/etc/kubernetes/admin.conf
[root@k8s-master ~]# scp /etc/kubernetes/admin.conf k8s-node1:/etc/kubernetes/admin.conf

# 将node节点加入集群，要在node节点机器上执行，node结点需要安装k8s相关组件，即kubeadm init之前的操作都需要执行
[root@k8s-master ~]# kubeadm join 172.20.73.73:6443 --token 0416zv.g44s1xnloyi8xjvm --discovery-token-ca-cert-hash sha256:37349a9729af525fe5716148545561d447b0a06ff6f9e76c4bc8269c79748017 
  

# 部署容器网络，CNI网络插件

#执行下面这条命令在线配置pod网络，因为是国外网站，所以可能报错，测试去http://ip.tool.chinaz.com/网站查到
#域名raw.githubusercontent.com对应的IP，把域名解析配置到/etc/hosts文件，然后执行在线配置pod网络，多尝试几次即可成功。
[root@k8s-master ~]# echo "185.199.110.133 raw.githubusercontent.com" >> /etc/hosts

[root@k8s-master ~]# kubectl apply -f https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml

#  查看运行状态          
[root@k8s-master ~]# kubectl get pods -n kube-system
NAME                                 READY   STATUS    RESTARTS   AGE
coredns-7f6cbbb7b8-2gwkr             1/1     Running   0          7m57s
coredns-7f6cbbb7b8-lwx8s             1/1     Running   0          7m57s
etcd-k8s-master                      1/1     Running   1          8m11s
kube-apiserver-k8s-master            1/1     Running   1          8m12s
kube-controller-manager-k8s-master   1/1     Running   1          8m11s
kube-proxy-2wbds                     1/1     Running   0          4m37s
kube-proxy-9lqxn                     1/1     Running   0          7m57s
kube-proxy-jjbtr                     1/1     Running   0          4m47s
kube-scheduler-k8s-master            1/1     Running   1          8m11s

# 验证集群是否正常，要注意看对外暴露的端口号（自动生成的）
[root@k8s-master ~]# kubectl create deployment httpd --image=httpd
deployment.apps/httpd created
[root@k8s-master ~]# kubectl expose deployment httpd --port=80 --type=NodePort
service/httpd exposed
[root@k8s-master ~]# kubectl get pod,svc
NAME                         READY   STATUS    RESTARTS   AGE
pod/httpd-757fb56c8d-zq2fj   1/1     Running   0          17s

NAME                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
service/httpd        NodePort    10.106.137.56   <none>        80:31537/TCP   7s
service/kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP        9m9s

# 安装kubeboard
[root@k8s-master k8s-yaml]# kubectl apply -f https://kuboard.cn/install-script/kuboard-beta.yaml
deployment.apps/kuboard created
service/kuboard created
serviceaccount/kuboard-user created
clusterrolebinding.rbac.authorization.k8s.io/kuboard-user created
serviceaccount/kuboard-viewer created
clusterrolebinding.rbac.authorization.k8s.io/kuboard-viewer created
[root@k8s-master k8s-yaml]# kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep kuboard-user | awk '{print $1}')
Name:         kuboard-user-token-6p88l
Namespace:    kube-system
Labels:       <none>
Annotations:  kubernetes.io/service-account.name: kuboard-user
              kubernetes.io/service-account.uid: a980558e-fd23-4e0e-a757-c7bac989a34d

Type:  kubernetes.io/service-account-token

Data
====
ca.crt:     1099 bytes
namespace:  11 bytes
token:      eyJhbGciOiJSUzI1NiIsImtpZCI6IlM4cFhRbW9Yb25HdGxzSm9HUWQwXzlfeUxXZVZ5UWZ4VHVxQUxwSnl3UXcifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrdWJvYXJkLXVzZXItdG9rZW4tNnA4OGwiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoia3Vib2FyZC11c2VyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiYTk4MDU1OGUtZmQyMy00ZTBlLWE3NTctYzdiYWM5ODlhMzRkIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOmt1Ym9hcmQtdXNlciJ9.q2fnZZN1-BN5QsLeoai1vQ3IkiuIeUTGt-lncgjFBZILxDoNRRQDlvIxPIvYu4PHGpbUf295vLRc5y0CvNizWAZqenb3exkuM4ahF0mNU68B2NYBnZ3No86GiFNj-pgrKpeaECHHERZn5I3h3aKupzYHlLgDes3-UATurf3UsgpUO7Gbp6xcT8OVDK_y81JnQLwXDIKEF09CbzgjF8EOn2_Gbi-zyapZQjVV3sDaRPs9-cXLRrD1n2FmRtYVb0fSCu_V0nnn_5gIy1OlI5wEHfUgrFOtEAg6B38AYyQamzQqPT6YIehpur1C7g9c5VCriOH1aAgj8WSAB31haeuvFw
[root@k8s-master k8s-yaml]# kubectl get svc -n kube-system
NAME       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                  AGE
kube-dns   ClusterIP   10.96.0.10       <none>        53/UDP,53/TCP,9153/TCP   76m
kuboard    NodePort    10.102.159.239   <none>        80:32567/TCP             21s

# 访问ip:32567
```

#### 清除node结点环境
```sh
#!/bin/sh
kubeadm reset
systemctl stop kubelet
systemctl stop docker
rm -rf /var/lib/cni/
rm -rf /var/lib/kubelet/*
rm -rf /etc/cni/
ifconfig cni0 down
ifconfig flannel.1 down
ifconfig docker0 down
ip link delete cni0
ip link delete flannel.1
systemctl start docker
systemctl start kubelet
```
