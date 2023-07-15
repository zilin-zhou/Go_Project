package tencent

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ses "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
	"mail/config"
	"mail/mail/input"
	"mail/mail/tencent/template"
)

type TencentSes struct {
}

func getCredential() *common.Credential {
	//实例化证书对象 传入
	return common.NewCredential(
		config.SesSecret.GetString("Credential.SecretId"),
		config.SesSecret.GetString("Credential.SecretKey"),
	)
}

// 获取client
func getClient() *ses.Client {
	credential := getCredential()
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.SesConf.GetString("SesConf.Endpoint")
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := ses.NewClient(credential, config.SesConf.GetString("SesConf.Region"), cpf)
	return client
}

func (ts *TencentSes) Send(input *input.ReqInput) (string, error) {

	request := ses.NewSendEmailRequest()
	request.FromEmailAddress = common.StringPtr(input.FromEmailAddress)
	request.Destination = common.StringPtrs(input.Destination)
	request.Template = &ses.Template{
		TemplateID:   common.Uint64Ptr(template.Templates[input.TemplateName].ID),
		TemplateData: common.StringPtr(input.TemplateData),
	}

	request.Subject = common.StringPtr(input.Subject)
	request.Attachments = []*ses.Attachment{}
	for _, attachment := range input.Attachments {
		request.Attachments = append(request.Attachments, &ses.Attachment{
			FileName: common.StringPtr(attachment.FileName),
			Content:  common.StringPtr(attachment.Content),
		})
	}
	request.Unsubscribe = common.StringPtr(input.Unsubscribe)
	request.TriggerType = common.Uint64Ptr(input.TriggerType)
	client := getClient()
	// 返回的resp是一个SendEmailResponse的实例，与请求对象对应
	response, err := client.SendEmail(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", err
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	return response.ToJsonString(), nil
}
