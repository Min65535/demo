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


## reload mysql slave
```text
CHANGE MASTER TO MASTER_HOST='192.168.199.8',MASTER_USER='xx',MASTER_PASSWORD='xxxxxxx',MASTER_PORT=3306,MASTER_LOG_FILE='mysql-bin.000059',MASTER_LOG_POS=545830;

start slave;

mysql> show slave status\G;
*************************** 1. row ***************************
               Slave_IO_State: Waiting for master to send event
                  Master_Host: 192.168.199.8
                  Master_User: xl
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File: mysql-bin.000060
          Read_Master_Log_Pos: 452258
               Relay_Log_File: ubuntu-relay-bin.000002
                Relay_Log_Pos: 320
        Relay_Master_Log_File: mysql-bin.000059
             Slave_IO_Running: Yes
            Slave_SQL_Running: No
              Replicate_Do_DB: 
          Replicate_Ignore_DB: 
           Replicate_Do_Table: 
       Replicate_Ignore_Table: 
      Replicate_Wild_Do_Table: 
  Replicate_Wild_Ignore_Table: 
                   Last_Errno: 1032
                   Last_Error: Could not execute Delete_rows event on table tradeserver_prod.tx_flows; Can't find record in 'tx_flows', Error_code: 1032; handler error HA_ERR_KEY_NOT_FOUND; the event's master log mysql-bin.000059, end_log_pos 546278
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 545830
              Relay_Log_Space: 6350650
              Until_Condition: None
               Until_Log_File: 
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File: 
           Master_SSL_CA_Path: 
              Master_SSL_Cert: 
            Master_SSL_Cipher: 
               Master_SSL_Key: 
        Seconds_Behind_Master: NULL
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error: 
               Last_SQL_Errno: 1032
               Last_SQL_Error: Could not execute Delete_rows event on table tradeserver_prod.tx_flows; Can't find record in 'tx_flows', Error_code: 1032; handler error HA_ERR_KEY_NOT_FOUND; the event's master log mysql-bin.000059, end_log_pos 546278
  Replicate_Ignore_Server_Ids: 
             Master_Server_Id: 1
                  Master_UUID: fc0a75d7-6f3a-11ea-9e30-f04da2736308
             Master_Info_File: /var/lib/mysql/master.info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
      Slave_SQL_Running_State: 
           Master_Retry_Count: 86400
                  Master_Bind: 
      Last_IO_Error_Timestamp: 
     Last_SQL_Error_Timestamp: 200520 11:23:02
               Master_SSL_Crl: 
           Master_SSL_Crlpath: 
           Retrieved_Gtid_Set: 
            Executed_Gtid_Set: 
                Auto_Position: 0
         Replicate_Rewrite_DB: 
                 Channel_Name: 
           Master_TLS_Version: 
1 row in set (0.00 sec)

ERROR: 
No query specified
发现主机上执行了删除操作，而从机上找不到要删除的记录。

执行命令：mysql> stop slave ;set global sql_slave_skip_counter=1;start slave;
```