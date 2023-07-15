package web

import "net/http"

func init() {
	http.HandleFunc("/ajax/upload", AjaxUploadHandeer)
	http.HandleFunc("/get/role/credential", getCredential)
	http.HandleFunc("/get/role/rolecredential", getRoleCredential)
}
