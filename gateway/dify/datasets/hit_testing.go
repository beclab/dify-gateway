package datasets

import "wzinc/dify"

func HitTestingApi(datasetId string, body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/hit-testing"
	return dify.GeneralPostForward(url, body)
}
