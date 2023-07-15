package jwt_plugin

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
)

var key = "123aasdasdeff"

type Data struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	Gender int    `json:"gender,omitempty"`
	jwt.RegisteredClaims
}

func Sign(data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	sign, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return sign, err
}

func Verify(sign string, data jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, data, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return err
}
