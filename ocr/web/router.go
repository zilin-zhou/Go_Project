package web

import "net/http"

func init() {
	http.HandleFunc("/orc/general/basic", GeneralBasicQCRHandler)
	http.HandleFunc("/orc/health/code", HealthCodeOCRHandler)
	http.HandleFunc("/orc/travel/card", TravelCardOCRHandler)
	http.HandleFunc("/orcv1/general/basic", GeneralBasicQCRHandlerV1)

}
