package service_api_dataset

import (
	"net/http"
	"wzinc/dify"
)

func GetDatasetApi(rawQuery string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/datasets"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForwardWithHeader(url, header)
}

func PostDatasetApi(body []byte, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/datasets"
	return dify.GeneralPostForwardWithHeader(url, body, header)
}
