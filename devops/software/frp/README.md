## frp
```text
export FRP_VERSION=0.29.1
sudo mkdir -p /etc/frp
cd /etc/frp
sudo wget "https://github.com/fatedier/frp/releases/download/v${FRP_VERSION}/frp_${FRP_VERSION}_linux_amd64.tar.gz"
sudo tar xzvf frp_${FRP_VERSION}_linux_amd64.tar.gz
sudo mv frp_${FRP_VERSION}_linux_amd64/* /etc/frp
```

## server
```text
vi frps.ini 
./frps -c ./frps.ini
```

## client
```text
vi frpc.ini 
./frpc -c ./frpc.ini
```