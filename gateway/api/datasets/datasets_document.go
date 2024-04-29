package datasets

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/datasets"
)

func GetProcessRuleApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.GetProcessRuleApi(r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call get process rule API")
	return
}

func DatasetDocumentListApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]

	if r.Method == http.MethodGet {
		code, response, err := datasets.GetDatasetDocumentListApi(datasetId, r.URL.RawQuery)
		api.GeneralGetResponse(w, code, response, err, "Failed to call get dataset document list API")
		return
	}

	if r.Method == http.MethodPost {
		// 读取请求的 JSON Body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		code, response, err := datasets.PostDatasetDocumentListApi(datasetId, body)
		api.GeneralPatchResponse(w, code, response, err, "Failed to call post dataset document list API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func DatasetInitApiHandler(w http.ResponseWriter, r *http.Request) {
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

	code, response, err := datasets.DatasetInitApi(body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call dataset init API")
	return
}

func DocumentIndexingEstimateApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentIndexingEstimateApi(datasetId, documentId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call document indexing estimate API")
	return
}

func DocumentBatchIndexingEstimateApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	batch := vars["batch"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentBatchIndexingEstimateApi(datasetId, batch)
	api.GeneralGetResponse(w, code, response, err, "Failed to call document batch indexing estimate API")
	return
}

func DocumentBatchIndexingStatusApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	batch := vars["batch"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentBatchIndexingStatusApi(datasetId, batch)
	api.GeneralGetResponse(w, code, response, err, "Failed to call document batch indexing status API")
	return
}

func DocumentIndexingStatusApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentIndexingStatusApi(datasetId, documentId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call document indexing status API")
	return
}

func DocumentDetailApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentDetailApi(datasetId, documentId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call document detail API")
	return
}

func DocumentProcessingApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]
	action := vars["action"]

	if r.Method != http.MethodPatch {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentProcessingApi(datasetId, documentId, action)
	api.GeneralPatchResponse(w, code, response, err, "Failed to call dataset document processing API")
	return
}

func DocumentDeleteApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

	if r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentDeleteApi(datasetId, documentId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call document delete API")
	return
}

func DocumentMetadataApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// 读取请求的 JSON Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	code, response, err := datasets.DocumentMetadataApi(datasetId, documentId, body)
	api.GeneralPutResponse(w, code, response, err, "Failed to call document metadata API")
	return
}

func DocumentStatusApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]
	action := vars["action"]

	if r.Method != http.MethodPatch {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentStatusApi(datasetId, documentId, action)
	api.GeneralPatchResponse(w, code, response, err, "Failed to call dataset document status API")
	return
}

func DocumentPauseApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

	if r.Method != http.MethodPatch {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentPauseApi(datasetId, documentId)
	api.GeneralPatchResponse(w, code, response, err, "Failed to call dataset document pause API")
	return
}

func DocumentRecoverApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

	if r.Method != http.MethodPatch {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentRecoverApi(datasetId, documentId)
	api.GeneralPatchResponse(w, code, response, err, "Failed to call dataset document recover API")
	return
}

func DocumentLimitApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DocumentLimitApi()
	api.GeneralGetResponse(w, code, response, err, "Failed to call document limit API")
	return
}
