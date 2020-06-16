##  customized installation of redis by source code in ubuntu
```text
apt install gcc

apt install make

wget http://download.redis.io/releases/redis-5.0.8.tar.gz
tar xzf redis-5.0.8.tar.gz
cd redis-5.0.8


mkdir -p /usr/lib/systemd/system
mkdir -p /usr/local/redis


vi /usr/lib/systemd/system/redis.service

[Unit]
Description=Redis Server
After=network.target

[Service]
ExecStart=/usr/local/redis/redis-5.0.8/src/redis-server  /usr/local/redis/redis-5.0.8/redis.conf  --daemonize no
ExecStop=/usr/local/redis/redis-5.0.8/src/redis-cli -p 6379 shutdown
Restart=always

[Install]
WantedBy=multi-user.target


systemctl daemon-reload

systemctl start redis.service
systemctl status redis.service
systemctl enable redis.service

ln -s /usr/local/redis/redis-5.0.8/src/redis-cli /usr/bin/redis-cli

redis-cli --version
```

## 修改配置文件
```text
# vim /etc/redis.conf
bind 127.0.0.1  //注释掉此行，不然只能本地访问 #bind 127.0.0.1
protected-mode yes    保护模式修改为no #product-mode no (关闭protected-mode模式,此时外部网络可以直接访问;开启protected-mode保护模式,需配置bind ip或者设置访问密码)
requirepass 123456 修改默认密码，查找 requirepass foobared 将 foobared 修改为你的密码

输入用户名密码 (auth  123456)
#127.0.0.1:6379> auth 123456
查看信息
#127.0.0.1:6379> info
查看所有keys
#127.0.0.1:6379> keys *
```

## redis 备份和恢复
```text
redis 数据的路径(其实备份和恢复就是cp)

导出：Redis写数据时先写到一个temp文件中，然后再把temp文件重命名为预定义的文件，所以即使Redis在运行，也可以直接用cp命令拷贝这个文件。

恢复：关闭Redis后直接覆盖掉demo.rdb，然后重启即可。

通过配置文件查看
#vim /etc/redis.conf
dbfilename dump.rdb
dir var/lib/redis
​
​
进行备份：
# redis-cli -h 192.168.25.65 -p 6379 -a "123456"
192.168.25.65:6379> ping
PONG
192.168.25.65:6379> SAVE //该命令将在 redis 安装目录中创建dump.rdb文件。
OK
​
进行恢复：
192.168.25.65:6379> CONFIG GET dir   //命令 CONFIG GET dir 输出的 redis 备份目录为 /var/lib/redis。
1) "dir"
2) "/var/lib/redis"
```

