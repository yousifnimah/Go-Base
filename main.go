package main

import (
	"base/Helper"
	"base/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	initMode()
	Helper.InitLocalize()
	r := gin.Default()
	InitCORS(r)
	Routes.InitRoutes(r)
	serve(r)
}
