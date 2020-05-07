## 应用分发平台

https://github.com/rock-app/fabu.love

#### 1.克隆
```
git clone https://github.com/HeadingMobile/fabu.love.git
cd docker
docker-compose up -d --build
// 访问 host:9898
```

#### 2. 注册账户

#### 3. 关闭注册功能并重新启动
```
修改 Dockerfile，在 CMD 前加

ENV FABU_ALLOW_REGISTER false

重新 up，需要 sudo 

sudo docker-compose up -d --build
```

#### 4. 关闭MongoDB外网端口
```
修改docker-compose.yml,将端口注释

  #ports:
     # - "27017:27017"
```