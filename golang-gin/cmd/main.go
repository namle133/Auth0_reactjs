package main

import (
	"CMS/config"
	"CMS/log"
	"CMS/router"
)

func main() {
	log.InitLogger()
	config.InitEnv()
	config.InitDB()
	router.InitRouter()
}
