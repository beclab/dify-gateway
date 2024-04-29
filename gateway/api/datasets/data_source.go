package datasets

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/datasets"
)

func GetDataSourceApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.GetDataSourceApi()
	api.GeneralGetResponse(w, code, response, err, "Failed to call get data source API")
	return
}

func PostDataSourceApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	bindingId := vars["binding_id"]
	action := vars["action"]

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.PostDataSourceApi(bindingId, action)
	api.GeneralPostResponse(w, code, response, err, "Failed to call get data source API")
	return
}

func DataSourceNotionListApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DataSourceNotionListApi(r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call data source notion list API")
	return
}

func GetDataSourceNotionApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	workspaceId := vars["workspace_id"]
	pageId := vars["page_id"]
	pageType := vars["page_type"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.GetDataSourceNotionApi(workspaceId, pageId, pageType)
	api.GeneralGetResponse(w, code, response, err, "Failed to call get data source notion API")
	return
}

func PostDataSourceNotionApiHandler(w http.ResponseWriter, r *http.Request) {
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

	code, response, err := datasets.PostDataSourceNotionApi(body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call post data source notion API")
	return
}

func DataSourceNotionDatasetSyncApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DataSourceNotionDatasetSyncApi(datasetId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call data source notion dataset sync API")
	return
}

func DataSourceNotionDocumentSyncApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DataSourceNotionDocumentSyncApi(datasetId, documentId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call data source notion document sync API")
	return
}
