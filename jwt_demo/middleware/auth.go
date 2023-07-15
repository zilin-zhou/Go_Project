package middleware

import (
	"github.com/gin-gonic/gin"
	"jwt_demo/jwt_plugin"
	"net/http"
)

// jwt中间件
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("access_token")
		data := &jwt_plugin.Data{}
		err := jwt_plugin.Verify(token, data)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "身份认证失败",
			})
			ctx.Abort()
		}
		ctx.Set("auth_info", data)
		ctx.Next()
	}
}
