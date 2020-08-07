package views

import (
	"HackInBox.online/service"
	"HackInBox.online/utils"
	"github.com/gin-contrib/sessions"
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

func ActionUserDetailByTeamUuid(c *gin.Context) {
	teamUuid := c.Query("team_uuid")
	ret := service.GetUserByTeamUuid(string(teamUuid))
	c.JSON(200, gin.H{
		"ret": ret,
	})
}

func ActionLogin(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	// 定义登录参数
	type LoginForm struct {
		Email    string `form:"user" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var form LoginForm

	// 获取登录参数
	if err := c.ShouldBindJSON(&form); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		utils.UtilsLogger.Error(err)
		retCode := utils.RetCode.HttpBadRequest
		ret["code"] = retCode
		ret["error_message"] = utils.ErrorCodeMessage[retCode]
		c.JSON(retCode, gin.H{"ret": ret})
		return
	}

	userModels := service.GetUserByEmail(string(form.Email))

	if userModels == nil || userModels[0].Password != form.Password {
		// 账号密码错误，登录失败
		retCode := utils.RetCode.HttpUnAuthorized
		ret["code"] = retCode
		ret["error_message"] = utils.ErrorCodeMessage[retCode]
		c.JSON(retCode, gin.H{"ret": ret})
		return
	}
	// 登录成功
	retCode := utils.RetCode.Success
	ret["code"] = retCode
	ret["error_message"] = utils.ErrorCodeMessage[retCode]

	// 返回数据
	var data = map[string]interface{}{}
	data["Name"] = userModels[0].Name
	data["Id"] = userModels[0].Id
	data["Email"] = userModels[0].Email
	data["DisplayName"] = userModels[0].DisplayName
	data["Name"] = userModels[0].Name
	data["TeamUuid"] = userModels[0].TeamUuid
	data["CreatedAt"] = userModels[0].CreatedAt
	// 把data添加到返回值
	ret["data"] = data
	// 写入session
	session := sessions.Default(c)
	session.Set("loginName", userModels[0].Name)
	_ = session.Save()
	c.JSON(retCode, gin.H{"ret": ret})
	return
}

func ActionRegister(c *gin.Context) {
	// 定义返回值
	var ret = map[string]interface{}{}

	c.JSON(200, gin.H{
		"ret": ret,
	})
}
