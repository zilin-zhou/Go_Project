package ocr

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"ocr/config"
)

// 内嵌
type RecognizeTravelCardOCRRequest struct {
	*ocr.RecognizeTravelCardOCRRequest
}

type RecognizeTravelCardOCRResponse struct {
	*ocr.RecognizeTravelCardOCRResponse
}

func NewRecognizeTravelCardOCRRequest() *RecognizeTravelCardOCRRequest {
	return &RecognizeTravelCardOCRRequest{
		ocr.NewRecognizeTravelCardOCRRequest(),
	}
}
func TravelCardOCR(request *RecognizeTravelCardOCRRequest) (response *RecognizeTravelCardOCRResponse, err error) {
	response = &RecognizeTravelCardOCRResponse{}
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.OcrConf.GetString("OcrConf.Endpoint")
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := ocr.NewClient(credential, config.OcrConf.GetString("OcrConf.Region"), cpf)

	// 返回的resp是一个GeneralBasicOCRResponse的实例，与请求对象对应
	response.RecognizeTravelCardOCRResponse, err = client.RecognizeTravelCardOCR(request.RecognizeTravelCardOCRRequest)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	return
}
