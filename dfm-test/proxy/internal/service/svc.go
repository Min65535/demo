package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

type ProxyServer interface {
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	// ------------------------------------------------------------------------------ Info(对外提供服务)
	// 对外Get
	ProxyGet(context.Context) error
	// 对外Post
	ProxyPost(context.Context) error
}
