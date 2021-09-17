package biz

import (
	"demo/dfm-test/pkg/decode"
	"demo/dfm-test/pkg/model"
)

type UserDao interface {
	GetUserByAcc(account string) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
}

type UserUseCase struct {
	repo UserDao
	jwt  decode.JwtToken
}

func NewUserUseCase(repo UserDao, jwt decode.JwtToken) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		jwt:  jwt,
	}
}

func (uuc *UserUseCase) Auth() error {

	return nil
}

func (uuc *UserUseCase) Login() error {

	return nil
}
