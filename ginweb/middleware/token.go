package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路径中间件  认证token
var token = "123456"

func TokenCheck(ctx *gin.Context) {
	accessToken := ctx.GetHeader("access_token")
	if accessToken != token {
		ctx.JSONP(http.StatusInternalServerError, gin.H{
			"message": "token check failed!",
		})
		//ctx.Abort() //不再执行后续的中间件
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("token check failed!"))
	}
	ctx.Set("user_name", "zilin")
	ctx.Set("user_id", "100")
	ctx.Next()
}
