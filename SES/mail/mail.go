package mail

import (
	"mail/mail/input"
	"mail/mail/tencent"
)

type MailSender interface {
	Send(input *input.ReqInput) (string, error)
}

func NewTencentMailSender() MailSender {
	return &tencent.TencentSes{}
}
