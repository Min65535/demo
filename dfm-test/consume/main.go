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
)

func init() {
	//insert into dfm_test.case(name,value,block) values ("w","ww","0");
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
	if err := conf.Table("case").Where(`id in (` + strId + `)`).Update(map[string]interface{}{"block": "0", "remark": "signed"}).Error; err != nil {
		log.QyLogger.Error("addTheRemark exists an error", zap.Error(err))
		return
	}
}

func getDataFromDb(conf *gorm.DB) (data []common.NameAndValue, err error) {
	if err = conf.Table("case").Where("block = ? and remark = ?", "0", "").Order("id asc", true).Limit(120).Find(&data).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.QyLogger.Error("getDataFromDb error", zap.Error(err))
	}
	return
}

//---------------------------whole tasks--------------------------//

func HugeTask(tasks chan func()) {
	dataChan := make(chan []common.NameAndValue)
	//Producer,生产数据
	go func() {
		var dbConf = common.GetDbConfig().GetDB()
		for {
			arr, err := getDataFromDb(dbConf)
			if err != nil {
				return
			}
			if len(arr) < 1 {
				time.Sleep(5 * time.Second)
				continue
			}
			closeTheBlock(dbConf, arr)
			dataChan <- arr
		}
	}()

	//将数据放入任务channel
	go func() {
		var dbConf = common.GetDbConfig().GetDB()
		for {
			select {
			case data := <-dataChan:
				tasks <- func() {
					addTheRemark(dbConf, data)
				}
			}
		}
	}()

	//Consumer,从channel取出任务并执行
	go func() {
		for {
			select {
			case m := <-tasks:
				m()
				log.QyLogger.Info("goroutine 1 finish the task")
			}
		}
	}()

	//Consumer,从channel取出任务并执行
	go func() {
		for {
			select {
			case m := <-tasks:
				m()
				log.QyLogger.Info("goroutine 2 finish the task")
			}
		}
	}()

	////Consumer,从channel取出任务并执行
	//go func() {
	//	for {
	//		select {
	//		case m := <-tasks:
	//			m()
	//			log.QyLogger.Info("goroutine 3 finish the task")
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
