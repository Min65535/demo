package access

import (
	"github.com/min65535/demo/dfm-test/pkg/sam/publish/role"
)

// ---------------- api ----------------
type Api = uint

const (
	// 推送
	ApiPush Api = 1 << iota
	// 发布
	ApiDeploy

	AllApi
)

var Apis = map[string]Api{
	// push
	"/api/v1/ping": ApiPush,
	"/api/v1/push": ApiPush,

	// deploy
	"/api/v1/deploy": ApiDeploy,
}

// ---------------- role ----------------

var Access = map[role.Role]Api{
	role.Admin:   AllApi - 1,
	role.Operate: ApiPush,
}
