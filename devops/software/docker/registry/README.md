#### docker-registry
```text
centos7(ubuntu同理)
执行: yum install -y httpd-tools
执行: htpasswd -B auth/htpasswd user

执行: docker run -d \
  -p 5000:5000 \
  --restart=always \
  --name registry \
  -v /data/auth:/auth \
  -v /data:/var/lib/registry \
  -e "REGISTRY_AUTH=htpasswd" \
  -e "REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm" \
  -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd \
  registry
  

测试登录: docker login my.image:5000 -u user -p 123456
```