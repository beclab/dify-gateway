package app

import (
	"net/http"
	"wzinc/dify"
)

func AppSite(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/site"
	return dify.GeneralPostForward(url, body)
}

func AppSiteAccessTokenReset(appId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/site/access-token-reset"
	return dify.GeneralPostForward(url, nil)
}
