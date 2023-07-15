package sms

import "message/sms/tencent"

type smsInterface interface {
	Send(phoneNumber []string, signName string, tempName string, params []string) (string, error)
}

func GetSmsSender() smsInterface {
	return &tencent.Sms{}
}
