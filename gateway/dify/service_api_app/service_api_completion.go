package service_api_app

import (
	"net/http"
	"wzinc/dify"
)

func CompletionApi(body []byte, header http.Header, streaming bool) (string, int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/completion-messages"
	if streaming {
		return url, http.StatusOK, []byte(""), nil
	}
	code, message, err := dify.GeneralPostForwardWithHeader(url, body, header)
	return url, code, message, err
}

func CompletionStopApi(taskId string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/completion-messages/" + taskId + "/stop"
	return dify.GeneralPostForwardWithHeader(url, nil, header)
}

func ChatApi(body []byte, header http.Header, streaming bool) (string, int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/chat-messages"
	if streaming {
		return url, http.StatusOK, []byte(""), nil
	}
	code, message, err := dify.GeneralPostForwardWithHeader(url, body, header)
	return url, code, message, err
}

func ChatStopApi(taskId string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/chat-messages/" + taskId + "/stop"
	return dify.GeneralPostForwardWithHeader(url, nil, header)
}
