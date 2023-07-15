package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthCheck(ctx *gin.Context) {
	username, _ := ctx.Get("user_name")
	userID, _ := ctx.Get("user_id")
	fmt.Printf("auth check called,userID:%v,user_name:%v\n", userID, username)
	ctx.Next()
}
