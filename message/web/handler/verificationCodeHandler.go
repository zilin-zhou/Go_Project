package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"message/sms"
	"message/verificationCode"
	"net/http"
)

func VerificationCodeHandler(w http.ResponseWriter, r *http.Request) {
	defaultSignName := "盈水展千华公众号"
	tempName := "验证码"
	signName := r.FormValue("sign_name")
	if signName == "" {
		signName = defaultSignName
	}

	sender := sms.GetSmsSender()

	phoneNumbersStr := r.FormValue("phone_numbers")
	//传递过来的为json，将其转为string
	phoneNumbers := make([]string, 0)
	err := json.Unmarshal([]byte(phoneNumbersStr), &phoneNumbers)
	if err != nil {
		log.Fatalln(err)
	}
	paramsStr := r.FormValue("params")
	//传递过来的为json，将其转为string
	params := make([]string, 0)
	if paramsStr != "" {
		err := json.Unmarshal([]byte(paramsStr), &params)
		if err != nil {
			log.Fatalln(err)
		}
	}
	res, err := sender.Send(phoneNumbers, signName, tempName, params)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, res)
}

// 引入数据库  验证码为随机验证码
func VerificationCodeV1Handler(w http.ResponseWriter, r *http.Request) {
	defaultSignName := "盈水展千华公众号"
	tempName := "验证码"
	signName := r.FormValue("sign_name")
	if signName == "" {
		signName = defaultSignName
	}

	phoneNumber := r.FormValue("phone_number")
	//传递过来的为json，将其转为string
	vc := verificationCode.NewVerificationCode()

	res, err := vc.Send(phoneNumber, tempName, signName)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, res)
}
