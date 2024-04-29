package service_api_dataset

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/service_api_dataset"
)

func DocumentAddByTextApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]

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

	code, response, err := service_api_dataset.DocumentAddByTextApi(datasetId, body, r.Header)
	api.GeneralPostResponse(w, code, response, err, "Failed to call document add by text API")
	return

}

// 本身接口都有问题，不可用，因此，这个实现未能真正完成，回头版本再考虑
func DocumentAddByFileApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20) // 32MB
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	formData := make(map[string]string)
	for key, values := range r.Form {
		formData[key] = values[0]
	}

	files := make(map[string]string)
	for key := range r.MultipartForm.File {
		fileHeaders := r.MultipartForm.File[key]
		if len(fileHeaders) > 0 {
			file := fileHeaders[0]
			files[key] = file.Filename
		}
	}
	fmt.Println(formData)
	fmt.Println(files)

	code, response, err := service_api_dataset.DocumentAddByFileApi(datasetId, formData, files, r.Header)
	fmt.Println(code, response, err)
	api.GeneralPostResponse(w, code, response, err, "Failed to call document add by text API")
	return

}
