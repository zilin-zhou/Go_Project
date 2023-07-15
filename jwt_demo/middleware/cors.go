package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//gin的跨域访问

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowHeaders: []string{
				"Origin", "Content-length", "Content-Type",
			},
			AllowMethods: []string{
				"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD",
			},
		})
}
