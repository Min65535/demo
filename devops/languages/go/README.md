## windows系统下ubuntu18.04子系统，go开发环境安装流程
```text
1. ubuntu18.04子系统
打开设置--->更新和安全--->开发者选项--->开发者人员模式打开，文件资源管理器应用；
打开控制面板--->程序--->启用或更换比windows功能--->适用于Linux的Windows子系统,勾选；
在micorosoft store搜索ubuntu18.04，安装后重启电脑，再次打开ubuntu18.04将进入安装步骤；
安装完后，更换ubuntu18.04的阿里软件源；
终端执行: sudo cp /etc/apt/sources.list /etc/apt/sources.list.bak
终端执行: sudo echo "">/etc/apt/sources.list
终端执行: sudo vi /etc/apt/sources.list
插入:
deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse

deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse

deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse

deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse

deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse

deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
保存

终端执行: sudo apt update
终端执行: sudo apt upgrade

2. go环境
在ubuntu子系统配置go环境，安装go的sdk于windows系统下的D盘,languages/linux_go目录下,即ubuntu系统下/mnt/d/languages/linux_go目录；
在终端执行: mkdir -p /mnt/d/languages/linux_go       (注意: 这里的languages/linux_go是为了GOROOT)
在终端执行: mkdir -p /mnt/d/code/go/src&&mkdir -p /mnt/d/code/go/bin       (注意: 这里的go目录是为了GOPATH,这里的src就是项目目录,bin目录是GOBIN需要的目录)
并在终端执行: vi ~/.bashrc
写入:
export GOROOT=/mnt/d/languages/linux_go/go
export GOPATH=/mnt/d/code/go
export PATH=$PATH:$GOROOT/bin
export GOBIN=$GOPATH/bin

然后再在终端执行: source ~/.bashrc

终端执行: go env
终端执行: go version
终端执行: go env -w GO111MODULE="on"
终端执行: go env -w GOPROXY="https://goproxy.cn"
终端执行: go env -w GONOPROXY="gitlab.net"
终端执行: go env -w GOPRIVATE="gitlab.net"

再次重启电脑

3. 其他tools工具&e.t.c
终端执行: sudo apt-get install gcc
终端执行: sudo apt-get install cmake
终端执行: sudo apt-get install unzip
终端执行: sudo apt-get install openssh-server
终端执行: sudo apt-get install net-tools
终端执行: sudo apt-get install mysql-server
终端执行: sudo service mysql start
终端执行: sudo apt-get install redis
终端执行: sudo service redis-server start
终端执行: sudo apt-get install docker.io
终端执行: git config --global credential.helper store (长期存储https方式clone的账户信息)

4. 对于电脑重新开机情况ubuntu18.04子系统本地mysql/redis等服务的重启
终端执行: cd ~
终端执行: local-db.sh
输入:
#!/usr/bin/env bash

service mysql start
service redis-server start
保存shell内容

终端执行: chmod +x local-db.sh
终端执行: sudo su
终端执行: chown root:root local-db.sh
终端执行: cp local-db.sh ~

以后每次电脑重新开机均须在终端切换到root用户,去执行这个local-db.sh文件来启动这些服务


5. go package依赖系列
安装protoc v3.11.2
https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
解压后放于/user/local目录下
举例执行: wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
执行: sudo mv protoc-3.11.2-linux-x86_64.zip /usr/local
执行: cd /usr/local
执行: sudo unzip protoc-3.11.2-linux-x86_64.zip
执行: sudo chmod +x /usr/local/bin/protoc

kratos涉及到的protoc-grpc包管理
执行:
go 1.16版本以前(go get有提示会废弃，但可使用)
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


go 1.17版本以后(go get不能使用)
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

执行:
cd $GOBIN
sudo cp kratos protoc-gen-* /usr/local/bin

对于protoc-gen-grpc-gateway，需要去官方拉最新的可执行文件
https://github.com/grpc-ecosystem/grpc-gateway/releases，然后部署在/usr/local/bin
举例执行: wget -O protoc-gen-grpc-gateway  https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v2.7.0/protoc-gen-grpc-gateway-v2.7.0-linux-x86_64
执行: chmod +x protoc-gen-grpc-gateway
执行: sudo cp protoc-gen-grpc-gateway /usr/local/bin


6. windows系统下软件
编辑器若采用vscode，扩展搜索wsl，配置默认ubuntu18.04(wsl)为默认终端;
编辑器若采用goland，在setting/tools/terminal，配置shell path: wsl.exe --distribution Ubuntu-18.04


ubuntu子系统docker集成参考:
https://docs.microsoft.com/zh-cn/windows/wsl/install-manual#step-4---download-the-linux-kernel-update-package
https://docs.docker.com/desktop/windows/wsl/
```