package login

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"jwt_demo/jwt_plugin"
	"net/http"
	"time"
)

//type User struct {
//	Name   string `json:"name,omitempty"`
//	Age    int    `json:"age,omitempty"`
//	Gender int    `json:"gender,omitempty"`
//}

func Login(ctx *gin.Context) {
	//user := &User{}
	//err := ctx.Bind(user)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"error": "参数绑定错误!",
	//	})
	//	return
	//}
	data := &jwt_plugin.Data{
		//Name:   user.Name,
		//Age:    user.Age,
		//Gender: user.Gender,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	ctx.Bind(data)
	sign, err := jwt_plugin.Sign(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"access_token": sign,
	})

}
