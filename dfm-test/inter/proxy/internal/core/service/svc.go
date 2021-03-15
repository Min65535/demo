package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

type ProxyServer interface {
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	// ------------------------------------------------------------------------------ Info(对外提供服务)
	// 对外
	Proxy(context.Context) error
}
