package app

import (
	"net/http"
	"wzinc/dify"
)

func ModelConfigResource(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/model-config"
	return dify.GeneralPostForward(url, body)
}
