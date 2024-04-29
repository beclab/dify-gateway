package service_api_app

import (
	"net/http"
	"wzinc/api"
	"wzinc/dify/service_api_app"
)

func AppParameterApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := service_api_app.AppParameterApi(r.URL.RawQuery, r.Header)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app parameter API")
	return
}
