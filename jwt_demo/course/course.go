package course

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Course struct {
	Name     string `json:"name" form:"name" `
	Teacher  string `json:"teacher" form:"teacher"`
	Duration string `json:"duration" form:"duration"`
}

// 增
func Add(ctx *gin.Context) {
	course := &Course{}
	err := ctx.Bind(course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "参数绑定错误!",
		})
		return
	}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"method": ctx.Request.Method,
	//	"url":    ctx.Request.URL.Path,
	//	"course": course,
	//})
	out := &AddCourseResponse{
		Name:     course.Name,
		Teacher:  course.Teacher,
		Duration: course.Duration,
	}
	//返回protobuf对象
	ctx.ProtoBuf(http.StatusOK, out)
}

// 删
func Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL.Path,
	})
}

// 改
func Updata(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL.Path,
	})
}

// 查
func Get(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"id":     id,
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL.Path,
	})
}
