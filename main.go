package main

import (
	"HackInBox.online/controller"
	"HackInBox.online/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	r := gin.Default()

	r.Use(middleware.LoggerToFile())

	controller.MakRouter(r)
	r.Run(":8080")
}
