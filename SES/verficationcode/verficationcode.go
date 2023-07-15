package verficationcode

import (
	"encoding/json"
	"log"
	"mail/mail"
	"mail/mail/input"
	"mail/utils"
	"time"
)

var verficationcodeKeyPrefix string = "verficationcode_"
var VC *verficationCode

func init() {
	VC = NewVerificationCode()
}

type verficationCode struct {
	s storge
}

type storge interface {
	Set(key, val string, duration time.Duration) error
	Get(key string) (string, error)
}

func (vc *verficationCode) getKey(mailAddress string) string {
	return verficationcodeKeyPrefix + mailAddress
}

// 检查验证码
func (vc *verficationCode) Check(mailAddress, code string) bool {
	key := vc.getKey(mailAddress)
	val, err := vc.s.Get(key)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if val != code {
		return false
	}
	return true
}

// 发送验证码
func (vc *verficationCode) Send(input *input.ReqInput) (string, error) {
	sender := mail.NewTencentMailSender()
	code := utils.GeneratorRandNo(4)
	mp := make(map[string]string)
	mp["code"] = code //产生随机验证码
	mp["m"] = "10"    //有效时间/分钟
	td, _ := json.Marshal(mp)
	input.TemplateData = string(td)
	key := vc.getKey(input.Destination[0])
	err := vc.s.Set(key, code, time.Second*600)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	res, err := sender.Send(input)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return res, nil
}

func NewVerificationCode() *verficationCode {
	return &verficationCode{
		//NewRedisStorage(), //使用redis数据库存值
		NewMapStorage(), //使用自带的map存储
	}
}
