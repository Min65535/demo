## README
```text
docker build -t image-web/v1:stable .
docker run -d --restart always --name image-web -p 80:80 --network my-net --network-alias image-web image-web/v1:stable
```