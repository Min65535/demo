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

## pv-pvc-nfs
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: test-pvc  #TODO: give right name of nfs pv
  namespace: test
spec:
  capacity:
    storage: 100Gi #TODO: give size of this pv
  accessModes:
    - ReadWriteMany
  nfs:
    # TODO: use the right IP
    server: 10.4.1.2
    # TODO: use the right export path
    path: "/data/nfs/test"
  persistentVolumeReclaimPolicy: Retain #TODO: specify relcaim policy Recycle or Retain
---
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: test-pvc
  namespace: test
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 100Gi
  selector:
    matchLabels:
      alicloud-pvname: test-pvc
```

## base-deployment
```yaml
apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
kind: Deployment
metadata:
  name: _svc_var_
  namespace: _namespace_var_
spec:
  selector:
    matchLabels:
      app: _svc_var_
  replicas: 1 # tells deployment to run 2 pods matching the template
  template:
    metadata:
      labels:
        app: _svc_var_
    spec:
      containers:
        - name: _svc_var_
          image: _image_name_var_
          command:
            - "/data/_app_name_var_/bin/server"
          args:
            - "-config"
            - "/data/_app_name_var_/configs/config._run_env_var_.yaml"
          volumeMounts:
            - mountPath: "/data/my/logs"
              name: _app_name_var_-log
      volumes:
        - name: _app_name_var_-log
          persistentVolumeClaim:
            claimName: pvc-match-_run_env_var_
---
apiVersion: v1
kind: Service
metadata:
  name: _svc_var_
  namespace: _namespace_var_
spec:
  type: NodePort
  selector:
    app: _svc_var_
  ports:
    - protocol: TCP
      port: _port_var_
      nodePort: _nodePort_var_
      name: rpc
```