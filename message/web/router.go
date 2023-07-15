package web

import (
	"message/web/handler"
	"net/http"
)

func init() {
	http.HandleFunc("/send/verification/code", handler.VerificationCodeHandler)
	http.HandleFunc("/send/verificationv1/code", handler.VerificationCodeV1Handler)
	http.HandleFunc("/check/verificationv1/code", handler.CheckCodeHandler)
}
