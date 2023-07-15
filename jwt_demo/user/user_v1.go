package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL.Path,
	})
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
