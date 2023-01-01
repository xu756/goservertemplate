package main

import (
	"server/common/config"
	"server/common/model"
	"server/common/router"
)

func main() {
	config.Config()
	model.InitsqLite()
	model.InitMysqlDB()
	router.InitRouter()
}
