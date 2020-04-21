## change the log dir
SHOW  GLOBAL VARIABLES LIKE '%slow%';
set global slow_query_log=1
set global slow_query_log_file='/data/mysql-slow/dfm-pc-slow.log'

/var/lib/mysql/dfm-pc-slow.log
set global slow_query_log_file='/var/lib/mysql/dfm-pc-slow.log'

ERROR 29 (HY000): File (Errcode: 13 - Permission denied)

/etc/apparmor.d/usr.sbin.mysqld

/data/ r,
/data/* rw,

/data/mysql-slow r,
/data/mysql-slow/* rw,

sudo /etc/init.d/apparmor reload

## using_indexes
mysql 慢查询日志导致时间无效的假象
```
[mysqld] 
log_slow_queries=ON 
long_query_time=5 
slow_query_log=ON 
#log_queries_not_using_indexes=ON 
slow_query_log_file=D:/phpStudy/MySQL/slow.log
```