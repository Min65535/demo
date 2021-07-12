package service

import (
	"demo/dfm-test/inter/order/dao"
	"demo/dfm-test/pkg/model"
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

type OrderSvc interface {
	Create(req *model.DemoOrder) error
	UpdateByNo(orderNo string, m map[string]interface{}) error
	QueryByNo(no string) (order model.DemoOrder, err error)
	QueryList(page, size int) (orders []*model.DemoOrder, err error)
	DeleteByNo(no string) (order model.DemoOrder, err error)
	OrderDownload() (string, error)
}

type OrderService struct {
	orderDao dao.OrderDao
}

func NewOrderService(orderDao dao.OrderDao) OrderSvc {
	return &OrderService{orderDao: orderDao}
}

func (o *OrderService) Create(req *model.DemoOrder) error {
	return o.orderDao.Create(req)
}

func (o *OrderService) UpdateByNo(orderNo string, m map[string]interface{}) error {
	return o.orderDao.UpdateByNo(orderNo, m)
}

func (o *OrderService) QueryByNo(no string) (order model.DemoOrder, err error) {
	return o.orderDao.QueryByNo(no)
}

func (o *OrderService) QueryList(page, size int) (orders []*model.DemoOrder, err error) {
	return o.orderDao.QueryList(page, size)
}

func (o *OrderService) DeleteByNo(no string) (order model.DemoOrder, err error) {
	return o.orderDao.DeleteByNo(no)
}

func (o *OrderService) OrderDownload() (string, error) {
	orders, _ := o.QueryList(0, 100)
	filename := "order.xlsx"
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return "", err
	}
	for _, v := range orders {
		fmt.Println(v)
		row := sheet.AddRow()
		id := row.AddCell()
		id.Value = strconv.Itoa(int(v.ID))
		CreatedAt := row.AddCell()
		CreatedAt.Value = v.CreatedAt.Format("2006-01-02 15:04:05")
		UpdatedAt := row.AddCell()
		UpdatedAt.Value = v.UpdatedAt.Format("2006-01-02 15:04:05")
		// DeletedAt := row.AddCell()
		// DeletedAt.Value = v.DeletedAt.Format("2006-01-02 15:04:05")
		orderNo := row.AddCell()
		orderNo.Value = v.OrderNo
		name := row.AddCell()
		name.Value = v.UserName
		Amount := row.AddCell()
		Amount.Value = strconv.Itoa(int(v.Amount))
		fileUrl := row.AddCell()
		fileUrl.Value = v.FileUrl
		Status := row.AddCell()
		Status.Value = v.Status
	}
	err = file.Save(filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}
