package app

import (
	"net/http"
	"wzinc/dify"
)

func CompletionMessageApi(appId string, body []byte, streaming bool) (string, int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return "", http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/completion-messages"
	if streaming {
		return url, http.StatusOK, []byte(""), nil
	}
	code, message, err := dify.GeneralPostForward(url, body)
	return url, code, message, err
}

func CompletionMessageStopApi(appId string, taskId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/completion-messages/" + taskId + "/stop"
	return dify.GeneralPostForward(url, nil)
}

func ChatMessageApi(appId string, body []byte, streaming bool) (string, int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return "", http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/chat-messages"
	if streaming {
		return url, http.StatusOK, []byte(""), nil
	}
	code, message, err := dify.GeneralPostForward(url, body)
	return url, code, message, err
}

func ChatMessageStopApi(appId string, taskId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/chat-messages/" + taskId + "/stop"
	return dify.GeneralPostForward(url, nil)
}
