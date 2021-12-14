## registry
```text
docker pull registry

docker run -di --restart=always --name=registry -p 5000:5000 -v /data:/var/lib/registry registry
```

## daemon
```text
{
  "registry-mirrors": ["https://dockerhub.azk8s.cn", "https://docker.mirrors.ustc.edu.cn"],
  "insecure-registries": ["127.0.0.1/8", "my.image:5000"],
  "max-concurrent-downloads": 10,
  "log-driver": "json-file",
  "log-level": "warn",
  "log-opts": {
    "max-size": "1g",
    "max-file": "7"
  },
  "data-root": "/data/var/lib/docker"
}
```