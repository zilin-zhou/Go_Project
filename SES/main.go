package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ses "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
	"log"
	"mail/config"
	_ "mail/web"
	"net/http"
)

func main() {
	addr := config.SesSecret.GetString("Http.IP") + ":" + config.SesConf.GetString("Http.Port")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
func main1() {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		"SecretId",
		"SecretKey",
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ses.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := ses.NewClient(credential, "ap-guangzhou", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := ses.NewSendEmailRequest()

	request.FromEmailAddress = common.StringPtr("291253464@qq.com")
	request.ReplyToAddresses = common.StringPtr("")
	request.Destination = common.StringPtrs([]string{""})
	request.Template = &ses.Template{
		TemplateID:   common.Uint64Ptr(11),
		TemplateData: common.StringPtr(""),
	}
	request.Simple = &ses.Simple{
		Html: common.StringPtr(""),
		Text: common.StringPtr(""),
	}
	request.Subject = common.StringPtr("")
	request.Attachments = []*ses.Attachment{
		&ses.Attachment{
			FileName: common.StringPtr(""),
			Content:  common.StringPtr(""),
		},
	}
	request.Unsubscribe = common.StringPtr("")
	request.TriggerType = common.Uint64Ptr(1111)

	// 返回的resp是一个SendEmailResponse的实例，与请求对象对应
	response, err := client.SendEmail(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
