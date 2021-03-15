package token

import (
	"demo/dfm-test/inter/mydfs/pkg/fileserver/token/jwt"
	"errors"
	"github.com/dipperin/go-ms-toolkit/log"
	"go.uber.org/zap"
	"time"
)

type Srv struct {
	*jwt.Jwt
	// todo: get exp from apollo
	expAt time.Duration
}

func NewSrv() *Srv {
	return &Srv{Jwt: jwt.NewJwt(), expAt: 7 * 24 * time.Hour}
}

var AuthFailedErr = errors.New("auth token fail")

// 验证token
func (s *Srv) Auth(userID uint, token string) error {
	uid, err := s.DecodeToken(token)
	if err != nil {
		return err
	}
	if uid != userID {
		log.QyLogger.Warn("auth token failed", zap.Uint("user_id", userID), zap.Uint("token_uid", uid))
		return AuthFailedErr
	}
	return nil
}

// 生成token
func (s *Srv) Token(userID uint) (string, error) {
	return s.GenerateToken(userID, s.expAt)
}
