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