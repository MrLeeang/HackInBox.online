package views

import (
	"HackInBox.online/utils"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		utils.UtilsLogger.Info("授权成功")
		context.Next()
	}
}
