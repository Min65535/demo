package db

import (
	db_config "github.com/dipperin/go-ms-toolkit/db-config"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
	"github.com/dipperin/go-ms-toolkit/qyenv"
)

var MyDbConfig *db_config.DbConfig

func GetDBConfig() *db_config.DbConfig {
	if MyDbConfig != nil {
		return MyDbConfig
	}
	MyDbConfig = db_config.NewDbConfig()
	MyDbConfig.DbCharset = "utf8mb4"
	MyDbConfig.MaxIdleConns = 100
	MyDbConfig.MaxOpenConns = 100
	dbName := "dfm"

	runEnv := qyenv.GetDBEnv()
	switch runEnv {
	case "test", "preprod", "prod":
		dbName += "_" + runEnv
	default:
		dbName += "_dev"
	}
	MyDbConfig.DbName = dbName
	return MyDbConfig
}

func GetDbConfigs() (app mysql.DB) {
	return mysql.MakeDB(GetDBConfig())
}
