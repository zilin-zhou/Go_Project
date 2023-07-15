package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL.Path,
	})
}

// 删
func DeleteV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL.Path,
	})
}

// 改
func UpdataV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"method": ctx.Request.Method,
		"url":    ctx.Request.URL.Path,
	})
}

// 查
func GetV2(ctx *gin.Context) {
	auth_info, _ := ctx.Get("auth_info")
	ctx.JSON(http.StatusOK, auth_info)
}
