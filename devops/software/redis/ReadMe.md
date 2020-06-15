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