## 集群搭建
```text
1. 创建6个Redis配置文件，并配置
# pwd
/usr/local/redis/redis-5.0.8
#mkdir 6380 6381 6382 6383 6384 6385
​
然后拷贝原来的配置文件到目录下，然就在到相应目录下进行重命名
#cp redis.conf 6380   mv redis.conf redis-6380.conf
#cp redis.conf 6381   mv redis.conf redis-6381.conf
#cp redis.conf 6382   mv redis.conf redis-6382.conf
#cp redis.conf 6383   mv redis.conf redis-6383.conf
#cp redis.conf 6384   mv redis.conf redis-6384.conf
#cp redis.conf 6385   mv redis.conf redis-6385.conf
​
修改所有的配置文件 (加端口号的以此类推)
具体修改：
(1)port 6380  #绑定端口
​
(2)bind 192.168.25.64  #定IP也(可以注释掉，这样任何桌面客户端都可以连接了)
​
(3)dir /usr/local/redis-cluster/3680 #指定文件存放路径 ( .rdb .aof nodes-xxxx.conf 这样的文件都会在此路径下)
​
(4)cluster-enabled yes   #启动集群模式 
​
(5)cluster-config-file #集群节点配置文件
​
(6)daemonize yes   #后台启动
​
(7)cluster-node-timeout 5000  #指定集群节点超时时间
​
(8)appendonly yes #指定持久化方式
   
(9)protected-mode no #非保护模式

2. 启动节点
# ./src/redis-server 6380/redis-6380.conf 
# ./src/redis-server 6381/redis-6381.conf 
# ./src/redis-server 6382/redis-6382.conf 
# ./src/redis-server 6383/redis-6383.conf 
# ./src/redis-server 6384/redis-6384.conf 
# ./src/redis-server 6385/redis-6385.conf

3. 启动集群
#./src/redis-cli --cluster create 192.168.25.64:6380 192.168.25.64:6381 192.168.25.64:6382 192.168.25.64:6383 192.168.25.64:6384 192.168.25.64:6385 --cluster-replicas 1    
// --replicas 1 表示我们希望为集群中的每个主节点创建一个从节点。(--cluster-replicas 1  命令的意思： 一主一从配置，六个节点就是 三主三从)

4. 客户端连接集群,查看集群信息
# ./redis-cli -c -h 192.168.25.64 -p 6380
192.168.25.64:6380>

查看集群信息
192.168.25.64:6380> cluster info
cluster_state:ok         //集群状态
cluster_slots_assigned:16384   槽分配
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:6
cluster_size:3
cluster_current_epoch:6
cluster_my_epoch:1
cluster_stats_messages_ping_sent:7977
cluster_stats_messages_pong_sent:8091
cluster_stats_messages_sent:16068
cluster_stats_messages_ping_received:8086
cluster_stats_messages_pong_received:7977
cluster_stats_messages_meet_received:5
cluster_stats_messages_received:16068
192.168.25.64:6380> 
​
查看节点信息
192.168.25.64:6380> cluster nodes
0853d9773fdcdad1bc0174d64990c40c39b4a2c7 192.168.25.64:6384@16384 slave d9da56a8f068d5529a7771addf586d14e14ca888 0 1547453355000 5 connected
d9da56a8f068d5529a7771addf586d14e14ca888 192.168.25.64:6381@16381 master - 0 1547453356838 2 connected 5461-10922
4b8e4fa926486c8ab675b7bf2fe36320657f1eae 192.168.25.64:6380@16380 myself,master - 0 1547453354000 1 connected 0-5460
0e585e6a0f45e293680e7aec052d942d85dd90c3 192.168.25.64:6382@16382 master - 0 1547453356000 3 connected 10923-16383
8d9e75800b6a4cc1d2815fbb8dd0036ecd0221ce 192.168.25.64:6385@16385 slave 0e585e6a0f45e293680e7aec052d942d85dd90c3 0 1547453355837 6 connected
53d3770f2ee72b3548b8a4cb26a23e617f7558bb 192.168.25.64:6383@16383 slave 4b8e4fa926486c8ab675b7bf2fe36320657f1eae 0 1547453356000 4 connected
192.168.25.64:6380> 


5. 集群测试

[root@LWJ01 src]# ./redis-cli -c -h 192.168.25.64 -p 6380
192.168.25.64:6380> set name ww
-> Redirected to slot [5798] located at 192.168.25.64:6381
OK
192.168.25.64:6381> set age 20
-> Redirected to slot [741] located at 192.168.25.64:6380
OK
192.168.25.64:6380> get name
-> Redirected to slot [5798] located at 192.168.25.64:6381
"ww"
192.168.25.64:6381> get age
-> Redirected to slot [741] located at 192.168.25.64:6380
"20"
192.168.25.64:6380> 
​
​
# ./redis-cli  -h 192.168.25.64 -p 6380
192.168.25.64:6380> set name ww
(error) MOVED 5798 192.168.25.64:6381
192.168.25.64:6380> set age 20
OK
192.168.25.64:6380> get name
(error) MOVED 5798 192.168.25.64:6381
192.168.25.64:6380> get age
"20"
192.168.25.64:6380>
```

## 集群新增节点
```text
## pwd
/usr/local/redis/redis-5.0.8
# mkdir 6386 6388
#cp redis.conf 6386     mv redis.conf redis-6386.conf
#cp redis.conf 6388     mv redis.conf redis-6388.conf
//进入到相应的目录下重命名，并进行配置

启动新增加的节点
# ./src/redis-server 6386/redis-6386.conf
​
进行查看
# ps -ef | grep redis


添加节点到集群
#./src/redis-cli --cluster add-node 192.168.25.64:6386 192.168.25.64:6386
#./src/redis-cli --cluster add-node 192.168.25.64:6388 192.168.25.64:6388
格式：
redis-cli --cluster add-node  {新节点IP}:{新节点端口} {任意集群节点IP}:{对应端口} #如果添加集群中的主节点，则新添加的就是主节点，如果是从节点则是从节点
​
查看集群信息
# ./src/redis-cli -c -h 192.168.25.64 -p 6388
192.168.25.64:6386> cluster nodes

或者用这样可以查看集群检查集群
#./src/redis-cli --cluster check 192.168.25.64:6380

自定义分配槽
从添加主节点输出信息和查看集群信息中可以看出，我们已经成功的向集群中添加了一个主节点，但是这个主节还没有成为真正的主节点，因为还没有分配槽(slot)，也没有从节点，现在要给它分配槽(slot)

自定义分配槽(slot)
#./redis-cli --cluster  reshard 192.168.25.64:6380
```