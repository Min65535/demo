## README
```text
go build -o publish
docker network create my-net
docker build -t publish-back/v1:stable .
docker run -d --restart always --name publish-back -p 10086 --network my-net --network-alias publish-back publish-back/v1:stable
```