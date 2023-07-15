package input

type Attachment struct {
	FileName string
	Content  string
}
type ReqInput struct {
	FromEmailAddress string        //发信邮件地址。请填写发件人邮箱地址
	Destination      []string      //收信人邮箱地址
	TemplateName     string        //模板名称
	TemplateData     string        //模板内容
	Subject          string        //邮件主题
	Unsubscribe      string        //退订链接选项 0: 不加入退订链接 1: 简体中文 2: 英文 3: 繁体中文 4: 西班牙语 5: 法语 6: 德语 7: 日语 8: 韩语 9: 阿拉伯语
	TriggerType      uint64        //邮件触发类型 0:非触发类，默认类型，营销类邮件、非即时类邮件等选择此类型 1:触发类，验证码等即时发送类邮件
	Attachments      []*Attachment //需要发送附件时，填写附件相关参数
}
