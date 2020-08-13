package controller

import (
	"HackInBox.online/controller/views"
	"github.com/gin-gonic/gin"
)

// MakRouter 是路由
func MakRouter(r *gin.Engine) {
	// 登录
	r.POST("/login", views.ActionLogin)
	// 注册
	r.POST("/register", views.ActionRegister)
	// 用户信息
	r.GET("/user", views.AuthHandler(), views.ActionUser)
	// 用户列表
	r.GET("/user/list", views.AuthHandler(), views.ActionUserList)
	// 战队成员
	r.GET("/team/user", views.AuthHandler(), views.ActionUserDetailByTeamUuid)
	// 虚拟机列表
	r.GET("/", views.AuthHandler(), views.ActionVMLIST)

}
