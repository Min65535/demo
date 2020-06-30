## README
```text
docker build -t dididi/1 .
docker run -d --restart always --name decode-app -p 80:80 dididi/1:latest
```