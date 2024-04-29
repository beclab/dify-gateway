package service_api_app

import (
	"net/http"
	"wzinc/dify"
)

func AppParameterApi(rawQuery string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/parameters"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForwardWithHeader(url, header)
}
