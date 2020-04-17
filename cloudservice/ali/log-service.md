## 日志库

#### service
配置名称：gapi-gateway
日志路径:/data/nfs_prod/xx-log/api-gateway /∗∗/ *.log
模式: 完整正则模式
单行模式:关
日志样例:2020-04-01 16:52:43.918 warn    gapigateway-deployment-548485fb6b-fkb8b    main.Ha:37     RegisterErr  {"error": "fail to register at 2020-04-01 16:52:43"}

行首正则表达式:\d+-\d+-\d+\s\d+:\d+:\d+\.\d+\s.*

提取字段:开
正则:(\d+-\d+-\d+\s\S+)\s(\w+)\s+(\S+)\s+(.*)
日志抽取内容: key|value
sys_time|2019-12-23 13:53:21.114
level|warn
thread|gapigateway-deployment-548485fb6b-fkb8b
content|main.Ha:37     RegisterErr  {"error": "fail to register at 2020-04-01 16:52:43"}

#### mysql-log
配置名称：slow-log-node
日志路径:/data/nfs_prod/mysql/slow/node /∗∗/ *.log
模式: 完整正则模式
单行模式:关
日志样例:
```
# Time: 2020-04-16T07:00:00.942128Z
# User@Host: root[root] @  [192.168.1.3]  Id:   12
# Query_time: 0.008652  Lock_time: 0.000323 Rows_sent: 1  Rows_examined: 6783
use gapi_prod;
SET timestamp=1587020400;
SELECT
    unix_timestamp() AS unix_time, user.user, txacc.gsl, txacc.bb_gsl
FROM
    (SELECT SUM(balance) AS gsl,
            SUM(IF(account_type = 0, balance, 0)) AS bb_gsl,
                        SUM(IF(account_type = 1, balance, 0)) AS ct_gsl
    FROM tradeserver_prod.accounts) txacc,
    (SELECT max(id) AS user FROM users_prod.users) user;
```

行首正则表达式:#\sTime:\s\S+

提取字段:开
正则:(#\s\S+\s\S+)((.|\n+)+)
日志抽取内容: key|value
```
sys_time|# Time: 2020-04-16T07:00:00.942128Z#
content|# User@Host: root[root] @  [192.168.1.3]  Id:   12
# Query_time: 0.008652  Lock_time: 0.000323 Rows_sent: 1  Rows_examined: 6783
use gapi_prod;
SET timestamp=1587020400;
SELECT
    unix_timestamp() AS unix_time, user.user, txacc.gsl, txacc.bb_gsl
FROM
    (SELECT SUM(balance) AS gsl,
            SUM(IF(account_type = 0, balance, 0)) AS bb_gsl,
                        SUM(IF(account_type = 1, balance, 0)) AS ct_gsl
    FROM tradeserver_prod.accounts) txacc,
    (SELECT max(id) AS user FROM users_prod.users) user;
```

## 仪表盘

#### service
请选择日志库
图表名称
显示标题边框背景
level:error
1分钟（相对）

#### mysql-log
请选择日志库
图表名称
显示标题边框背景
sys_time
1分钟（相对）

## 告警

#### service
关联图表：查询区间1分钟（相对）
频率：固定间隔 1 分钟
触发条件：contains(level,"error")
触发通知阈值：1
通知间隔：无间隔

通知列表
 *短信
  ** 手机号码：
  ** 发送内容：${Results[0].FireResult}
 *邮件
  ** 收件人：
  ** 主题：
  ** 发送内容：${Project}, ${Condition}, ${AlertName}, ${AlertID}, ${Dashboard}, ${FireTime}, ${Results}

#### mysql-log
关联图表：查询区间1分钟（相对）
频率：固定间隔 1 分钟
触发条件：contains(sys_time,"Time")
触发通知阈值：1
通知间隔：无间隔

通知列表
 *短信
  ** 手机号码：
  ** 发送内容：${Results[0].FireResult}
 *邮件
  ** 收件人：
  ** 主题：
  ** 发送内容：${Project}, ${Condition}, ${AlertName}, ${AlertID}, ${Dashboard}, ${FireTime}, ${Results}
