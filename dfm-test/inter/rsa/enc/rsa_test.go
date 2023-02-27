package enc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNewRsaEncrypt(t *testing.T) {
	t.Skip()
	var pri, pub string
	rst := NewRsaEncrypt([]byte(pri), []byte(pub))
	f, err := os.Open("/home/min/edc2.txt")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	defer f.Close()
	out, err := os.Create("/home/min/output3.txt")
	if err != nil {
		fmt.Println("Create#err:", err.Error())
		return
	}
	defer out.Close()
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		arr := strings.Split(scn.Text(), " ")
		// fmt.Println("arr:", arr)
		timStr := arr[0] + arr[1]
		reqStr := arr[2]
		trStr := reqStr[1 : len(reqStr)-2]
		resByt, err := rst.Decrypt(trStr)
		if err != nil {
			fmt.Println("res str err:", err.Error())
			return
		}
		// fmt.Println("resByt:",timStr+ " " +string(resByt))
		te := timStr + " " + string(resByt)
		n3, err := out.WriteString(te + "\n") // 写入文件(字节数组)
		if err != nil {
			fmt.Println("write err:", err.Error())
			return
		}
		fmt.Printf("写入 %d 个字节n", n3)
	}
	out.Sync()
}
