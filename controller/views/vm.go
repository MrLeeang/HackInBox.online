package views

import (
	"HackInBox.online/service"
	"HackInBox.online/utils"
	"github.com/gin-gonic/gin"
)

// ActionVMLIST 查询虚拟机列表
func ActionVMLIST(c *gin.Context) {
	ret := service.GetUsers()
	utils.Utilslogger.Info(ret)
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
