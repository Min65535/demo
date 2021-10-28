package main

import (
	"errors"
	"fmt"
	"github.com/thinkeridea/go-extend/exnet"
	"net"
	"reflect"
	"strconv"
	"strings"
)

type IntIP struct {
	IP    string
	Intip int
}

func main() {
	//var x = IntIP{IP: "192.168.1.1"}
	//fmt.Println(x)
	//x.ToIntIp()
	//fmt.Println(x)

	ip := "192.168.1.1"

	n, _ := exnet.IPString2Long(ip)
	s, _ := exnet.Long2IPString(n)

	fmt.Println(n, s == ip)

	Ip1 := net.ParseIP(ip) // 会得到一个16字节的byte，主要为了兼容ipv6
	n, _ = exnet.IP2Long(Ip1)

	Ip2, _ := exnet.Long2IP(n)

	fmt.Println(n, reflect.DeepEqual(Ip1[12:], Ip2))
}

func (s *IntIP) String() string {
	return s.IP
}

func (s *IntIP) ToIntIp() error {
	inTip, err := ConvertToIntIP(s.IP)
	if err != nil {
		return err
	}
	s.Intip = inTip
	return nil
}

func (s *IntIP) ToString() (string, error) {
	i4 := s.Intip & 255
	i3 := s.Intip >> 8 & 255
	i2 := s.Intip >> 16 & 255
	i1 := s.Intip >> 24 & 255
	if i1 > 255 || i2 > 255 || i3 > 255 || i4 > 255 {
		return "", errors.New("Isn't a IntIP Type")
	}
	ipstring := fmt.Sprintf("%d.%d.%d.%d", i4, i3, i2, i1)
	s.IP = ipstring
	return ipstring, nil
}
func ConvertToIntIP(ip string) (int, error) {
	ips := strings.Split(ip, ".")
	E := errors.New("Not A IP")
	if len(ips) != 4 {
		return 0, E
	}
	var intIP int
	for k, v := range ips {
		i, err := strconv.Atoi(v)
		if err != nil || i > 255 {
			return 0, E
		}
		intIP = intIP | i<<uint(8*(3-k))
		fmt.Println(fmt.Sprintf("第%s个的值是%s", strconv.Itoa(k+1), strconv.Itoa(i<<uint(8*(3-k)))))
		fmt.Println(fmt.Sprintf("第%s个的值是%s", strconv.Itoa(k+1), strconv.Itoa(intIP)))
	}
	return intIP, nil
}
