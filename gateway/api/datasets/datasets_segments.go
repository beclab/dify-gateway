package datasets

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/datasets"
)

func DatasetDocumentSegmentListApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DatasetDocumentSegmentListApi(datasetId, documentId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call dataset document segment list API")
	return
}

func DatasetDocumentSegmentApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	segmentId := vars["segment_id"]
	action := vars["action"]

	if r.Method != http.MethodPatch {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.DatasetDocumentSegmentApi(datasetId, segmentId, action)
	api.GeneralPatchResponse(w, code, response, err, "Failed to call dataset document segment API")
	return
}

func DatasetDocumentSegmentAddApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

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

	code, response, err := datasets.DatasetDocumentSegmentAddApi(datasetId, documentId, body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call dataset document segment add API")
	return
}

func DatasetDocumentSegmentUpdateApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]
	segmentId := vars["segment_id"]

	if r.Method == http.MethodPatch {
		// 读取请求的 JSON Body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		code, response, err := datasets.PatchDatasetDocumentSegmentUpdateApi(datasetId, documentId, segmentId, body)
		api.GeneralPatchResponse(w, code, response, err, "Failed to call patch dataset document segment update API")
		return
	}

	if r.Method == http.MethodDelete {
		code, response, err := datasets.DeleteDatasetDocumentSegmentUpdateApi(datasetId, documentId, segmentId)
		api.GeneralGetResponse(w, code, response, err, "Failed to call delete dataset document segment update API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func PostDatasetDocumentSegmentBatchImportApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	datasetId := vars["dataset_id"]
	documentId := vars["document_id"]

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

	code, response, err := datasets.PostDatasetDocumentSegmentBatchImportApi(datasetId, documentId, filename, contentStr)
	api.GeneralPostResponse(w, code, response, err, "Failed to call post dataset document segment batch import API")
	return
}

func GetDatasetDocumentSegmentBatchImportApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	jobId := vars["job_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := datasets.GetDatasetDocumentSegmentBatchImportApi(jobId)
	api.GeneralGetResponse(w, code, response, err, "Failed to call get dataset document segment batch import API")
	return
}
