package datasets

import (
	"wzinc/dify"
)

func GetDatasetListApi(rawQuery string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func PostDatasetListApi(body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets"
	return dify.GeneralPostForward(url, body)
}

func GetDatasetApi(datasetId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId
	return dify.GeneralGetForward(url)
}

func PatchDatasetApi(datasetId string, body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId
	return dify.GeneralPatchForward(url, body)
}

func DeleteDatasetApi(datasetId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId
	return dify.GeneralDeleteForward(url)
}

func DatasetQueryApi(datasetId string, rawQuery string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/queries"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func DatasetIndexingEstimateApi(body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/indexing-estimate"
	return dify.GeneralPostForward(url, body)
}

func DatasetRelatedAppListApi(datasetId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/related-apps"
	return dify.GeneralGetForward(url)
}

func DatasetIndexingStatusApi(datasetId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/indexing-status"
	return dify.GeneralGetForward(url)
}

func GetDatasetApiKeyApi() (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/api-keys"
	return dify.GeneralGetForward(url)
}

func PostDatasetApiKeyApi() (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/api-keys"
	return dify.GeneralPostForward(url, nil)
}

func DatasetApiDeleteApi(apiKeyId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/api-keys/" + apiKeyId
	return dify.GeneralDeleteForward(url)
}

func DatasetApiBaseUrlApi() (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/api-base-info"
	return dify.GeneralGetForward(url)
}
