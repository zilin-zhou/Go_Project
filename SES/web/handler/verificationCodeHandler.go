package handler

import (
	"fmt"
	"log"
	"mail/mail/input"
	"mail/verficationcode"
	"net/http"
)

func VerificationCodeHandler(w http.ResponseWriter, r *http.Request) {
	reqInput := &input.ReqInput{}
	reqInput.TriggerType = 1
	reqInput.Subject = r.FormValue("subject")
	reqInput.Unsubscribe = "0"
	reqInput.TemplateName = "xxxxxx"          //验证码需要固定的模板/根据自己的邮件模板名
	reqInput.FromEmailAddress = "xxxxxx@.com" //验证码需要把发件人固定
	destination := r.FormValue("destination")
	reqInput.Destination = []string{destination}

	vc := verficationcode.VC
	res, err := vc.Send(reqInput)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, res)
}
