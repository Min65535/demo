FROM debian:10.10-slim
RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    netbase telnet iputils-ping  net-tools procps lsof \
    curl wget vim gdb pstack tree binutils \
    && mkdir -p /match /data/config \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y


COPY min/ /min
COPY docker-entrypoint.sh /min
RUN tree /min

COPY config/ /data/config
RUN tree /data/config


WORKDIR /min

ENTRYPOINT [ "/min/docker-entrypoint.sh" ]
