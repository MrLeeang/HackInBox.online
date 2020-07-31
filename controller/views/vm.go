package views

import (
	"github.com/gin-gonic/gin"
)

// ActionVMLIST 查询虚拟机列表
func ActionVMLIST(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
