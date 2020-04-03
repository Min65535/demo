## 日志库
配置名称：gotocash-gapi-gateway
日志路径:/data/nfs_ind_prod/qy-log/api-gateway /∗∗/ *.log
模式: 完整正则模式
单行模式:关
日志样例:2020-04-01 16:52:43.918 warn    gapigateway-deployment-548485fb6b-fkb8b    main.Ha:37     RegisterErr  {"error": "fail to register at 2020-04-01 16:52:43"}

行首正则表达式:\d+-\d+-\d+\s\d+:\d+:\d+\.\d+\s.*

提取字段:开
正则:(\d+-\d+-\d+\s\S+)\s(\w+)\s+(\S+)\s+(.*)
日志抽取内容: key|value
sys_time|2019-12-23 13:53:21.114
level|warn
thread|ind-gapigateway-deployment-548485fb6b-fkb8b
content|main.Ha:37     RegisterErr  {"error": "fail to register at 2020-04-01 16:52:43"}


## 仪表盘
请选择日志库
图表名称
显示标题边框背景
level:error
1分钟（相对）

## 告警
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

