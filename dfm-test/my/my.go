package main

import (
	"os/exec"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"errors"
)

//获取可执行程序当前路径
func pathGet() (path string, err error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return
	}

	path, err = filepath.Abs(file)
	if err!=nil{
		return
	}
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}

	i := strings.LastIndex(path, "/")
	if i < 0 {
		err = errors.New(`没有找到字符“/”或者"\"`)
		return
	}
	path = string(path[0 : i+1])
	return
}

func main()  {

}
