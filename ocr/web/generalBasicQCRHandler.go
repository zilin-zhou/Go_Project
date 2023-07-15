package web

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"ocr/ocr"
	"ocr/ocrV1"
	"strconv"
)

// 路由
func GeneralBasicQCRHandler(w http.ResponseWriter, r *http.Request) {
	var request = ocr.NewGeneralBasicOCRRequest()
	//绑定表单内容
	imageUrl := r.FormValue("image_url")
	scene := r.FormValue("secne")
	languageType := r.FormValue("language_type")
	isPdf, _ := strconv.ParseBool(r.FormValue("is_pdf"))
	pdfPageNumber, _ := strconv.ParseUint(r.FormValue("pdf_page_number"), 10, 64)
	isWords, _ := strconv.ParseBool(r.FormValue("is_words"))

	//赋值
	request.ImageUrl = &imageUrl
	request.Scene = &scene
	request.LanguageType = &languageType
	request.IsPdf = &isPdf
	request.IsWords = &isWords
	request.PdfPageNumber = &pdfPageNumber
	/*
		图片/PDF的 Base64 值。 要求图片/PDF经Base64编码后不超过 7M，分辨率建议600*800以上，
		支持PNG、JPG、JPEG、BMP、PDF格式。 图片的 ImageUrl、ImageBase64
		必须提供一个，如果都提供，只使用 ImageUrl
	*/
	if imageUrl == "" {
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
	response, err := ocr.GeneralBasicQCR(request)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, response.ToJsonString())
}

func GeneralBasicQCRHandlerV1(w http.ResponseWriter, r *http.Request) {
	gb := &ocrV1.GeneralBasic{}
	input := make(map[string]string)
	//绑定表单内容
	input["image_url"] = r.FormValue("image_url")
	input["secne"] = r.FormValue("secne")
	input["language_type"] = r.FormValue("language_type")
	input["is_pdf"] = r.FormValue("is_pdf")
	input["pdf_page_number"] = r.FormValue("pdf_page_number")
	input["is_words"] = r.FormValue("is_words")

	if input["image_url"] == "" {
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
		input["image_ase64"] = imageBase64
	}
	response, err := gb.Ocr(input)

	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, response)
}
