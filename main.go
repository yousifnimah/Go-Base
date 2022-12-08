package main

import (
	"gateway_api/Helper"
	"gateway_api/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	initMode()
	Helper.InitLocalize()
	r := gin.Default()
	InitCORS(r)
	Routes.InitRoutes(r)
	//Cron.Init() //Must enabled on live
	serve(r)
}
