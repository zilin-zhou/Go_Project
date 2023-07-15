package router

import (
	"ginweb/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.Use(middleware.LogInput) //全局中间件的使用
	InitApi(r)
	InitCourse(r)
}
