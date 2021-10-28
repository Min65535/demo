## busybox-curl
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busyboxplus
  namespace: default
spec:
  containers:
    - name: busyboxplus
      image: radial/busyboxplus:curl
      command:
        - sleep
        - "8640000"
      imagePullPolicy: IfNotPresent
  restartPolicy: Always
```


## redis
```yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: redis-master
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis-master
  template:
    metadata:
      labels:
        app: redis-master
    spec:
      containers:
      - name: redis-master
        image: "redis_5.0.8"
        imagePullPolicy: IfNotPresent
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
---
apiVersion: v1
kind: Service
metadata:
  name: redis-master
  namespace: default
spec:
  selector:
    app: redis-master
  ports:
    - protocol: TCP
      port: 6379
```

## fastdfs
```yaml
apiVersion: v1
kind: Service
metadata:
  name: fastdfs
spec:
  type: NodePort
  ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
      nodePort: 20000
---
apiVersion: v1
kind: Endpoints
metadata:
  name: fastdfs
subsets:
  - addresses:
      - ip: 172.20.3.55
    ports:
      - port: 8080
```