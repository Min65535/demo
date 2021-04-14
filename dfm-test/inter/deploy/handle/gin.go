package handle

import "context"

type Publish interface {
	Login(ctx context.Context) error
	Auth(ctx context.Context) error
	ImagePush(ctx context.Context) error
	SvcDeploy(ctx context.Context) error
}
