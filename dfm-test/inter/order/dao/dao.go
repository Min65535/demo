package dao

import (
	"github.com/min65535/demo/dfm-test/pkg/model"
	"github.com/jinzhu/gorm"

)

//go:generate mockgen -source=dao.go -package=dao -destination=dao_mock.go OrderDao
type OrderDao interface {
	Create(req *model.DemoOrder) error
	UpdateByNo(orderNo string, m map[string]interface{}) error
	QueryByNo(no string) (order model.DemoOrder, err error)
	QueryList(page, size int) (orders []*model.DemoOrder, err error)
	DeleteByNo(no string) (order model.DemoOrder, err error)
}

type orderDao struct {
	db *gorm.DB
}

func NewOrderDao(db *gorm.DB) *orderDao {
	return &orderDao{db: db}
}

func (o *orderDao) Create(req *model.DemoOrder) error {
	return o.db.Create(req).Error
}

func (o *orderDao) UpdateByNo(orderNo string, m map[string]interface{}) error {
	return o.db.Model(&model.DemoOrder{}).Where("order_no=?", orderNo).Updates(m).Error
}

func (o *orderDao) QueryByNo(no string) (order model.DemoOrder, err error) {
	// err = o.db.Where("order_no=?", no).Find(&order).Error
	err = o.db.First(&order, "order_no = ?", no).Error
	return
}

func (o *orderDao) QueryList(page, size int) (orders []*model.DemoOrder, err error) {
	err = o.db.Offset(page * size).Limit(size).Find(&orders).Error
	return
}

func (o *orderDao) DeleteByNo(no string) (order model.DemoOrder, err error) {
	err = o.db.Delete(&order, "order_no=?", no).Error
	// err = o.db.Where("order_no=?", no).Delete(&order).Error
	return
}
