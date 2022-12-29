package main

import (
	"server/common/config"
	"server/router"
)

func main() {
	config.Config()
	router.InitRouter()
}
