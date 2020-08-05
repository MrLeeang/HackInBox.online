package views

import (
	"HackInBox.online/service"
	"github.com/gin-gonic/gin"
)

// ActionVMLIST 查询虚拟机列表
func ActionUser(c *gin.Context) {
	userUuid := c.Query("uuid")
	ret := service.GetUserByUuid(string(userUuid))
	c.JSON(200, gin.H{
		"ret": ret,
	})
}

func ActionUserList(c *gin.Context) {
	ret := service.GetUsers()
	c.JSON(200, gin.H{
		"ret": ret,
	})
}
