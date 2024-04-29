package app

import (
	"net/http"
	"wzinc/dify"
)

func DailyConversationStatistic(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/statistics/daily-conversations"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func DailyTerminalsStatistic(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/statistics/daily-end-users"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func DailyTokenCostStatistic(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/statistics/token-costs"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func AverageSessionInteractionStatistic(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/statistics/average-session-interactions"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func UserSatisfactionRateStatistic(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/statistics/user-satisfaction-rate"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func AverageResponseTimeStatistic(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/statistics/average-response-time"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func TokensPerSecondStatistic(appId string, rawQuery string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/statistics/tokens-per-second"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}
