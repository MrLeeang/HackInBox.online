package controller

import (
	"HackInBox.online/controller/views"
	"github.com/gin-gonic/gin"
)

// MakRouter 是路由
func MakRouter(r *gin.Engine) {
	r.GET("/", views.ActionVMLIST)
	r.GET("/user", views.ActionUser)
}
