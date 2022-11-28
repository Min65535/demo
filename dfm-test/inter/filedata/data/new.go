package logger

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"os"
)

func InitLogger(name string) {
	switch qyenv.GetDockerEnv() {
	case "", "0":
	default:
		// logDir := "/data/match/logs/toptop_match_data/"
		logDir := "/mnt/d/code/go/src/demo/dfm-test/inter/filedata/data/"
		if err := os.MkdirAll(logDir, 0755); err != nil {
			panic(err)
		}
		byS := "SS"
		if _, err := os.Create(logDir + name + ".log"); err != nil {
			fmt.Println("ccccc: ", err)
		}
		if err := os.WriteFile(logDir+name+".log", []byte(byS), 755); err != nil {
			fmt.Println("wwww: ", err)
		}
	}
}
