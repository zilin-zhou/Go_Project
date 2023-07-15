package ocr

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"ocr/config"
)

// 内嵌
type GeneralBasicOCRRequest struct {
	*ocr.GeneralBasicOCRRequest
}

type GeneralBasicOCRResponse struct {
	*ocr.GeneralBasicOCRResponse
}

func NewGeneralBasicOCRRequest() *GeneralBasicOCRRequest {
	return &GeneralBasicOCRRequest{
		ocr.NewGeneralBasicOCRRequest(),
	}
}
func GeneralBasicQCR(request *GeneralBasicOCRRequest) (response *GeneralBasicOCRResponse, err error) {
	response = &GeneralBasicOCRResponse{}
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.OcrConf.GetString("OcrConf.Endpoint")
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := ocr.NewClient(credential, config.OcrConf.GetString("OcrConf.Region"), cpf)

	// 返回的resp是一个GeneralBasicOCRResponse的实例，与请求对象对应
	response.GeneralBasicOCRResponse, err = client.GeneralBasicOCR(request.GeneralBasicOCRRequest)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	return
}
