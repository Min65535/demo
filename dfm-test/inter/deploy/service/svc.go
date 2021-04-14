package service

import (
	"context"
	dao2 "demo/dfm-test/inter/deploy/dao"
	"demo/dfm-test/pkg/command"
)

func NewSvc(dao dao2.Dao) *Svc {
	return &Svc{dao: dao, cli: command.NewToolClient()}
}

type Svc struct {
	dao dao2.Dao
	cli command.ToolCommand
}

func (s Svc) Login(ctx context.Context) error {
	panic("implement me")
}

func (s Svc) Auth(ctx context.Context) error {
	panic("implement me")
}

func (s Svc) ImagePush(ctx context.Context) error {
	panic("implement me")
}

func (s Svc) SvcDeploy(ctx context.Context) error {
	panic("implement me")
}
