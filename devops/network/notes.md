## ubuntu iptables
```text
iptables -L -n --line-number

iptables -D INPUT 2

iptables -D OUTPUT

iptables -D FORWARD

iptables -X chain

iptables -S

iptables -P FORWARD ACCEPT
```

## gateway
```text
route add default gw 192.168.1.1
```

## readme
```text
network:
  ethernets:
    enp3s0:
      addresses: [172.16.1.62/24]
      gateway4: 172.16.1.1
      dhcp4: no
      nameservers:
        addresses: [114.114.114.114, 8.8.8.8]
  version: 2
```