package web

import (
	"mail/web/handler"
	"net/http"
)

func init() {
	http.HandleFunc("/send/email", handler.SendMailHandler) //发送邮件
	http.HandleFunc("/send/email/verification/code", handler.VerificationCodeHandler)
	http.HandleFunc("/verification/email/code", handler.CheckMailCode) //发送邮件验证码
}
