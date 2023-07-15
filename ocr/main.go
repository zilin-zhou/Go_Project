package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"log"
	"net/http"
	"ocr/config"
	ocr "ocr/ocr"
	_ "ocr/web"
)

func main() {
	addr := config.OcrConf.GetString("Http.IP") + ":" + config.OcrConf.GetString("Http.Port")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
func main1() {
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := ocr.NewGeneralBasicOCRRequest()
	request.ImageBase64 = common.StringPtr("1234567")
	request.ImageUrl = common.StringPtr("https://ocr-demo-1254418846.cos.ap-guangzhou.myqcloud.com/general/GeneralBasicOCR/GeneralBasicOCR1.jpg")
	request.Scene = common.StringPtr("12345")
	request.LanguageType = common.StringPtr("zh")
	request.IsPdf = common.BoolPtr(false)
	request.PdfPageNumber = common.Uint64Ptr(0)
	request.IsWords = common.BoolPtr(false)

	response, err := ocr.GeneralBasicQCR(request)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(response.ToJsonString())
}
