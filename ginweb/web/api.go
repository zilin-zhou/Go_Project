package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerReq struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Phone    string `form:"phone" binding:"required,e164"`
	Email    string `form:"email" binding:"omitempty,email"`
}

func Ping(cctx *gin.Context) {

}
func Login(cctx *gin.Context) {

}

func Register(ctx *gin.Context) {
	req := &registerReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, req)
}
