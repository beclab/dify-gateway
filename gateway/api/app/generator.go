package app

import (
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/app"
)

func RuleGenerateApiHandler(w http.ResponseWriter, r *http.Request) {
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

	// 调用另一个服务的 createapp 接口
	code, response, err := app.RuleGenerateApi(body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call rule generation API")
	return
}
