package datasets

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/datasets"
)

func FileApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		code, response, err := datasets.GetFileApi()
		api.GeneralGetResponse(w, code, response, err, "Failed to call get file API")
		return
	}

	if r.Method == http.MethodPost {
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

		code, response, err := datasets.PostFileApi(filename, contentStr)
		api.GeneralPostResponse(w, code, response, err, "Failed to call create file API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func FilePreviewApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	fileId := vars["file_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.FilePreviewApi(fileId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call file preview API")
	return
}
