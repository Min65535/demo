package email

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	"github.com/alibabacloud-go/tea/tea"
)

type NotifyEmail interface {
	SendEmailLatest(req *dm20151123.SingleSendMailRequest) error
}

type Client struct {
	cli *dm20151123.Client
}

func (c *Client) SendEmailLatest(req *dm20151123.SingleSendMailRequest) error {
	req.AddressType = tea.Int32(1)
	req.ReplyToAddress = tea.Bool(false)
	// 复制代码运行请自行打印 API 的返回值
	_, err := c.cli.SingleSendMail(req)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (*dm20151123.Client, error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dm.aliyuncs.com")
	return dm20151123.NewClient(config)
}

func NewClient(accessKeyId *string, accessKeySecret *string) NotifyEmail {
	cli, err := CreateClient(accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	return &Client{cli: cli}
}
