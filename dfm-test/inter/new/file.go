package main

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/philchia/agollo"
	"github.com/tidwall/gjson"
	"math/rand"
	"strings"
	"time"
)

func main() {
	httpUploadDemo()
}

var cli = resty.New()

func httpUploadDemo() {
	e := gin.Default()
	{
		e.POST("/upload", func(c *gin.Context) {
			form, err := c.MultipartForm()
			if err != nil {
				fmt.Println("ParseForm:", err.Error())
				c.JSON(400, gin.H{"success": false, "err_msg": "参数错误"})
				return
			}
			// 参数校验
			key := "file"
			if _, ok := form.File[key]; !ok {
				c.JSON(400, gin.H{"success": false, "err_msg": "参数缺失"})
				return
			}
			fs := form.File[key]
			if len(fs) != 1 {
				c.JSON(400, gin.H{"success": false, "err_msg": "没有上传文件或数量过大"})
				return
			}

			// 获取文件数据
			fileOne := fs[0]
			fileReader, err := fileOne.Open()
			if err != nil {
				fmt.Println("OpenFile:", err.Error())
				c.JSON(400, gin.H{"success": false, "err_msg": "上传失败"})
				return
			}
			fix := strings.Split(fileOne.Filename, ".")
			lastFixName := strings.ToLower(fix[len(fix)-1])
			var suffix string
			switch { // https://github.com/gabriel-vasile/mimetype/blob/master/tree.go
			case strings.Contains("jpeg,png,jpg", lastFixName):
				suffix = lastFixName
			default:
				log.QyLogger.Info("lastFixName:" + lastFixName)
				c.JSON(400, gin.H{"success": false, "err_msg": "文件格式非法"})
				return
			}
			resp, err := cli.R().
				SetFileReader("file", getFileName(suffix), fileReader).
				SetFormData(map[string]string{
					"output": "json",
					"scene":  "",
					"path":   "",
				}).Post("http://172.30.1.1:8080/upload")
			if err != nil {
				fmt.Println("OpenFile:", err.Error())
				c.JSON(400, gin.H{"success": false, "err_msg": "上传失败"})
				return
			}
			if resp.StatusCode() != 200 {
				fmt.Println("requests fail, body:", resp.String(), ", statusCode:", resp.StatusCode())
				c.JSON(400, gin.H{"success": false, "err_msg": "上传失败"})
				return
			}
			fmt.Println("respBody:", resp.String())

			result := gjson.ParseBytes(resp.Body())
			if result.Get("url").String() == "" {
				fmt.Println("requests fail, body:", resp.String(), ", statusCode:", resp.StatusCode())
				c.JSON(400, gin.H{"success": false, "err_msg": "上传失败"})
				return
			}
			c.JSON(200, gin.H{"success": true, "data": gin.H{"file_url": result.Get("url").String()}})
		})
	}
	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}

func genRandNum() int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63()
}

func getFileName(suffixName string) string {
	return fmt.Sprintf("%d%s%s", genRandNum(), time.Now().Format("20060102150405"), suffixName)
}

func apolloDemo() {
	if err := agollo.StartWithConf(&agollo.Conf{
		AppID:          "demo",
		Cluster:        "default",
		NameSpaceNames: []string{"application"},
		IP:             "172.30.9.73:8180",
	}); err != nil {
		panic(err)
	}

	fmt.Println(agollo.GetStringValue("msg", ""))
}
