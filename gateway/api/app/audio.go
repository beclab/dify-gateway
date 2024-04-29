package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/app"
)

func ChatMessageAudioApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		// 处理错误
		return
	}
	defer file.Close()

	filename := fileHeader.Filename

	// 将文件内容读取到字节切片中
	content, err := ioutil.ReadAll(file)
	if err != nil {
		// 处理错误
		return
	}

	// 将文件内容转换为字符串
	contentStr := string(content)

	fmt.Println(filename)
	fmt.Println(contentStr)

	code, response, err := app.ChatMessageAudioApi(appId, filename, contentStr)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app message audio API")
	return
}

func ChatMessageTextApiHandler(w http.ResponseWriter, r *http.Request) {
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

	code, response, err := app.ChatMessageTextApi(appId, body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app message text API")
	return
}

func TextModesApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.TextModesApi(appId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app text modes API")
	return
}
