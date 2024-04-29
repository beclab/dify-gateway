package service_api_app

import (
	"net/http"
	"wzinc/dify"
)

func MessageListApi(rawQuery string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/messages"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForwardWithHeader(url, header)
}

func MessageFeedbackApi(messageId string, body []byte, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/messages/" + messageId + "/feedbacks"
	return dify.GeneralPostForwardWithHeader(url, body, header)
}

func MessageSuggestedApi(messageId string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/messages/" + messageId + "/suggested"
	return dify.GeneralGetForwardWithHeader(url, header)
}
