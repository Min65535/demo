FROM golang:1.17.4

RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    make git docker.io openssh-client
RUN mkdir -p ~/go/src && mkdir -p ~/go/bin && mkdir -p ~/go/pkg \
    && go env -w GO111MODULE="on" \
    && go env -w GOPROXY="https://goproxy.cn" \
    && go env -w GONOPROXY="gitlab.net" \
    && go env -w GOPRIVATE="gitlab.net"

RUN git config --global --add url."git@gitlab.net:".insteadOf "https://gitlab.net/"
RUN docker login my.image:5000 -u user -p 123456
ADD upx /usr/local/bin

