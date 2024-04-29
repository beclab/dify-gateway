package service_api_app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify"
	"wzinc/dify/service_api_app"
)

func AudioApiHandler(w http.ResponseWriter, r *http.Request) {
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

	code, response, err := service_api_app.AudioApi(filename, contentStr)
	api.GeneralPostResponse(w, code, response, err, "Failed to call message audio API")
	return
}

func TextApiHandler(w http.ResponseWriter, r *http.Request) {
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
		url, code, response, err = service_api_app.TextApi(body, r.Header, true)
		if url != "" {
			err = dify.GeneralSSEForward(url, body, w, r)
			if err != nil {
				code = http.StatusInternalServerError
			}
		}
	} else {
		_, code, response, err = service_api_app.TextApi(body, r.Header, false)
	}
	api.GeneralPostResponse(w, code, response, err, "Failed to call message text API")
	return
}
