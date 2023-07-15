package handler

import (
	"fmt"
	"log"
	"message/verificationCode"
	"net/http"
)

func CheckCodeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	phoneNumber := r.FormValue("phone_number")
	vc := verificationCode.NewVerificationCode()
	ok := vc.Check(phoneNumber, code)
	if !ok {
		log.Println("验证码不正确！")
		fmt.Fprintf(w, "验证码不正确！")
		return
	}
	fmt.Fprintf(w, "验证码正确！")
}
