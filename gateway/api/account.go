package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wzinc/dify"
)

func CallbackCreateHandler(w http.ResponseWriter, r *http.Request) {
	//host := r.Host
	//fmt.Println("Host:", host)
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	type RequestBody struct {
		Name string `json:"name"`
	}

	// 读取请求体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 解析请求体
	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// 提取 name 字段的值
	name := requestBody.Name

	// 处理 name 字段的值
	fmt.Fprintf(w, "Name: %s\n", name)
	if name == "" {
		http.Error(w, "Name Empty", http.StatusBadRequest)
	}

	// 调用 ListApp 函数获取数据，传递 limit 和 page 参数
	data := dify.CallbackCreate(name)

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

func CallbackDeleteHandler(w http.ResponseWriter, r *http.Request) {
	//host := r.Host
	//fmt.Println("Host:", host)
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	type RequestBody struct {
		Name string `json:"name"`
	}

	// 读取请求体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 解析请求体
	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// 提取 name 字段的值
	name := requestBody.Name

	// 处理 name 字段的值
	fmt.Fprintf(w, "Name: %s\n", name)
	if name == "" {
		http.Error(w, "Name Empty", http.StatusBadRequest)
	}

	// 调用 ListApp 函数获取数据，传递 limit 和 page 参数
	data := dify.CallbackDelete(name)

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
