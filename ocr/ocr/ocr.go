package ocr

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"ocr/config"
)

var credential *common.Credential

func init() {
	credential = common.NewCredential(
		config.OcrSecret.GetString("OcrSecret.SecretId"),
		config.OcrSecret.GetString("OcrSecret.SecretKey"),
	)

}
