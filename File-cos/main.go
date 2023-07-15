package main

import (
	"File-cos/config"
	_ "File-cos/web"
	"log"
	"net/http"
)

func main() {
	//监听端口
	addr := config.Confs.GetString("HttpConf.IP") + ":" + config.Confs.GetString("HttpConf.Port")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println(err.Error())
	}

}
