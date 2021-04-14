package dao

import (
	"demo/dfm-test/pkg/model"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
)

type Dao interface {
	GetUserByAcc(account string) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
}

type Db struct {
	db mysql.DB
}

func NewDb(db mysql.DB) Dao {
	return &Db{db: db}
}

func (d *Db) GetUserByAcc(account string) (*model.User, error) {
	panic("implement me")
}

func (d *Db) GetUserById(id uint) (*model.User, error) {
	panic("implement me")
}
