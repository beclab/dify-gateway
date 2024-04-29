package service_api_app

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/service_api_app"
)

func MessageListApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := service_api_app.MessageListApi(r.URL.RawQuery, r.Header)
	api.GeneralGetResponse(w, code, response, err, "Failed to call message list API")
	return
}

func MessageFeedbackApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	messageId := vars["message_id"]

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// 读取请求的 JSON Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	code, response, err := service_api_app.MessageFeedbackApi(messageId, body, r.Header)
	api.GeneralPostResponse(w, code, response, err, "Failed to call message feedback API")
	return
}

func MessageSuggestedApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	messageId := vars["message_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := service_api_app.MessageSuggestedApi(messageId, r.Header)
	api.GeneralGetResponse(w, code, response, err, "Failed to call message suggested API")
	return
}
