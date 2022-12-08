package main

import (
	"gateway_api/Helper"
	"github.com/gin-gonic/gin"
	"os"
)

func initMode() {
	Helper.LoadEnv()
	mode := os.Getenv("GIN_MODE")
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
