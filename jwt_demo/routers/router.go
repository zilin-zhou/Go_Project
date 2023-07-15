package routers

import (
	"github.com/gin-gonic/gin"
	"jwt_demo/middleware"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.Cors(), middleware.Auth())
	{
		//课程相关接口
		InitCourse(api)
		//用户相关接口
		InitUser(api)
	}
	notAuthApi := r.Group("api")
	notAuthApi.Use(middleware.Cors())
	//登录
	InitLogin(notAuthApi)
}
