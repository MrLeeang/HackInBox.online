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

	// 注册logger
	r.Use(middleware.Logger())
	// 注册路由
	controller.MakeRouter(r)
	// 启动
	utils.UtilsLogger.Info("Server Run Success: 0.0.0.0:8080")
	_ = r.Run(":8080")
}
