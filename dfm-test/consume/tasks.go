package main

import (
	"demo/dfm-test/consume/common"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"strconv"
)

func init() {
	// insert into dfm_test.case(name,value,block) values ("w","ww","0");
	var str string
	str = `insert into dfm_test.case (name,value,block,remark) values `
	for i := 0; i < 5000; i++ {
		str = str + `("` + strconv.Itoa(i) + `","w` + strconv.Itoa(i) + `","0",""),`
	}
	str = str[:len(str)-1]
	if err := common.GetDbConfig().GetDB().Exec(str).Error; err != nil {
		log.QyLogger.Error("insert data to db exists an error", zap.Error(err))
		return
	}
}

func closeTheBlock(conf *gorm.DB, data []*common.NameAndValue) {
	strId := ``
	for i := range data {
		if i < len(data)-1 {
			strId = strId + strconv.Itoa(int(data[i].Id)) + `,`
		} else {
			strId = strId + strconv.Itoa(int(data[i].Id))
		}
	}
	if err := conf.Table("case").Where(`id in (` + strId + `)`).Update(map[string]interface{}{"block": "1"}).Error; err != nil {
		log.QyLogger.Error("closeTheBlock exists an error", zap.Error(err))
		return
	}
}

func addTheRemark(conf *gorm.DB, data []*common.NameAndValue) {
	// con := conf.Model("case")
	con := conf
	switch len(data) {
	case 0:
		return
	case 1:
		con = conf.Table("case").Where("id=?", data[0].Id)

	default:
		strId := ``
		for i := range data {
			if i < len(data)-1 {
				strId = strId + strconv.Itoa(int(data[i].Id)) + `,`
			} else {
				strId = strId + strconv.Itoa(int(data[i].Id))
			}
		}
		con = con.Table("case").Where(`id in (` + strId + `)`)
	}
	if err := con.Table("case").Update(map[string]interface{}{"block": "0", "remark": "signed"}).Error; err != nil {
		log.QyLogger.Error("addTheRemark exists an error", zap.Error(err))
		return
	}
}

func getDataFromDb(conf *gorm.DB) (data []*common.NameAndValue, err error) {
	if err = conf.Table("case").Where("block = ? and remark = ?", "0", "").Order("id asc", true).Limit(120).Find(&data).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.QyLogger.Error("getDataFromDb error", zap.Error(err))
	}
	return
}
