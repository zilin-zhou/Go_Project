package web

import (
	"File-cos/cos"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//获取签名
func getCredential(w http.ResponseWriter, r *http.Request) {
	res, err := cos.NewGetCredential().GetCredential()
	if err != nil {
		log.Println(err)
		fmt.Println(err.Error())
		return
	}
	bytes, _ := json.Marshal(res)
	fmt.Fprintf(w, string(bytes))
	return
}
func getRoleCredential(w http.ResponseWriter, r *http.Request) {
	res, err := cos.NewGetRoleCredential().GetCredential()
	if err != nil {
		log.Println(err)
		fmt.Println(err.Error())
		return
	}
	bytes, _ := json.Marshal(res)
	fmt.Fprintf(w, string(bytes))
	return
}
