package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Redis 配置
type redisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// Mysql 配置
type mysqlConfig struct {
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}

// Mongodb 配置
type mongodb struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type Data struct {
	// Redis 配置
	Redis redisConfig `yaml:"redis"`
	// Mysql 配置
	Mysql mysqlConfig `yaml:"mysql"`
	// Mongodb 配置
	//Mongo mongodb `yaml:"mongo"`
}

var InitData Data

func Config() {
	CreateDir("./media/upload")
	file, err := os.Open("./config.yaml")
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
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("配置文件解码失败", err.Error())
	} else {
		fmt.Println("配置文件解码成功")
		InitData = data
	}

}
