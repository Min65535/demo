package main

import (
	"demo/dfm-test/inter/rsa/enc"
	"fmt"
)

func main() {
	pub, pri, err := enc.GenRsaKey(2048)
	if err != nil {
		fmt.Println("GenRsaKey err", err)
		return
	}

	msg := `this is unsafe!!!`
	r := enc.NewRsaEncrypt([]byte(pri), []byte(pub))
	eMsg, err := r.EncryptBlock(msg)
	if err != nil {
		fmt.Println("EncryptBlock err", err)
		return
	}
	fmt.Println("eMsg:", eMsg)

	msgBt, err := r.Decrypt(eMsg)
	if err != nil {
		fmt.Println("Decrypt err", err)
		return
	}
	fmt.Println("msgBt:", string(msgBt))

}
