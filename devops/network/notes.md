## ubuntu iptables
iptables -L -n --line-number

iptables -D INPUT 2

iptables -D OUTPUT

iptables -D FORWARD

iptables -X chain

iptables -S

iptables -P FORWARD ACCEPT

## gateway
route add default gw 192.168.1.1