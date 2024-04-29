package service_api_app

import (
	"github.com/gorilla/mux"
	"net/http"
	"wzinc/api"
	"wzinc/dify"
	"wzinc/dify/service_api_app"
)

func CompletionApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	streaming, body, err := dify.GeneralSSEPrepare(w, r)
	if err != nil {
		return
	}

	var url string = ""
	var code int = http.StatusOK
	var response []byte = []byte("")
	if streaming == "streaming" {
		url, code, response, err = service_api_app.CompletionApi(body, r.Header, true)
		if url != "" {
			err = dify.GeneralSSEForward(url, body, w, r)
			if err != nil {
				code = http.StatusInternalServerError
			}
		}
	} else {
		_, code, response, err = service_api_app.CompletionApi(body, r.Header, false)
	}
	api.GeneralPostResponse(w, code, response, err, "Failed to call completion message API")
	return
}

func CompletionStopApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	taskId := vars["task_id"]

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := service_api_app.CompletionStopApi(taskId, r.Header)
	api.GeneralPostResponse(w, code, response, err, "Failed to call completion message stop API")
	return
}

func ChatApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		streaming, body, err := dify.GeneralSSEPrepare(w, r)
		if err != nil {
			return
		}

		var url string = ""
		var code int = http.StatusOK
		var response []byte = []byte("")
		if streaming == "streaming" {
			url, code, response, err = service_api_app.ChatApi(body, r.Header, true)
			if url != "" {
				err = dify.GeneralSSEForward(url, body, w, r)
				if err != nil {
					code = http.StatusInternalServerError
				}
			}
		} else {
			_, code, response, err = service_api_app.ChatApi(body, r.Header, false)
		}
		api.GeneralPostResponse(w, code, response, err, "Failed to call chat message API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func ChatStopApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	taskId := vars["task_id"]

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := service_api_app.ChatStopApi(taskId, r.Header)
	api.GeneralPostResponse(w, code, response, err, "Failed to call chat message stop API")
	return
}
