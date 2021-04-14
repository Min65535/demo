package decode

import (
	"github.com/dipperin/go-ms-toolkit/log"
	"go.uber.org/zap"
)

func NewAdmJwt() (admJ JwtToken, cf func(), err error) {
	var (
		cfg struct {
			AdmJwtSecret string `json:"adm_jwt_secret" toml:"adm_jwt_secret"`
		}
	)

	admJ, err = NewJwtToken(cfg.AdmJwtSecret)
	if err != nil {
		log.QyLogger.Error("adm jwt token load fail", zap.Error(err))
		return
	}

	return
}
