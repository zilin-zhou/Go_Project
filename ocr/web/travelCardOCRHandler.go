package web

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"ocr/ocr"
)

func TravelCardOCRHandler(w http.ResponseWriter, r *http.Request) {
	var request = ocr.NewRecognizeTravelCardOCRRequest()
	//绑定表单内容
	image_url := r.FormValue("image_url")
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

	response, err := ocr.TravelCardOCR(request)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, response.ToJsonString())
}
