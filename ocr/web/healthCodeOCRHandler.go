package web

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"ocr/ocr"
	"strconv"
)

func HealthCodeOCRHandler(w http.ResponseWriter, r *http.Request) {
	var request = ocr.NewRecognizeHealthCodeOCRRequest()
	//绑定表单内容
	image_url := r.FormValue("image_url")
	Type, _ := strconv.ParseInt(r.FormValue("type"), 10, 64)
	request.Type = &Type
	if image_url == "" {
		file, _, err := r.FormFile("image_file")
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		//转化为base64
		imageBase64 := base64.StdEncoding.EncodeToString(bytes)
		request.ImageBase64 = &imageBase64
	}
	request.ImageUrl = &image_url

	response, err := ocr.HealthCodeOCR(request)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, response.ToJsonString())

}
