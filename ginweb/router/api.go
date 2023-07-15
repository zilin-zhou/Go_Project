package router

import (
	"ginweb/web"
	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {
	//分组
	api := r.Group("/api")
	v1 := api.Group("/v1")
	//创建路由
	{
		v1.GET("/ping", web.Ping)
		v1.POST("/login", web.Login)
		v1.POST("/register", web.Register)
	}

}
