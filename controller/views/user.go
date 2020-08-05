package views

import (
	"HackInBox.online/service"
	"github.com/gin-gonic/gin"
)

// ActionVMLIST 查询虚拟机列表
func ActionUser(c *gin.Context) {
	ret := service.GetUserByUuid("887f1149-d62d-11ea-ac10-0242c0a80007")
	c.JSON(200, gin.H{
		"ret": ret,
	})
}
