package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourseDetail(ctx *gin.Context) {
	//从URL获取参数
	//http://127.0.0.1:8080/course/v1/detail/10?name=golang
	id := ctx.Param("id")     //id=10
	name := ctx.Query("name") //name=golang
	ctx.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}
func GetCourseView(ctx *gin.Context) {

}
func AddCourse(ctx *gin.Context) {

}
func PublicCourse(ctx *gin.Context) {

}

// 文件上传 单文件上传
func UploadCourse(ctx *gin.Context) {
	//单文件上传
	//file, _ := ctx.FormFile("file")
	//ctx.SaveUploadedFile(file, "upload/"+file.Filename)
	//ctx.JSON(http.StatusOK, gin.H{
	//	"filename": file.Filename,
	//})

	//多文件上传
	form, _ := ctx.MultipartForm()
	files, _ := form.File["files"]
	list := make([]string, len(files))

	for i, file := range files {
		ctx.SaveUploadedFile(file, "upload/"+file.Filename)
		list[i] = file.Filename
	}
	ctx.JSON(http.StatusOK, gin.H{
		"fiels": list,
	})
}
