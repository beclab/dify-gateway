package service_api_app

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/service_api_app"
)

func ConversationRenameApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	cId := vars["c_id"]

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

	code, response, err := service_api_app.ConversationRenameApi(cId, body, r.Header)
	api.GeneralPostResponse(w, code, response, err, "Failed to call conversation rename API")
	return
}

func ConversationApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := service_api_app.ConversationApi(r.URL.RawQuery, r.Header)
	api.GeneralGetResponse(w, code, response, err, "Failed to call conversation API")
	return
}

func ConversationDetailApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	cId := vars["c_id"]

	if r.Method == http.MethodGet {
		code, response, err := service_api_app.GetConversationDetailApi(cId, r.Header)
		api.GeneralGetResponse(w, code, response, err, "Failed to call get conversation detail API")
		return
	}

	if r.Method == http.MethodDelete {
		code, response, err := service_api_app.DeleteConversationDetailApi(cId, r.Header)
		api.GeneralGetResponse(w, code, response, err, "Failed to call delete conversation detail API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}
