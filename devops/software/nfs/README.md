## nfs
```text
https://www.cnblogs.com/canflyfish/p/11385613.html
假设server机器ip是192.168.1.188
centos: 
执行: yum -y install rpcbind
执行: service rpcbind start
执行: systemctl enable rpcbind
执行: ss -tnulp | grep 111
执行: yum -y install nfs-utils
编辑/etc/exports ，添加以下内容

      /data    192.168.1.0/24(rw,async)
      
启动nfs服务，systemctl start nfs ,启动后 使用rpcinfo -p 192.168.1.188 查看:
[root@mdev data]# rpcinfo -p 192.168.1.188
   program vers proto   port  service
    100000    4   tcp    111  portmapper
    100000    3   tcp    111  portmapper
    100000    2   tcp    111  portmapper
    100000    4   udp    111  portmapper
    100000    3   udp    111  portmapper
    100000    2   udp    111  portmapper
    100024    1   udp  49766  status
    100024    1   tcp  58190  status
    100005    1   udp  20048  mountd
    100005    1   tcp  20048  mountd
    100005    2   udp  20048  mountd
    100005    2   tcp  20048  mountd
    100005    3   udp  20048  mountd
    100005    3   tcp  20048  mountd
    100003    3   tcp   2049  nfs
    100003    4   tcp   2049  nfs
    100227    3   tcp   2049  nfs_acl
    100003    3   udp   2049  nfs
    100003    4   udp   2049  nfs
    100227    3   udp   2049  nfs_acl
    100021    1   udp  13451  nlockmgr
    100021    3   udp  13451  nlockmgr
    100021    4   udp  13451  nlockmgr
    100021    1   tcp  17685  nlockmgr
    100021    3   tcp  17685  nlockmgr
    100021    4   tcp  17685  nlockmgr
[root@dev data]# showmount -e localhost
Export list for localhost:
/data 192.168.1.188

执行: chown -R nfsnobody.nfsnobody /data

//----------------------------
假设client机器ip是192.168.1.189执行:
执行: yum -y install rpcbind
执行: service rpcbind start
执行: systemctl enable rpcbind
执行: ss -tnulp | grep 111
执行: yum -y install nfs-utils
执行: showmount -e 192.168.1.188

挂载至本地/data/nfs目录
执行:  mount -t nfs 192.168.1.188:/data /data/nfs
```