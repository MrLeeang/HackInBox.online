package views

import (
	"HackInBox.online/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		loginName := session.Get("loginName")
		_ = session.Save()
		if loginName == nil {
			var ret = map[string]interface{}{}
			retCode := utils.RetCode.VerifyFailed
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
