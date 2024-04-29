package app

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify"
	"wzinc/dify/app"
)

func MessageMoreLikeThisApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	messageId := vars["message_id"]

	if r.Method != http.MethodGet {
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
		url, code, response, err = app.MessageMoreLikeThisApi(appId, messageId, r.URL.RawQuery, false)
		if url != "" {
			err = dify.GeneralSSEForward(url, body, w, r)
			if err != nil {
				code = http.StatusInternalServerError
			}
		}
	} else {
		_, code, response, err = app.MessageMoreLikeThisApi(appId, messageId, r.URL.RawQuery, false)
	}
	api.GeneralGetResponse(w, code, response, err, "Failed to call app message more like this API")
	return
}

func MessageSuggestedQuestionApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	messageId := vars["message_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.MessageSuggestedQuestionApi(appId, messageId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app message suggested question API")
	return
}

func MessageFeedbackApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

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

	code, response, err := app.MessageFeedbackApi(appId, body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app message feedback API")
	return
}

func MessageAnnotationApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

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

	code, response, err := app.MessageAnnotationApi(appId, body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app message annotation API")
	return
}

func MessageAnnotationCountApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.MessageAnnotationCountApi(appId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app message annotation count API")
	return
}

func MessageApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	messageId := vars["message_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.MessageApi(appId, messageId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app message API")
	return
}
