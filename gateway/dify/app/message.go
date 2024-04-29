package app

import (
	"net/http"
	"wzinc/dify"
)

func MessageMoreLikeThisApi(appId string, messageId string, rawQuery string, streaming bool) (string, int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return "", http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/completion-messages/" + messageId + "/more-like-this"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	if streaming {
		return url, http.StatusOK, []byte(""), nil
	}
	code, message, err := dify.GeneralGetForward(url)
	return url, code, message, err
}

func MessageSuggestedQuestionApi(appId string, messageId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/chat-messages/" + messageId + "/suggested-questions"
	return dify.GeneralGetForward(url)
}

func ChatMessageListApi(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/chat-messages"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func MessageFeedbackApi(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/feedbacks"
	return dify.GeneralPostForward(url, body)
}

func MessageAnnotationApi(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/annotations"
	return dify.GeneralPostForward(url, body)
}

func MessageAnnotationCountApi(appId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/annotations/count"
	return dify.GeneralGetForward(url)
}

func MessageApi(appId string, messageId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/messages/" + messageId
	return dify.GeneralGetForward(url)
}
