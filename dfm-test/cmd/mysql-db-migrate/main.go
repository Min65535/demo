package main

import (
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
	"github.com/min65535/demo/dfm-test/inter/consume/common"
	"github.com/min65535/demo/dfm-test/pkg/common/db"
	"github.com/min65535/demo/dfm-test/pkg/model"
	"github.com/urfave/cli"
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
	utilDB.CreateDB()
}

func doDrop(c *cli.Context) {
	conf := db.GetDBConfig()
	if strings.Contains(conf.DbName, "prod") {
		return
	}
	utilDB := mysql.MakeDBUtil(conf)
	utilDB.DropDB()
}

func doMigrate(c *cli.Context) {
	mysqlDB := mysql.MakeDB(db.GetDBConfig()).GetDB()
	mysqlDB.AutoMigrate(
		&common.NameAndValue{}, &model.DemoOrder{}, &model.User{},
	)
}
