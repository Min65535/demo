## README
```text
docker build -t publish-web/v1:stable .
docker run -d --restart always --name publish-web -p 8200:80 --network my-net --network-alias publish-web publish-web/v1:stable
```