package main

import (
	"demo/dfm-test/inter/order/dao"
	"demo/dfm-test/inter/order/handle"
	"demo/dfm-test/inter/order/router"
	"demo/dfm-test/inter/order/service"
	"demo/dfm-test/pkg/common/db"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	orm := mysql.MakeDB(db.GetDBConfig())
	orderDao := dao.NewOrderDao(orm.GetDB())
	svcOrd := service.NewOrderService(orderDao)
	hld := handle.NewHandler(svcOrd)
	server := router.NewServer(hld, gin.Default())
	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
