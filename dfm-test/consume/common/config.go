package common

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/db-config"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
)

//---------------------------DB--------------------------//
var DbConfig *db_config.DbConfig

func init() {
	DbConfig = db_config.NewDbConfig()
	DbConfig.DbName = "dfm_test"
	DbConfig.MaxIdleConns = 100
	DbConfig.MaxOpenConns = 100
}

func GetDbConfig() (mysql.DB) {
	//DbConfig = db_config.NewDbConfig()
	//DbConfig.DbName = "dfm_test"
	return mysql.MakeDB(DbConfig)
}

//---------------------------base data--------------------------//
type NameAndValue struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Block  string `json:"block"`
	Remark string `json:"remark"`
}

type AllData struct {
	Data []NameAndValue
}

func (a *AllData) GetData() ([]NameAndValue) {
	fmt.Println("len of a.Data in GetData:", len(a.Data))
	if len(a.Data) <= 150 {
		defer func(a *AllData) {
			a.Data = nil
		}(a)
		return a.Data
	}
	defer func(a *AllData) {
		a.Data = a.Data[150:]
	}(a)
	return a.Data[:150]
}

func (a *AllData) SetData(data []NameAndValue) {
	a.Data = append(a.Data, data...)
	fmt.Println("the len of a.Data in SetData:", len(a.Data))
}
