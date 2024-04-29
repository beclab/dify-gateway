package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"wzinc/api"
	"wzinc/dify/app"
)

func CompletionConversationApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.CompletionConversationApi(appId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app completion conversation API")
	return
}

func CompletionConversationDetailApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	conversationId := vars["conversation_id"]

	if r.Method == http.MethodGet {
		code, response, err := app.GetCompletionConversationDetailApi(appId, conversationId)
		api.GeneralGetResponse(w, code, response, err, "Failed to call get app completion conversation detail API")
		return
	}

	if r.Method == http.MethodDelete {
		code, response, err := app.DeleteCompletionConversationDetailApi(appId, conversationId)
		api.GeneralGetResponse(w, code, response, err, "Failed to call delete app completion conversation detail API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func ChatConversationApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.ChatConversationApi(appId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app chat conversation API")
	return
}

func ChatConversationDetailApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	conversationId := vars["conversation_id"]

	if r.Method == http.MethodGet {
		code, response, err := app.GetChatConversationDetailApi(appId, conversationId)
		api.GeneralGetResponse(w, code, response, err, "Failed to call get app chat conversation detail API")
		return
	}

	if r.Method == http.MethodDelete {
		code, response, err := app.DeleteChatConversationDetailApi(appId, conversationId)
		api.GeneralGetResponse(w, code, response, err, "Failed to call delete app chat conversation detail API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}
