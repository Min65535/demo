package main

import (
	"demo/dfm-test/pkg/common/db"
	"demo/dfm-test/pkg/model"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "Rule Engine Mysql DB Gen"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "action", Usage: "create, migrate, drop"},
	}
	app.Commands = []cli.Command{
		{Name: "create", Usage: "create db", Action: doCreate},
		{Name: "drop", Usage: "drop db", Action: doDrop},
		{Name: "migrate", Usage: "migrate db", Action: doMigrate},
	}
	app.Action = action
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func action(c *cli.Context) {
	_ = cli.ShowAppHelp(c)
}

const _logName = "db name"

func doCreate(c *cli.Context) {
	conf := db.GetDBConfig()
	utilDB := mysql.MakeDBUtil(conf)
	log.QyLogger.Info("run db: "+c.Command.Name, zap.String(_logName, conf.DbName))
	utilDB.CreateDB()
}

func doDrop(c *cli.Context) {
	conf := db.GetDBConfig()
	if strings.Contains(conf.DbName, "prod") {
		log.QyLogger.Warn("can't drop prod db", zap.String(_logName, conf.DbName))
		return
	}
	utilDB := mysql.MakeDBUtil(conf)
	log.QyLogger.Info("run db: "+c.Command.Name, zap.String(_logName, conf.DbName))
	utilDB.DropDB()
}

func doMigrate(c *cli.Context) {
	log.QyLogger.Info("run db: " + c.Command.Name)
	mysqlDB := mysql.MakeDB(db.GetDBConfig()).GetDB()
	mysqlDB.AutoMigrate(
		&model.User{},
	)
}
