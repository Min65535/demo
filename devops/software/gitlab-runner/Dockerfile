FROM golang:1.17.4

RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    make git docker.io openssh-client
RUN mkdir -p ~/go/src && mkdir -p ~/go/bin && mkdir -p ~/go/pkg && mkdir -p ~/.ssh/ && mkdir -p ~/.docker/ \
    && go env -w GO111MODULE="on" \
    && go env -w GOPROXY="https://goproxy.cn" \
    && go env -w GONOPROXY="gitlab.net" \
    && go env -w GOPRIVATE="gitlab.net"

ADD id_rsa ~/.ssh
ADD id_rsa.pub ~/.ssh
ADD known_hosts ~/.ssh
ADD upx /usr/local/bin
ADD config.json ~/.docker/

