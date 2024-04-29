package service_api_dataset

import (
	"net/http"
	"wzinc/dify"
)

func DocumentAddByTextApi(datasetId string, body []byte, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/datasets/" + datasetId + "/document/create_by_text"
	return dify.GeneralPostForwardWithHeader(url, body, header)
}

// 本身接口都有问题，不可用，因此，这个实现未能真正完成，回头版本再考虑
func DocumentAddByFileApi(datasetId string, formData map[string]string, files map[string]string, header http.Header) (int, []byte, error) {
	url := dify.WholeServiceApiPrefix + "/datasets/" + datasetId + "/document/create_by_file"
	//return dify.FilePostForward(url, formData, files, header)
	return dify.FilePostForward(url, "", "")
}
