package main

import (
	"HackInBox.online/controller"
	"HackInBox.online/db"
	"HackInBox.online/middleware"
	"HackInBox.online/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		operation := os.Args[1]
		if operation == "create_db" {
			utils.UtilsLogger.Info("create db...")
			db.CreateDb()
			utils.UtilsLogger.Info("Success")
		}
		return
	}
	gin.SetMode(gin.ReleaseMode)
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	r := gin.New()

	r.Use(middleware.Logger())
	controller.MakRouter(r)
	utils.UtilsLogger.Info("Server Run Success: 0.0.0.0:8080")
	r.Run(":8080")
}
