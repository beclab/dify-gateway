package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"wzinc/api"
	"wzinc/dify/app"
)

func ListAppHandler(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	fmt.Println("Host:", host)

	page, limit, err := api.GeneralGetPageLimit(r, 1, 100)
	if err != nil {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	// 调用 ListApp 函数获取数据，传递 limit 和 page 参数
	data := app.ListApp(host, page, limit)

	// 将数据转换为 JSON 格式
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 返回响应
	fmt.Fprint(w, string(response))
}

func AppListApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		host := r.Host
		fmt.Println("Host:", host)

		page, limit, err := api.GeneralGetPageLimit(r, 1, 20)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}
		fmt.Println(page, limit)

		// 调用 ListApp 函数获取数据，传递 limit 和 page 参数
		code, response, err := app.GetApps(host, page, limit)
		api.GeneralGetResponse(w, code, response, err, "Failed to call list app API")
		return
	}

	if r.Method == http.MethodPost {
		// 读取请求的 JSON Body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		code, response, err := app.PostApps(body)
		api.GeneralPostResponse(w, code, response, err, "Failed to call create app API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func AppTemplateApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.GetAppTemplate()
	api.GeneralGetResponse(w, code, response, err, "Failed to call app template API")
	return
}

func AppApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]

	if r.Method == http.MethodGet {
		code, response, err := app.GetAppDetail(appId)
		api.GeneralGetResponse(w, code, response, err, "Failed to call get app detail API")
		return
	}

	if r.Method == http.MethodDelete {
		code, response, err := app.DeleteApp(appId)
		api.GeneralDeleteResponse(w, code, response, err, "Failed to call delete app API")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

func AppCopyHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	fmt.Println(appId)

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
	code, response, err := app.AppCopy(appId, body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call copy app API")
	return
}

func AppNameApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	fmt.Println(appId)

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
	code, response, err := app.AppName(appId, body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app name API")
	return
}

func AppIconApiHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	fmt.Println(appId)

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
	code, response, err := app.AppIcon(appId, body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app icon API")
	return
}

func AppSiteStatusHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	fmt.Println(appId)

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
	code, response, err := app.AppSiteStatus(appId, body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app site-status API")
	return
}

func AppApiStatusHandler(w http.ResponseWriter, r *http.Request) {
	// 获取可变内容参数
	vars := mux.Vars(r)
	appId := vars["app_id"]
	fmt.Println(appId)

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
	code, response, err := app.AppApiStatus(appId, body)
	api.GeneralPostResponse(w, code, response, err, "Failed to call app api-status API")
	return
}
