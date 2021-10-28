package captcha

import (
	"errors"
	tcCli "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha/v20190722"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

type TcCaptcha interface {
	DescribeCaptchaResultLatest(req *Req) (bool, error)
}

type Cfg struct {
	SecretId     string `json:"secret_id" toml:"secret_id"`
	SecretKey    string `json:"secret_key" toml:"secret_key"`
	CaptchaAppId uint64 `json:"captcha_app_id" toml:"captcha_app_id"`
	AppSecretKey string `json:"app_secret_key" toml:"app_secret_key"`
}

type ClientCaptcha struct {
	cli *tcCli.Client
	cfg *Cfg
}

type Req struct {
	// 前端回调函数返回的用户验证票据
	Ticket string `json:"ticket,omitempty" name:"Ticket"`

	// 业务侧获取到的验证码使用者的外网IP
	// UserIp string `json:"user_ip" name:"UserIp"`

	// 前端回调函数返回的随机字符串
	Randstr string `json:"randstr,omitempty" name:"Randstr"`
}

func (r *Req) Valid() error {
	if r.Ticket == "" {
		return errors.New("参数缺失或非法")
	}
	if r.Randstr == "" {
		return errors.New("参数缺失或非法")
	}
	return nil
}

func (tc *ClientCaptcha) DescribeCaptchaResultLatest(req *Req) (bool, error) {
	request := tcCli.NewDescribeCaptchaResultRequest()
	request.CaptchaType = common.Uint64Ptr(9)
	request.Ticket = common.StringPtr(req.Ticket)
	request.UserIp = common.StringPtr("127.0.0.1")
	request.Randstr = common.StringPtr(req.Randstr)
	request.CaptchaAppId = common.Uint64Ptr(tc.cfg.CaptchaAppId)
	request.AppSecretKey = common.StringPtr(tc.cfg.AppSecretKey)
	res, err := tc.cli.DescribeCaptchaResult(request)
	if err != nil {
		return false, err
	}
	if res.Response == nil || res.Response.CaptchaCode == nil || res.Response.CaptchaMsg == nil {
		return false, errors.New("响应非法，请稍后再试")
	}
	code := *res.Response.CaptchaCode
	msg := *res.Response.CaptchaMsg
	if code == 1 && msg == "OK" {
		return true, nil
	}
	return false, errors.New(msg)
}

func NewClient(cfg *Cfg) *ClientCaptcha {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "captcha.ap-hongkong.tencentcloudapi.com"
	// cpf.HttpProfile.Endpoint = "captcha.tencentcloudapi.com"
	cli, err := tcCli.NewClient(common.NewCredential(cfg.SecretId, cfg.SecretKey), regions.HongKong, cpf)
	if err != nil {
		panic("getCaptchaSecret#NewClient#err:" + err.Error())
	}
	return &ClientCaptcha{
		cli: cli,
		cfg: cfg,
	}
}
