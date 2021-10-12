package email

import (
	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	"github.com/alibabacloud-go/tea/tea"
	"testing"
)

func TestCreateClient(t *testing.T) {
	t.Skip()
	id, sec := "", ""
	cl := NewClient(&id, &sec)
	req := &dm20151123.SingleSendMailRequest{
		AccountName: tea.String("vc@mail.xhpl.com"),
		ToAddress:   tea.String("min65535zxcv@gmail.com"),
		Subject:     tea.String("ssr系列"),
		TextBody:    tea.String("尊敬的用户：您的验证码为：65535，该验证码5分钟内有效，如非本人操作请忽略！"),
	}
	if err := cl.SendEmailLatest(req); err != nil {
		t.Log("err:", err.Error())
	}
}
