package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mail/mail"
	"mail/mail/input"
	"net/http"
	"strconv"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	reqInput := &input.ReqInput{}
	triggerType, _ := strconv.ParseUint(r.FormValue("trigger_type"), 10, 64)
	reqInput.TriggerType = triggerType
	reqInput.Subject = r.FormValue("subject")
	reqInput.Unsubscribe = r.FormValue("unsubscribe")
	reqInput.TemplateName = r.FormValue("template_name")
	reqInput.TemplateData = r.FormValue("template_data")
	reqInput.FromEmailAddress = r.FormValue("from_email_address")
	destination := r.FormValue("destination")
	//接受的json数据，进行反序列化
	json.Unmarshal([]byte(destination), &reqInput.Destination)
	files := r.MultipartForm.File["attachments"]
	count := len(files)
	if count > 0 {
		reqInput.Attachments = make([]*input.Attachment, count)
		for i := 0; i < count; i++ {
			filename := files[i].Filename
			file, _ := files[i].Open()
			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println(err)
				fmt.Fprintln(w, err.Error())
				return
			}
			content := base64.StdEncoding.EncodeToString(bytes)
			reqInput.Attachments[i] = &input.Attachment{
				FileName: filename,
				Content:  content,
			}

		}
	}
	send := mail.NewTencentMailSender()
	res, err := send.Send(reqInput)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, res)
}
