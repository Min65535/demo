## registry
docker pull registry

docker run -di --restart=always --name=registry -p 5000:5000 -v /data:/var/lib/registry registry