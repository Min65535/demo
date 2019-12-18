package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"runtime"
	"errors"
	"image/jpeg"
)

const DEFAULT_MAX_WIDTH float64 = 320
const DEFAULT_MAX_HEIGHT float64 = 180

// 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func makeThumbnail(imagePath, savePath string) (string, error) {

	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("os.Open err:", err.Error())
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("image.Decode err:", err.Error())
		return "", err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// 需要保存的文件
	imgFile, _ := os.Create(savePath)
	defer imgFile.Close()

	//// 以png格式保存文件
	//err = png.Encode(imgFile, m)
	// 以jpeg格式保存文件
	err = jpeg.Encode(imgFile, m, &jpeg.Options{Quality: 100})
	if err != nil {
		fmt.Println("png.Encode err:", err.Error())
		return "", err
	}

	return "", nil
}

//获取可执行程序当前路径
func pathGet() (path string, err error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return
	}

	path, err = filepath.Abs(file)
	if err != nil {
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

func main() {
	enter, _ := pathGet()
	str, err := makeThumbnail(enter+"haha.jpg", "ha1.jpg")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("str:", str)
}
