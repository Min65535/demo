package common

import (
	"github.com/dipperin/go-ms-toolkit/db-config"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
)

// ---------------------------DB--------------------------//
var DbConfig *db_config.DbConfig

func init() {
	DbConfig = db_config.NewDbConfig()
	DbConfig.DbName = "dfm_test"
	DbConfig.MaxIdleConns = 100
	DbConfig.MaxOpenConns = 100
}

func GetDbConfig() mysql.DB {
	// DbConfig = db_config.NewDbConfig()
	// DbConfig.DbName = "dfm_test"
	return mysql.MakeDB(DbConfig)
}

// ---------------------------base data--------------------------//
type NameAndValue struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Block  string `json:"block"`
	Remark string `json:"remark"`
	Status uint   `json:"status" gorm:"column:status;default:0;index"`
}

func (n NameAndValue) TableName() string {
	return "case"
}
