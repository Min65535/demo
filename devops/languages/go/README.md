## windows系统下ubuntu20.04子系统，go开发环境安装流程
```text
1. ubuntu20.04子系统
打开设置--->更新和安全--->开发者选项--->开发者人员模式打开，文件资源管理器应用；
打开控制面板--->程序--->启用或更换比windows功能--->适用于Linux的Windows子系统,勾选；
在micorosoft store搜索ubuntu20.04，安装后重启电脑，再次打开ubuntu20.04将进入安装步骤；
安装完后，更换ubuntu20.04的阿里软件源；
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
终端执行: go env -w GONOPROXY="gitlab-vywrajy.micoworld.net"

再次重启电脑

3. 其他tools工具&e.t.c
终端执行: sudo apt-get install gcc
终端执行: sudo apt-get install cmake
终端执行: sudo apt-get install unzip
终端执行: sudo apt-get install ssh-tools
终端执行: sudo apt-get install net-tools
终端执行: git config --global credential.helper store (长期存储https方式clone的账户信息)

4. go package依赖系列
安装protoc v3.11.2
https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
解压后放于/user/local目录下

kratos涉及到的protoc-grpc包管理
执行:
go 1.16版本以前(go get有提示会废弃，但可使用)
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
go get -u google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


go 1.17版本以后(go get不能使用)
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
go install  github.com/go-kratos/kratos/cmd/kratos/v2@latest
go install  google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install  google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install  github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go install  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2@latest

执行:
cd $GOBIN
sudo cp kratos protoc-gen-* /usr/local/bin
```

5. windows系统下软件
   编辑器采用vscode，扩展搜索wsl，配置默认ubuntu20.04(wsl)为默认终端;

windows redis下载地址: https://github.com/tporadowski/redis/releases

windows mysql version: https://cdn.mysql.com/archives/mysql-installer/mysql-installer-community-5.7.32.0.msi
```