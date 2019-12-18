package common

import (
	"testing"
	"strconv"
	"github.com/dipperin/go-ms-toolkit/log"
	"go.uber.org/zap"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func TestGetDbConfig(t *testing.T) {
	var str string
	str = `insert into dfm_test.case (name,value,block,remark) values `
	for i := 0; i < 100000; i++ {
		str = str + `("` + strconv.Itoa(i) + `","i` + strconv.Itoa(i) + `","0",""),`
	}
	str = str[:len(str)-1]
	if err := GetDbConfig().GetDB().Exec(str).Error; err != nil {
		log.QyLogger.Error("insert data to db exists an error", zap.Error(err))
		return
	}
}

func TestNew(t *testing.T) {
	var data NameAndValue
	conf := GetDbConfig().GetDB()
	conf.LogMode(false)
	sign := conf.Table("case").Where("name=?", "ww").First(&data).RecordNotFound()
	fmt.Println("sign:", sign)
}
