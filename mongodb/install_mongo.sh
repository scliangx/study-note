#!/bin/sh

# download mongo.tar.gz
curl -O https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-3.2.9.tgz

tar zxvf mongodb-linux-x86_64-3.2.9.tgz

# move software
mv mongodb-linux-x86_64-3.2.9/ /usr/local/mongodb

# create data/log/conf directory
mkdir -p  /usr/local/mongodb/data
touch /usr/local/mongodb/mongod.log
touch /usr/local/mongodb/mongodb.conf

# modify config
vim /usr/local/mongodb/mongodb.conf


dbpath=/usr/local/mongodb/data
logpath=/usr/local/mongodb/mongod.log
logappend = true 
port = 27017 
fork = true 
auth = true

# create mongo.service
vim /usr/lib/systemd/system

[Unit]
Description=mongodb
After=network.target remote-fs.target nss-lookup.target
[Service]
Type=forking
ExecStart=/usr/local/mongodb/bin/mongod --config /usr/local/mongodb/mongodb.conf
ExecReload=/bin/kill -s HUP $MAINPID
ExecStop=/usr/local/mongodb/bin/mongod -shutdown -dbpath=/usr/local/mongodb/data
PrivateTmp=true
[Install]
WantedBy=multi-user.targe

# $PATH
vim ~/.bashrc
export PATH=$PATH:/usr/local/mongodb/bin

source ~/.bashrc

/sbin/iptables -I INPUT -p tcp --dport 27017 -j ACCEPT

# start
systemctl start mongo.servie

# stop
systemctl stop mongo.service

