package views

import (
	"HackInBox.online/utils"
	"github.com/gin-gonic/gin"
)

// ActionVMLIST 查询虚拟机列表
func ActionVMLIST(c *gin.Context) {
	utils.Utilslogger.Info("fdsfd")
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
