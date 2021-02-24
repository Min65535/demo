## README
```text
docker build -t image-back/v1:stable .
docker run -d --restart always --name image-back -p 9527:9527 image-back/v1:stable
```