package logger

import (
	"testing"
)

// func TestInitLogger(t *testing.T) {
// 	InitLogger("fffff")
// 	// 1669621402550827
// 	// 1669060799803
// 	// 1669702997779
// 	// 1669621429
// 	fmt.Println("nt1:", time.Now().Unix())
// 	fmt.Println("nt2:", time.Now().UnixMilli())
// 	fmt.Println("nt3:", time.Now().UnixMicro())
// 	fmt.Println("1669060799803:", 1669060799803)
// }

func TestFileIsExist(t *testing.T) {

	// FileIsExist("/mnt/d/code/go/src/demo/dfm-test/inter/filedata/data/new.go")
	// FileIsExist("/mnt/d/code/go/src/demo/dfm-test/inter/filedata/data1/new1.go")
	Init(WithDirRoot("/mnt/d/code/go/src/demo/dfm-test/inter/data_logger"), WithLogName("data"))

}
