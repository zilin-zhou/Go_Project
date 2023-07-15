package main

import (
	"github.com/gin-gonic/gin"
	"jwt_demo/routers"
)

func main() {
	r := gin.Default()
	routers.InitRouter(r)
	r.Run()
}
