package inotify

//import (
//	"os"
//)
//
//var NginxPrefix = os.Getenv("PREFIX")
//
//type RequestProvider struct {
//	Op       string               `json:"op"`
//	DataType string               `json:"datatype"`
//	Version  string               `json:"version"`
//	Group    string               `json:"group"`
//	Token    string               `json:"token"`
//	Data     *RequestProviderData `json:"data"`
//}
//
//type RequestProviderData struct {
//	OpType int                    `json:"op_type"` // 1: GetDatasets; 2: PostDatasets; 3: GetIndexingStatus; 4: UpdateCallback; 5: DeleteDatasets
//	OpData map[string]interface{} `json:"op_data"`
//}

//func DifyGatewayBaseProviderHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("You got the dify gateway base provider~!")
//
//	// 解析请求体
//	var token RequestProvider
//	err := json.NewDecoder(r.Body).Decode(&token)
//	if err != nil {
//		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
//		return
//	}
//
//	fmt.Println("search folder paths")
//	fmt.Println(token)
//	fmt.Println(token.Data)
//
//	opType := token.Data.OpType
//	opData := token.Data.OpData
//	fmt.Println("Accept Request:")
//	fmt.Println("OpType:", opType)
//	fmt.Println("opData:", opData)
//
//	var req *http.Request
//	switch opType {
//	case 1:
//		fmt.Println("OpType: Get Datasets")
//
//		url := "http://127.0.0.1:6317" + NginxPrefix + dify.ConsoleApiPrefix + "/datasets"
//		fmt.Println("Real URL: ", url)
//		req, err = http.NewRequest("GET", url, nil)
//		if err != nil {
//			http.Error(w, "Failed to create new request", http.StatusInternalServerError)
//			return
//		}
//
//		query := req.URL.Query()
//		query.Set("page", "1")
//		query.Set("limit", "10000")
//		req.URL.RawQuery = query.Encode()
//		fmt.Println(req.URL.String())
//	case 2:
//		name := opData["name"].(string)
//
//		fmt.Println("OpType: Create Dataset")
//		fmt.Println("Dataset Name:", name)
//
//		opDataBytes, err := json.Marshal(opData)
//		if err != nil {
//			http.Error(w, "Failed to marshal opData", http.StatusInternalServerError)
//			return
//		}
//
//		url := "http://127.0.0.1:6317" + NginxPrefix + dify.ConsoleApiPrefix + "/datasets"
//		req, err = http.NewRequest("POST", url, bytes.NewBuffer(opDataBytes))
//		if err != nil {
//			http.Error(w, "Failed to create new request", http.StatusInternalServerError)
//			return
//		}
//		fmt.Println(req.URL.String())
//		// 在这里处理opType为2的逻辑，使用opDataStruct.Name变量
//	case 3:
//		datasetID := opData["datasetID"].(string)
//
//		fmt.Println("OpType: Get Dataset Indexing Status")
//		fmt.Println("Dataset ID:", datasetID)
//
//		url := "http://127.0.0.1:6317" + NginxPrefix + dify.ConsoleApiPrefix + "/datasets/" + datasetID + "/indexing-status"
//		req, err = http.NewRequest("GET", url, nil)
//		if err != nil {
//			http.Error(w, "Failed to create new request", http.StatusInternalServerError)
//			return
//		}
//		fmt.Println(req.URL.String())
//		// 在这里处理opType为3的逻辑，使用opDataStruct.DatasetID变量
//	case 4:
//		fmt.Println("OpType: Update Callback")
//		err = GetAndUpdateDatasetFolderStatus()
//		if err != nil {
//			http.Error(w, "Failed to callback update", http.StatusInternalServerError)
//			return
//		}
//		req = nil
//	case 5:
//		datasetID := opData["datasetID"].(string)
//
//		fmt.Println("OpType: Delete Dataset")
//		fmt.Println("Dataset ID:", datasetID)
//
//		//opDataBytes, err := json.Marshal(opData)
//		//if err != nil {
//		//	http.Error(w, "Failed to marshal opData", http.StatusInternalServerError)
//		//	return
//		//}
//
//		url := "http://127.0.0.1:6317" + NginxPrefix + dify.ConsoleApiPrefix + "/datasets/" + datasetID
//		req, err = http.NewRequest("DELETE", url, nil) // bytes.NewBuffer(opDataBytes))
//		if err != nil {
//			http.Error(w, "Failed to delete dataset", http.StatusInternalServerError)
//			return
//		}
//		fmt.Println(req.URL.String())
//	default:
//		http.Error(w, "Invalid opType", http.StatusBadRequest)
//		return
//	}
//
//	if req != nil {
//		client := &http.Client{}
//		resp, err := client.Do(req)
//		if err != nil {
//			http.Error(w, "Failed to send request", http.StatusInternalServerError)
//			return
//		}
//		defer resp.Body.Close()
//
//		fmt.Println("resp: ", resp)
//
//		body, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
//			return
//		}
//
//		fmt.Println("body: ", string(body))
//
//		//api.GeneralGetResponse(w, resp.StatusCode, body, err, "Failed to provider get datasets!")
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(resp.StatusCode)
//
//		// 将响应体的内容解析为map[string]interface{}
//		var data map[string]interface{}
//		err = json.Unmarshal(body, &data)
//		if err != nil {
//			http.Error(w, "Failed to decode response body", http.StatusInternalServerError)
//			return
//		}
//		fmt.Println(data)
//
//		// 将转换后的数据编码并返回给调用方
//		if err := json.NewEncoder(w).Encode(data); err != nil {
//			http.Error(w, "Failed to write response", http.StatusInternalServerError)
//			return
//		}
//	} else {
//		// 设置响应头
//		w.Header().Set("Content-Type", "application/json")
//		// 发送响应
//		response := map[string]interface{}{
//			"message": "Success",
//		}
//		err = json.NewEncoder(w).Encode(response)
//		if err != nil {
//			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
//			return
//		}
//	}
//}
