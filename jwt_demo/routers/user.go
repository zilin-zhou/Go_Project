package routers

import (
	"github.com/gin-gonic/gin"
	"jwt_demo/user"
)

func InitUser(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/user/:id", user.Get)
		v1.POST("/user", user.Add)
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/user/:id", user.GetV2)
		v2.POST("/user", user.AddV2)
	}
}
