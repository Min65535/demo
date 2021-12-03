FROM golang:1.17.3

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase \
    make git openssh-client \
    && mkdir -p ~/go/src \
    && mkdir -p ~/go/bin \
    && mkdir -p ~/go/pkg \
    && go env -w GO111MODULE="on" \
    && go env -w GOPROXY="https://goproxy.cn" \
    && go env -w GONOPROXY="gitlab.net" \
    && go env -w GOPRIVATE="gitlab.net"
