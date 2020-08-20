package views

import (
	"HackInBox.online/service"
	"HackInBox.online/utils"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Request.Header.Get("token")
		verify := service.VerifyToken(tokenStr)
		utils.UtilsLogger.Info(verify)
		if verify == nil {
			var ret = map[string]interface{}{}
			retCode := utils.RetCode.LoginError
			ret["code"] = retCode
			ret["ret_message"] = utils.ErrorCodeMessage[retCode]
			context.JSON(200, gin.H{
				"ret": ret,
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
