package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"wzinc/api"
	"wzinc/dify"
	"wzinc/dify/app"
)

// 复制响应头
func copyHeaders(dst, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dst.Add(key, value)
		}
	}
}

func CompletionMessageApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

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
		url, code, response, err = app.CompletionMessageApi(appId, body, true)
		if url != "" {
			err = dify.GeneralSSEForward(url, body, w, r)
			if err != nil {
				code = http.StatusInternalServerError
			}
		}
	} else {
		_, code, response, err = app.CompletionMessageApi(appId, body, false)
	}
	api.GeneralPostResponse(w, code, response, err, "Failed to call app completion message API")
	return
}

func CompletionMessageStopApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	taskId := vars["task_id"]

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.CompletionMessageStopApi(appId, taskId)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app completion message stop API")
	return
}

func ChatMessageApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	fmt.Println("appId: ", appId, " chat messages post!")

	if r.Method == http.MethodPost {
		streaming, body, err := dify.GeneralSSEPrepare(w, r)
		if err != nil {
			return
		}

		var url string = ""
		var code int = http.StatusOK
		var response []byte = []byte("")
		if streaming == "streaming" {
			url, code, response, err = app.ChatMessageApi(appId, body, true)
			if url != "" {
				err = dify.GeneralSSEForward(url, body, w, r)
				if err != nil {
					code = http.StatusInternalServerError
				}
			}
		} else {
			_, code, response, err = app.ChatMessageApi(appId, body, false)
			api.GeneralPostResponse(w, code, response, err, "Failed to call app chat message API")
		}
		return
	}

	if r.Method == http.MethodGet {
		code, response, err := app.ChatMessageListApi(appId, r.URL.RawQuery)
		api.GeneralGetResponse(w, code, response, err, "Failed to call app chat message list API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func ChatMessageStopApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	taskId := vars["task_id"]

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.ChatMessageStopApi(appId, taskId)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app chat message stop API")
	return
}
