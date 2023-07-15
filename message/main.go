package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"log"
	"message/config"
	_ "message/web"
	"net/http"
)

func main() {
	addr := config.SmsConf.GetString("Http.IP") + ":" + config.SmsConf.GetString("Http.Port")
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
		config.SmsSecret.GetString("Credential.SecretId"),
		config.SmsSecret.GetString("Credential.SecretKey"),
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.SmsConf.GetString("SmsConf.Endpoint")
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, config.SmsConf.GetString("SmsConf.Region"), cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs([]string{"13353271663", "13523321524", "13800138000"})
	request.SmsSdkAppId = common.StringPtr("1400823653")
	request.SignName = common.StringPtr("123")
	request.TemplateId = common.StringPtr("1231")
	request.TemplateParamSet = common.StringPtrs([]string{"1234", "3"})

	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := client.SendSms(request)
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
