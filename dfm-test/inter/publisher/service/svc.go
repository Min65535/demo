package service

import (
	"context"
	"github.com/min65535/demo/dfm-test/inter/publisher/api"
	"github.com/min65535/demo/dfm-test/inter/publisher/biz"
)

func NewSvc(uuc *biz.UserUseCase, tuc *biz.ToolUseCase) *Svc {
	return &Svc{uuc: uuc, tuc: tuc}
}

type Svc struct {
	uuc *biz.UserUseCase
	tuc *biz.ToolUseCase
}

func (s Svc) Auth(ctx context.Context) error {
	panic("implement me")
}

func (s Svc) Login(ctx context.Context, req *api.LoginReq) (*api.LoginResp, error) {
	panic("implement me")
}

func (s Svc) ImagePush(ctx context.Context, req *api.ImagePushReq) error {
	panic("implement me")
}

func (s Svc) SvcDeploy(ctx context.Context, req *api.SvcDeployReq) error {
	panic("implement me")
}
