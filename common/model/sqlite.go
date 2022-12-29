package model

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var SqliteDB *gorm.DB

func InitsqLite() {
	db, err := gorm.Open(sqlite.Open("logs.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("sqlite数据库连接失败")
	}
	fmt.Println("sqlite数据库连接成功")
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	SqliteDB = db
}

func GetSqlitedb() *gorm.DB {
	sqlDB, err := SqliteDB.DB()
	if err != nil {
		log.Print("connect db server failed.")
		InitsqLite()
	}
	if err := sqlDB.Ping(); err != nil {
		err := sqlDB.Close()
		if err != nil {
			return nil
		}
		InitsqLite()
	}
	return SqliteDB
}

type NewTable struct {
	Table interface{}
	Name  string
}

func SqliteCreateTable(Name string, Table interface{}) {
	err := GetSqlitedb().AutoMigrate(Table)
	if err != nil {
		fmt.Printf("sqlite创建 %s 表失败!!!", Name)
	}
}
