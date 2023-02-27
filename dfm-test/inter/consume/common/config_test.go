package common

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/dipperin/go-ms-toolkit/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/min65535/demo/dfm-test/pkg/common/db"
	"go.uber.org/zap"
	"os"
	"strconv"
	"testing"
)

func TestGetDbConfig(t *testing.T) {
	os.Setenv("db_env", "test")
	// defer db.GetDbConfigs().ClearAllData()
	con := db.GetDbConfigs()
	var str string
	str = `insert into dfm_test.case (name,value,block,remark) values `
	for i := 0; i < 10000; i++ {
		str = str + `("` + strconv.Itoa(i) + `","i` + strconv.Itoa(i) + `","0",""),`
	}
	str = str[:len(str)-1]
	if err := con.GetDB().Exec(str).Error; err != nil {
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
	// db.getDB().Model(&model.Verification{}).Select("distinct block_number from verifications v where v.address=?").Count(&num)
	var num int
	var c []NameAndValue
	conf.Raw("select distinct name from `case` v where v.block=?", 0).Scan(&c)
	num = len(c)
	fmt.Println("num:", num)

}

func TestHaHa(t *testing.T) {
	conf := GetDbConfig().GetDB()
	conf.LogMode(true)

	// str := `insert into dfm_test.case (name,value,block,remark) values `
	// for i := 0; i < 100000; i++ {
	// 	str = str + `("` + fmt.Sprintf("%d", i) + `","i` + fmt.Sprintf("%d", i) + `","0",""),`
	// }
	// str = str[:len(str)-1]
	// if err := GetDbConfig().GetDB().Exec(str).Error; err != nil {
	// 	log.QyLogger.Error("insert data to db exists an error", zap.Error(err))
	// 	return
	// }

	type NS struct {
		Name string `json:"name"`
	}
	var arrNs []*NS
	var arrStr []string
	var bmgCIDs = []string{"2", "1111111", "3333333333", "444444444"}
	addr2 := "2"
	addr3 := "3"
	// unionStr := `select name from dfm_test.case where name='` + addr2 + `' union select name from dfm_test.case where value='` + "i" + addr3 + `'`
	unionStr := `(select name from dfm_test.case where case.name='` + addr2 + `' union select name from dfm_test.case where case.value='` + "i" + addr3 + `') as t1`
	// conf = conf.Select("cid").Where("cid in (?)", bmgCIDs).Where("messages.from = ? or messages.to = ?", addr, addr)
	// if err := conf.Raw(`select name from dfm_test.case where name in (?) and name in (?)`, bmgCIDs, gorm.Expr(unionStr)).Scan(&arrNs).Error; err != nil {
	// if err := conf.Raw(`select name from dfm_test.case where name in (?)`, gorm.Expr(unionStr)).Scan(&arrNs).Error; err != nil {
	if err := conf.Raw(`select t1.name as name from `+unionStr+` where t1.name in (?)`, bmgCIDs).Scan(&arrNs).Error; err != nil {
		// if err := conf.Raw(unionStr).Scan(&arrNs).Error; err != nil {
		log.QyLogger.Error("TestHaHa Raw exists an error", zap.Error(err))

	}

	fmt.Println("arrNs:", json.StringifyJson(arrNs))
	fmt.Println("arrStr:", json.StringifyJson(arrStr))

	var totalCount int
	var data []string
	conf.Table("case").Count(&totalCount)
	pageSize := 500
	totalPages := (totalCount + pageSize - 1) / pageSize

	for page := 1; page <= totalPages; page++ {
		var arr []*NS
		if err := conf.Table("case").
			Offset((page - 1) * pageSize).
			Limit(pageSize).Find(&arr).Error; err != nil {
			fmt.Println("err:", err.Error())
		}
		for i := range arr {
			data = append(data, arr[i].Name)
		}
	}
	fmt.Println("len(data):", len(data))
}
