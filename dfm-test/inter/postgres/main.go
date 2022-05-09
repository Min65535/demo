package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var dbConfig DatabaseConfig

func init() {
	// 在init方法中配置
	dbConfig = DatabaseConfig{
		User:     "",
		Password: "",
		Addr:     "127.0.0.1",
		Database: "lotus",
		PoolSize: 20,
		Slow:     50,
		Port:     5432,
	}
}

type DatabaseConfig struct {
	User     string
	Password string
	Addr     string
	Database string
	PoolSize int
	Slow     int
	Port     int
}

type Database struct {
	pg    *pg.DB
	pgorm *gorm.DB
}

func (db *Database) GetPg() *pg.DB {
	return db.pg
}

// func (db *Database) GetGorm() *gorm.DB {
// 	return db.pgorm
// }

func NewDataBase() *Database {
	return &Database{
		pg:    newPgDB(),
		pgorm: newGormDB(),
	}
}

func newPgDB() *pg.DB {
	return connectPg(
		&pg.Options{
			User:         dbConfig.User,
			Password:     dbConfig.Password,
			Database:     dbConfig.Database,
			Addr:         fmt.Sprintf("%s:%d", dbConfig.Addr, dbConfig.Port),
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 5,
			IdleTimeout:  time.Second * 120,
			PoolSize:     dbConfig.PoolSize,
		}, dbConfig.Slow)
}

func newGormDB() *gorm.DB {
	instance, err := gorm.Open("postgres", connStrGet())
	if err != nil {
		return nil
	}
	instance.DB().SetConnMaxLifetime(time.Minute * 5)
	instance.DB().SetMaxIdleConns(10)
	instance.DB().SetMaxOpenConns(dbConfig.PoolSize)
	instance.LogMode(true)
	return instance
}

func connectPg(opt *pg.Options, slow int) *pg.DB {
	db := pg.Connect(opt)
	var n string
	res, err := db.QueryOne(pg.Scan(&n), "select now() ")
	if err != nil {
		panic(err)
	}
	fmt.Println("res:", res)
	fmt.Printf("connect pg %s %s success on %s\n", db.String(), opt.Database, n)
	return db
}

func main() {
	db := NewDataBase()
	// 测试查库
	var n string
	_, err := db.pg.QueryOne(pg.Scan(&n), "select now() ")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func connStrGet() string {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d",
		dbConfig.Addr, dbConfig.User, dbConfig.Database, dbConfig.Password, dbConfig.Port)
	return connStr
}
