package datasets

import "wzinc/dify"

func GetProcessRuleApi(rawQuery string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/process-rule"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func GetDatasetDocumentListApi(datasetId string, rawQuery string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func PostDatasetDocumentListApi(datasetId string, body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents"
	return dify.GeneralPostForward(url, body)
}

func DatasetInitApi(body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/init"
	return dify.GeneralPostForward(url, body)
}

func DocumentIndexingEstimateApi(datasetId string, documentId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/indexing-estimate"
	return dify.GeneralGetForward(url)
}

func DocumentBatchIndexingEstimateApi(datasetId string, batch string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/batch/" + batch + "/indexing-estimate"
	return dify.GeneralGetForward(url)
}

func DocumentBatchIndexingStatusApi(datasetId string, batch string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/batch/" + batch + "/indexing-status"
	return dify.GeneralGetForward(url)
}

func DocumentIndexingStatusApi(datasetId string, documentId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/indexing-status"
	return dify.GeneralGetForward(url)
}

func DocumentDetailApi(datasetId string, documentId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId
	return dify.GeneralGetForward(url)
}

func DocumentProcessingApi(datasetId string, documentId string, action string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/processing/" + action
	return dify.GeneralPatchForward(url, nil)
}

func DocumentDeleteApi(datasetId string, documentId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId
	return dify.GeneralDeleteForward(url)
}

func DocumentMetadataApi(datasetId string, documentId string, body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/metadata"
	return dify.GeneralPutForward(url, body)
}

func DocumentStatusApi(datasetId string, documentId string, action string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/status/" + action
	return dify.GeneralPatchForward(url, nil)
}

func DocumentPauseApi(datasetId string, documentId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/processing/pause"
	return dify.GeneralPatchForward(url, nil)
}

func DocumentRecoverApi(datasetId string, documentId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/processing/resume"
	return dify.GeneralPatchForward(url, nil)
}

func DocumentLimitApi() (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/limit"
	return dify.GeneralGetForward(url)
}
