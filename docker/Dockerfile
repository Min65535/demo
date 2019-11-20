FROM basesystem:1.0
MAINTAINER qy

# 时区设置成当前时区
RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN mkdir -p /usr/local/.db
# 这个密码要跟db中的一致
RUN echo hello_qy > /usr/local/.db/mysql.pas
RUN echo qy > /usr/local/.db/mysql.uname

COPY start_app.sh /
RUN chmod +x /start_app.sh

COPY tmp/demo_db /usr/bin
COPY tmp/demo /usr/bin

ENTRYPOINT /start_app.sh
