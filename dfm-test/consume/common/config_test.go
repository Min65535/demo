package common

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/log"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"strconv"
	"testing"
)

func TestGetDbConfig(t *testing.T) {
	var str string
	str = `insert into dfm_test.case (name,value,block,remark) values `
	for i := 0; i < 10000; i++ {
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
	conf.LogMode(true)
	sign := conf.Table("case").Where("name=?", "ww").First(&data).RecordNotFound()
	fmt.Println("sign:", sign)
	//db.getDB().Model(&model.Verification{}).Select("distinct block_number from verifications v where v.address=?").Count(&num)
	var num int
	var c []NameAndValue
	conf.Raw("select distinct name from `case` v where v.block=?", 0).Scan(&c)
	num = len(c)
	fmt.Println("num:", num)

}
