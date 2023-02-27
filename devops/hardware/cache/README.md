## 清除buff/cache
Linux服务器用一段时间就会出现buff/cache内存占用过多的情况，导致free空闲内存变得非常少，严重影响使用；
```
echo 1 > /proc/sys/vm/drop_caches
echo 2 > /proc/sys/vm/drop_caches
echo 3 > /proc/sys/vm/drop_caches
```

touch cleanCache.sh

vi cleanCache.sh

```
#!/bin/bash
#每天0点清除一次缓存
echo "开始清理缓存"
sync;sync;sync #写入硬盘，防止数据丢失
sleep 10 #延迟10秒
echo 1 > /proc/sys/vm/drop_caches
echo 2 > /proc/sys/vm/drop_caches
echo 3 > /proc/sys/vm/drop_caches
echo "清理结束"
```

chmod 777 cleanCache.sh

./cleanCache.sh

下面开始设置自动定期清理

输入如下命令，打开配置文件
crontab -e
在末尾添加如下内容：（每天0点的时候执行一次，可以按需更改）
```
* 0 * * * ./路径/cleanCache.sh
```
然后输入如下指令可以查看是否成功

crontab -l

