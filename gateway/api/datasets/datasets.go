package datasets

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/datasets"
)

func DatasetListApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		code, response, err := datasets.GetDatasetListApi(r.URL.RawQuery)
		fmt.Println(code, response)
		api.GeneralGetResponse(w, code, response, err, "Failed to call list datasets API")
		return
	}

	if r.Method == http.MethodPost {
		// 读取请求的 JSON Body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		code, response, err := datasets.PostDatasetListApi(body)
		api.GeneralPostResponse(w, code, response, err, "Failed to call create datasets API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func DatasetApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]

	if r.Method == http.MethodGet {
		code, response, err := datasets.GetDatasetApi(datasetId)
		api.GeneralGetResponse(w, code, response, err, "Failed to call get dataset API")
		return
	}

	if r.Method == http.MethodPatch {
		// 读取请求的 JSON Body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		code, response, err := datasets.PatchDatasetApi(datasetId, body)
		api.GeneralPatchResponse(w, code, response, err, "Failed to call patch dataset API")
		return
	}

	if r.Method == http.MethodDelete {
		code, response, err := datasets.DeleteDatasetApi(datasetId)
		api.GeneralGetResponse(w, code, response, err, "Failed to call delete dataset API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func DatasetQueryApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DatasetQueryApi(datasetId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call dataset query API")
	return
}

func DatasetIndexingEstimateApiHandler(w http.ResponseWriter, r *http.Request) {
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

	code, response, err := datasets.DatasetIndexingEstimateApi(body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call patch dataset API")
	return
}

func DatasetRelatedAppListApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DatasetRelatedAppListApi(datasetId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call dataset related app list API")
	return
}

func DatasetIndexingStatusApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DatasetIndexingStatusApi(datasetId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call dataset indexing status API")
	return
}

func DatasetApiKeyApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		code, response, err := datasets.GetDatasetApiKeyApi()
		api.GeneralGetResponse(w, code, response, err, "Failed to call get dataset api key API")
		return
	}

	if r.Method == http.MethodPost {
		code, response, err := datasets.PostDatasetApiKeyApi()
		api.GeneralPatchResponse(w, code, response, err, "Failed to call create dataset api key API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func DatasetApiDeleteApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	apiKeyId := vars["api_key_id"]

	if r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DatasetApiDeleteApi(apiKeyId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call delete dataset api key API")
	return
}

func DatasetApiBaseUrlApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DatasetApiBaseUrlApi()
	api.GeneralGetResponse(w, code, response, err, "Failed to call dataset api base url API")
	return
}
