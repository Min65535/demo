## README
```text
docker build -t dididi/1 .
docker run -d --name decode-app -p 80:80 dididi/1:latest
```