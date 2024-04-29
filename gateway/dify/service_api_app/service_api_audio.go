package service_api_app

import (
	"net/http"
	"wzinc/dify"
)

func AudioApi(filename string, content string) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/audio-to-text"
	return dify.FilePostForward(url, filename, content)
}

func TextApi(body []byte, header http.Header, streaming bool) (string, int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/text-to-audio"
	if streaming {
		return url, http.StatusOK, []byte(""), nil
	}
	code, message, err := dify.GeneralPostForwardWithHeader(url, body, header)
	return url, code, message, err
}
