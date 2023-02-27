package captcha

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Skip()
	type Resp struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data,omitempty"`
		ErrCode string      `json:"err_code,omitempty"`
		ErrMsg  string      `json:"err_msg,omitempty"`
	}
	clt := NewClient(&Cfg{
		// need param
		SecretId:     "",
		SecretKey:    "",
		CaptchaAppId: 0,
		AppSecretKey: "",
	})
	fmt.Println("clt:cfg:", json.StringifyJson(clt.cfg))
	const (
		XForwardedFor = "X-Forwarded-For"
		XRealIP       = "X-Real-IP"
	)
	eg := gin.Default()
	eg.Handle("POST", "/api/v1/captcha/test", func(c *gin.Context) {
		req := new(Req)
		if err := c.MustBindWith(req, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
			c.JSON(400, &Resp{Success: false, ErrCode: "000-000-A-000"})
			return
		}
		rqt := c.Request
		remoteAddr := rqt.RemoteAddr
		if ip := rqt.Header.Get(XRealIP); ip != "" {
			remoteAddr = ip
		} else if ip = rqt.Header.Get(XForwardedFor); ip != "" {
			remoteAddr = ip
		} else {
			remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
		}

		if remoteAddr == "::1" || remoteAddr == "" {
			remoteAddr = "127.0.0.1"
		}

		fmt.Println("req:", req)
		fmt.Println("remoteAddr:", remoteAddr)

		if err := req.Valid(); err != nil {
			fmt.Println("DescribeCaptchaResultRequestValid#err:", err.Error())
			c.JSON(400, &Resp{Success: false, ErrCode: "400", ErrMsg: err.Error()})
			return
		}
		res, err := clt.DescribeCaptchaResultLatest(req)
		if err != nil {
			fmt.Println("err:", err.Error())
			c.JSON(400, &Resp{Success: false, ErrCode: "400", ErrMsg: err.Error()})
			return
		}
		fmt.Println("res:", json.StringifyJson(res))
		c.JSON(200, &Resp{Success: true, Data: res})
	})
	eg.Run("0.0.0.0:8081")
}
