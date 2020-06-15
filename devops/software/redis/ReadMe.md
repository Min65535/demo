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