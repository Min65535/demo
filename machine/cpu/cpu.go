package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"runtime"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Press Test")
	fmt.Println("---------------------")
	i := 1
	var cores = runtime.NumCPU()
	for {
		for t := 0; t < cores; t++ {
			go func() {
				for f := 0; f < 2*i; f++ {
					i++
				}
			}()
		}
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		//convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, It's finished and the size of 'i' is ", i)
			return
		}
	}
}
