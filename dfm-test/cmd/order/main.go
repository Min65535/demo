package main

import (
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
	"github.com/gin-gonic/gin"
	"github.com/min65535/demo/dfm-test/inter/order/dao"
	"github.com/min65535/demo/dfm-test/inter/order/handle"
	"github.com/min65535/demo/dfm-test/inter/order/router"
	"github.com/min65535/demo/dfm-test/inter/order/service"
	"github.com/min65535/demo/dfm-test/pkg/common/db"
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
