package main

import (
	"sync"
	"github.com/jinzhu/gorm"
	"github.com/dipperin/go-ms-toolkit/log"
	"go.uber.org/zap"
	"strconv"
	"demo/dfm-test/consume/common"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/dipperin/go-ms-toolkit/json"
)

func init() {
	//insert into dfm_test.case(name,value,block) values ("w","ww","0");
	var str string
	str = `insert into dfm_test.case (name,value,block) values `
	for i := 0; i < 5000; i++ {
		str = str + `("` + strconv.Itoa(i) + `","w` + strconv.Itoa(i) + `","0"),`
	}
	str = str[:len(str)-1]
	if err := common.GetDbConfig().GetDB().Exec(str).Error; err != nil {
		log.QyLogger.Error("insert data to db exists an error", zap.Error(err))
		return
	}
}

var TheData common.AllData
//var dbConf = common.GetDbConfig().GetDB()

//读取数据//Exec(`select * from dfm_test.case where block = "0" order by id desc limit 150`)
func readDataFromDb(conf *gorm.DB) {
	var data []common.NameAndValue
	if err := conf.Table("case").Where("block = ?", "0").Order("id desc", true).Limit(150).Find(&data).Error; err != nil && err != gorm.ErrRecordNotFound {
		//if err := common.GetDbConfig().GetDB().Exec(`select * from dfm_test.case where block = "0" order by id desc limit 150`).Scan().Error; err != nil && err != gorm.ErrRecordNotFound {
		log.QyLogger.Error("readDataFromDb error", zap.Error(err))
		return
	}
	if len(data) > 0 {
		closeTheBlock(conf, data)
	}
	TheData.SetData(data)
}

func closeTheBlock(conf *gorm.DB, data []common.NameAndValue) {
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

func addTheRemark(conf *gorm.DB, data []common.NameAndValue) {
	strId := ``
	for i := range data {
		if i < len(data)-1 {
			strId = strId + strconv.Itoa(int(data[i].Id)) + `,`
		} else {
			strId = strId + strconv.Itoa(int(data[i].Id))
		}
	}
	if err := conf.Table("case").Where(`id in (` + strId + `)`).Update(map[string]interface{}{"remark": "signed"}).Error; err != nil {
		log.QyLogger.Error("addTheRemark exists an error", zap.Error(err))
		return
	}
}

//---------------------------whole tasks--------------------------//

func HugeTask(tasks chan func()) {

	go func() {
		var dbConf = common.GetDbConfig().GetDB()
		for {
			readDataFromDb(dbConf)
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		var dbConf = common.GetDbConfig().GetDB()
		for {
			tasks <- func() {
				data := TheData.GetData()
				log.QyLogger.Info("GetData:", zap.String("data", json.StringifyJson(data)))
				if len(data) > 0 {
					addTheRemark(dbConf, data)
				}
			}
			time.Sleep(10 * time.Second)
		}

	}()

	go func() {
		for {
			select {
			case m := <-tasks:
				m()
			}
		}
	}()

	//go func() {
	//	for {
	//		select {
	//		case m := <-tasks:
	//			m()
	//		}
	//	}
	//}()

	select {}

}

var wg sync.WaitGroup

func main() {
	//var totalNum int
	//flag.IntVar(&totalNum, "total", 1, "总数量")
	//var mobileStartNum int
	//flag.IntVar(&mobileStartNum, "mobile", 1, "手机号开始数字")
	//var profileIdStartNum int
	//flag.IntVar(&profileIdStartNum, "profile", 1, "profileId开始数字")
	//flag.Parse()
	//fmt.Println("totalNum:", totalNum)
	//fmt.Println("mobileStartNum:", mobileStartNum)
	//fmt.Println("profileIdStartNum:", profileIdStartNum)

	tasks := make(chan func(), 10)
	HugeTask(tasks)
	close(tasks)
	wg.Wait()
}
