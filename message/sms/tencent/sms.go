package tencent

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"log"
	"message/config"
	"message/sms/tencent/template"
)

type Sms struct {
}

func getCredential() *common.Credential {
	//实例化证书对象 传入
	return common.NewCredential(
		config.SmsSecret.GetString("Credential.SecretId"),
		config.SmsSecret.GetString("Credential.SecretKey"),
	)
}

// 获取client
func getClient() *sms.Client {
	credential := getCredential()
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.SmsConf.GetString("SmsConf.Endpoint")
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, config.SmsConf.GetString("SmsConf.Region"), cpf)
	return client
}

func (s *Sms) Send(phoneNumber []string, signName string, tempName string, params []string) (string, error) {
	t, ok := template.Templates[tempName]
	if !ok {
		log.Fatalln("没有短信模板")
	}
	client := getClient()
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs(phoneNumber)
	request.SmsSdkAppId = common.StringPtr(config.SmsConf.GetString("SmsConf.SmsSdkAppID"))
	request.SignName = common.StringPtr(signName)
	request.TemplateId = common.StringPtr(t.ID)
	request.TemplateParamSet = common.StringPtrs(params)

	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", err
	}
	if err != nil {
		panic(err)
	}
	return response.ToJsonString(), nil
}
