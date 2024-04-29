package app

import (
	"net/http"
	"wzinc/dify"
)

func CompletionConversationApi(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/completion-conversations"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func GetCompletionConversationDetailApi(appId string, conversationId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/completion-conversations/" + conversationId
	return dify.GeneralGetForward(url)
}

func DeleteCompletionConversationDetailApi(appId string, conversationId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/completion-conversations/" + conversationId
	return dify.GeneralDeleteForward(url)
}

func ChatConversationApi(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/chat-conversations"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func GetChatConversationDetailApi(appId string, conversationId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/chat-conversations/" + conversationId
	return dify.GeneralGetForward(url)
}

func DeleteChatConversationDetailApi(appId string, conversationId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/chat-conversations/" + conversationId
	return dify.GeneralDeleteForward(url)
}
