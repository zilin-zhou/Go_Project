package web

import (
	"html/template"
	"net/http"
)

func AjaxUploadHandeer(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web/template/ajaxUpload.html")
	t.Execute(w, nil)
}
