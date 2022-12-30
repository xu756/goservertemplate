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
		log.Print("连接错误")
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

func SqliteCreateTable(Table ...interface{}) error {
	return GetSqlitedb().AutoMigrate(Table...)
}
