package service_api_dataset

import (
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/service_api_dataset"
)

func DatasetApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// 调用 ListApp 函数获取数据，传递 limit 和 page 参数
		code, response, err := service_api_dataset.GetDatasetApi(r.URL.RawQuery, r.Header)
		api.GeneralGetResponse(w, code, response, err, "Failed to call get dataset API")
		return
	}

	if r.Method == http.MethodPost {
		// 读取请求的 JSON Body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		code, response, err := service_api_dataset.PostDatasetApi(body, r.Header)
		api.GeneralPostResponse(w, code, response, err, "Failed to call post dataset API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}
