package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Redis 配置
type redisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// Mysql 配置
type mysqlConfig struct {
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

// Mongodb 配置
type mongodb struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
}
type Data struct {
	// Redis 配置
	Redis redisConfig `json:"redis"`
	// Mysql 配置
	Mysql mysqlConfig `json:"mysql"`
	// Mongodb 配置
	Mongo mongodb `json:"mongo"`
}

var InitData Data

func Config() {
	CreateDir("./media/upload")
	file, err := os.Open("./config.yml")
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Print("失败")
		}
	}(file)
	// 读取文件内容
	var data Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("配置文件解码失败", err.Error())
	} else {
		fmt.Println("配置文件解码成功")
		InitData = data
	}

}
