package ocrV1

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"ocr/config"
	"strconv"
)

type GeneralBasic struct {
}

func (g *GeneralBasic) Ocr(input map[string]string) (string, error) {
	//得到凭证    getCredential
	credential := g.getCredential()
	g.getCredential()
	//得到请求体	getRequest
	request := g.getRequest(input)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.OcrConf.GetString("OcrConf.Endpoint")
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := ocr.NewClient(credential, config.OcrConf.GetString("OcrConf.Region"), cpf)

	// 返回的resp是一个GeneralBasicOCRResponse的实例，与请求对象对应
	response, err := client.GeneralBasicOCR(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", nil
	}

	//得到响应体	getResponse
	res := g.getResponse(response)
	return res, nil
}

func (g *GeneralBasic) getCredential() *common.Credential {
	return common.NewCredential(
		config.OcrSecret.GetString("OcrSecret.SecretId"),
		config.OcrSecret.GetString("OcrSecret.SecretKey"),
	)
}

func (g *GeneralBasic) getRequest(input map[string]string) *ocr.GeneralBasicOCRRequest {
	request := ocr.NewGeneralBasicOCRRequest()
	imageUrl := input["image_url"]
	scene := input["secne"]
	languageType := input["language_type"]
	isPdf, _ := strconv.ParseBool(input["is_pdf"])
	pdfPageNumber, _ := strconv.ParseUint(input["pdf_page_number"], 10, 64)
	isWords, _ := strconv.ParseBool(input["is_words"])
	imageBase64 := input["inage_base64"]
	//赋值
	request.ImageUrl = &imageUrl
	request.Scene = &scene
	request.LanguageType = &languageType
	request.IsPdf = &isPdf
	request.IsWords = &isWords
	request.PdfPageNumber = &pdfPageNumber
	request.ImageBase64 = &imageBase64
	return request
}

func (g *GeneralBasic) getResponse(response *ocr.GeneralBasicOCRResponse) string {
	return response.ToJsonString()
}
