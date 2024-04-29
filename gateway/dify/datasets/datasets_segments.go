package datasets

import "wzinc/dify"

func DatasetDocumentSegmentListApi(datasetId string, documentId string, rawQuery string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/segments"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func DatasetDocumentSegmentApi(datasetId string, segmentId string, action string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/segments/" + segmentId + "/" + action
	return dify.GeneralPatchForward(url, nil)
}

func DatasetDocumentSegmentAddApi(datasetId string, documentId string, body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/segment"
	return dify.GeneralPostForward(url, body)
}

func PatchDatasetDocumentSegmentUpdateApi(datasetId string, documentId string, segmentId string, body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/segments/" + segmentId
	return dify.GeneralPatchForward(url, body)
}

func DeleteDatasetDocumentSegmentUpdateApi(datasetId string, documentId string, segmentId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/segments/" + segmentId
	return dify.GeneralDeleteForward(url)
}

func PostDatasetDocumentSegmentBatchImportApi(datasetId string, documentId string, filename string, content string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/segments/batch_import"
	return dify.FilePostForward(url, filename, content)
}

func GetDatasetDocumentSegmentBatchImportApi(jobId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/batch_import_status/" + jobId
	return dify.GeneralGetForward(url)
}
