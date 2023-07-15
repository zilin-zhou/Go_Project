package handler

import (
	"fmt"
	"log"
	"mail/verficationcode"

	"net/http"
)

func CheckMailCode(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	destination := r.FormValue("destination")
	vc := verficationcode.VC
	ok := vc.Check(destination, code)
	if !ok {
		log.Println("验证码不正确！")
		fmt.Fprintf(w, "验证码不正确！")
		return
	}
	fmt.Fprintf(w, "验证码正确！")
}
