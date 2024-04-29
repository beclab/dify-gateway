package app

import (
	"net/http"
	"wzinc/dify"
)

func ChatMessageAudioApi(appId string, filename string, content string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/audio-to-text"
	return dify.FilePostForward(url, filename, content)
}

func ChatMessageTextApi(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/text-to-audio"
	return dify.GeneralPostForward(url, body)
}

func TextModesApi(appId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/text-to-audio/voices"
	return dify.GeneralGetForward(url)
}
