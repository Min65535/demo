package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// RSA公钥私钥产生
func GenRsaKey(bits int) (publicKeyStr, privateKeyStr string, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	bufferPrivate := new(bytes.Buffer)
	err = pem.Encode(bufferPrivate, block)
	if err != nil {
		return
	}
	privateKeyStr = bufferPrivate.String()
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	bufferPublic := new(bytes.Buffer)
	err = pem.Encode(bufferPublic, block)
	if err != nil {
		return
	}
	publicKeyStr = bufferPublic.String()
	fmt.Println("-------------公钥----------------")
	fmt.Println("\r", publicKeyStr)
	fmt.Println("--------------私钥---------------")
	fmt.Println("\r", privateKeyStr)
	return

}

type RsaEncrypt struct {
	priKey *rsa.PrivateKey
	pubKey *rsa.PublicKey
}

func NewRsaEncrypt(privateKey, publicKey []byte) *RsaEncrypt {
	block, _ := pem.Decode(privateKey)
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("NewRsaEncrypt#parse private key error. err_msg: %s", err.Error()))
	}

	pubBlock, _ := pem.Decode(publicKey)
	pubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		panic(fmt.Sprintf("parse public key error. err_msg: %s", err.Error()))
	}
	return &RsaEncrypt{
		priKey: priKey,
		pubKey: pubKey.(*rsa.PublicKey),
	}
}

func (r *RsaEncrypt) Decrypt(msg string) ([]byte, error) {
	cryptText, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return []byte{}, err
	}
	srcSize := len(cryptText)
	keySize := r.priKey.Size()
	offset := 0
	buffer := bytes.Buffer{}
	for offset < srcSize {
		endIndex := offset + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}

		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, r.priKey, cryptText[offset:endIndex])
		if err != nil {
			return []byte{}, err
		}
		buffer.Write(bytesOnce)
		offset = endIndex
	}
	return buffer.Bytes(), nil
}

func (r *RsaEncrypt) EncryptBlock(msg string) (strEncrypt string, err error) {
	src := []byte(msg)
	keySize, srcSize := r.pubKey.Size(), len(src)
	// 单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, r.pubKey, src[offSet:endIndex])
		if err != nil {
			return "", err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	strEncrypt = base64.StdEncoding.EncodeToString(buffer.Bytes())
	return
}

func main() {
	pub, pri, err := GenRsaKey(2048)
	if err != nil {
		fmt.Println("GenRsaKey err", err)
		return
	}

	msg := `this is unsafe!!!`
	r := NewRsaEncrypt([]byte(pri), []byte(pub))
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
