package service_api_app

import (
	"net/http"
	"wzinc/dify"
)

func ConversationRenameApi(cId string, body []byte, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/conversations/" + cId + "/name"
	return dify.GeneralPostForwardWithHeader(url, body, header)
}

func ConversationApi(rawQuery string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/conversations/"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForwardWithHeader(url, header)
}

func GetConversationDetailApi(cId string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/conversations/" + cId
	return dify.GeneralGetForwardWithHeader(url, header)
}

func DeleteConversationDetailApi(cId string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/conversations/" + cId
	return dify.GeneralDeleteForwardWithHeader(url, header)
}
