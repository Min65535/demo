package dao

import (
	"github.com/min65535/demo/dfm-test/inter/publisher/biz"
	"github.com/min65535/demo/dfm-test/pkg/model"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
)


type Db struct {
	db mysql.DB
}

func NewDb(db mysql.DB) biz.UserDao {
	return &Db{db: db}
}

func (d *Db) GetUserByAcc(account string) (*model.User, error) {
	panic("implement me")
}

func (d *Db) GetUserById(id uint) (*model.User, error) {
	panic("implement me")
}
