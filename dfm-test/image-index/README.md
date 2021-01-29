## README
```text
docker network create my-net

docker run -d --restart always  --network my-net --name decode-app -p 80:80 dididi/1:latest
```