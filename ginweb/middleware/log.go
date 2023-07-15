package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// 全局中间件 负责打印URL内容
func LogInput(ctx *gin.Context) {
	requestBody, _ := ctx.GetRawData()
	//GetRawData函数 是io流读取，读取完之后就没有了，全局中间件处理在前所以要在复制一份进去，负责后续读取都为空
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))

	mp := make(map[string]interface{})
	mp["request_url"] = ctx.Request.URL
	mp["status"] = ctx.Writer.Status()
	mp["method"] = ctx.Request.Method
	mp["access_token"] = ctx.GetHeader("access_token")
	mp["body"] = string(requestBody)
	//log.Println(mp)
	ctx.Next()
}
