package router

import (
	"ginweb/middleware"
	"ginweb/web"
	"github.com/gin-gonic/gin"
)

func InitCourse(r *gin.Engine) {
	course := r.Group("/course", middleware.TokenCheck)
	v1 := course.Group("/v1")
	{
		//http://localhost:8080/course/v1/detail/10
		v1.GET("/detail/:id", web.GetCourseDetail)
		v1.POST("/view/:id", web.GetCourseView)
	}
	admin := course.Group("/admin", middleware.AuthCheck)
	adminV1 := admin.Group("/v1")
	{
		adminV1.POST("/add", web.AddCourse)
		adminV1.POST("/public", web.PublicCourse)
		adminV1.POST("/upload", web.UploadCourse)
	}
}
