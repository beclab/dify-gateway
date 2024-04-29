package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"wzinc/api"
	"wzinc/dify/app"
)

func DailyConversationStatisticHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.DailyConversationStatistic(appId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app daily conversation statistic API")
	return
}

func DailyTerminalsStatisticHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.DailyTerminalsStatistic(appId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app daily terminals statistic API")
	return
}

func DailyTokenCostStatisticHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.DailyTokenCostStatistic(appId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app daily token cost statistic API")
	return
}

func AverageSessionInteractionStatisticHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.AverageSessionInteractionStatistic(appId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app average session interaction statistic API")
	return
}

func UserSatisfactionRateStatisticHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.UserSatisfactionRateStatistic(appId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app user satisfaction rate statistic API")
	return
}

func AverageResponseTimeStatisticHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.AverageResponseTimeStatistic(appId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app average response time statistic API")
	return
}

func TokensPerSecondStatisticHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.TokensPerSecondStatistic(appId, r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app tokens per second statistic API")
	return
}
