## Harbor 部署

### 1. 安装harbor
```sh
# 安装docker
yum install -y yum-utils

yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

yum list docker-ce --showduplicate

# 设置开机自启动
systemctl enable docker

# 启动docker
systemctl start docker

# 安装docker-compose
curl -L https://github.com/docker/compose/releases/download/1.25.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose

# 安装harbor
cd /opt && mkdir -p src && cd src

wget https://github.com/goharbor/harbor/releases/download/v1.9.4/harbor-offline-installer-v1.9.4.tgz

tar xf harbor-offline-installer-v1.9.4.tgz -C /opt/

mv harbor /opt/harbor-v1.9.4

ln -s /opt/harbor-v1.9.4 /opt/harbor

# 修改harbor 配置文件
vim /opt/harbor/harbor.yml

hostname: my_ip
http:
  port: 180
data_volume: /mydata/harbor
location: /mydata/harbor/logs
mkdir -p /mydata/harbor /mydata/harbor/logs

cd /opt/harbor/

systemctl restart  docker

./install.sh # 该参数启用chart 【./install.sh --with-chartmuseum】

yum install nginx -y

# 配置nginx
vim /etc/nginx/nginx.cfg
# 在nginx上添加一个server
server {
    listen       80;
    server_name  my_ip;
    
    client_max_body_size 1000m;

    location / {
        proxy_pass http://127.0.0.1:180;
    }
}

# 浏览器访问
${my_ip}:180

# 登陆
username: admin
password: harbor配置文件中
```


### 2. 部署rancher
```sh
# 拉取镜像
docker pull rancher/rancher:v2.5.10

# 查看容器卷相关信息
docker inspect rancher/rancher:v2.5.10
"Volumes": {
    "/var/lib/cni": {},
    "/var/lib/kubelet": {},
    "/var/lib/rancher": {},
    "/var/log": {}
}

# 创建挂载卷
mkdir -p /home/rancher/var/lib/cni
mkdir -p /home/rancher/var/lib/kubelet
mkdir -p /home/rancher/var/lib/rancher
mkdir -p /home/rancher/var/log

# 启动rancher
docker run -d --restart=unless-stopped \
   -p 18081:80 -p 18443:443 \
   -v /home/rancher/var/lib/cni:/var/lib/cni \
   -v /home/rancher/var/lib/kubelet:/var/lib/kubelet \
   -v /home/rancher/var/lib/rancher:/var/lib/rancher \
   -v /home/rancher/var/log:/var/log \
   --privileged \
   --name myRancher rancher/rancher:v2.5.10


# 浏览器访问 18081 是自己映射的端口
${my_ip}:18081 
```
