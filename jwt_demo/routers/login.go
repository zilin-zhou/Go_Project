package routers

import (
	"github.com/gin-gonic/gin"
	"jwt_demo/login"
)

func InitLogin(r *gin.RouterGroup) {
	v1 := r.Group("v1")
	{
		v1.POST("/login", login.Login)
	}
}
