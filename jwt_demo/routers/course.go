package routers

import (
	"github.com/gin-gonic/gin"
	"jwt_demo/course"
)

func InitCourse(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		//路径传参
		v1.GET("/course/:id", course.Get)
		//添加
		v1.POST("/course", course.Add)
		//更新
		v1.PUT("/course", course.Updata)
		//删除
		v1.DELETE("/course", course.Delete)
	}
}
