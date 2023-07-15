package verificationCode

import (
	"fmt"
	"log"
	"math/rand"
	"message/sms"
	"time"
)

var verificationKeyPrefix = "verification"

type verificationCode struct {
	s storge
}
type storge interface {
	Set(key, val string, duration time.Duration) error
	Get(key string) (string, error)
}

func (vc *verificationCode) getKey(phoneNumber string) string {
	return verificationKeyPrefix + phoneNumber
}

// 验证验证码是否正确
func (vc *verificationCode) Check(phoneNumber, code string) bool {
	key := vc.getKey(phoneNumber)
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

func (vc *verificationCode) Send(phoneNumber, tempName, signName string) (string, error) {
	sender := sms.GetSmsSender()
	phoneNumbers := []string{phoneNumber}
	//随机产生验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	//code:=utils.GeneratorRandNo(4)
	key := vc.getKey(phoneNumber)
	err := vc.s.Set(key, code, time.Second*300)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	params := []string{code}
	//发送短信
	res, err := sender.Send(phoneNumbers, signName, tempName, params)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return res, nil
}

func NewVerificationCode() *verificationCode {
	return &verificationCode{
		s: NewRedisStorage(),
	}
}
