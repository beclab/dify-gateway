package dify

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
	"wzinc/database"
)
import "net/http"

var DifyHeaders map[string]string = nil
var DatasetId string = ""
var AshiaAgentId string = ""

var CurrentAccountId = ""
var CurrentAccountEmail = ""
var DifyHost = os.Getenv("DIFY_HOST")
var ConsoleApiPrefix = "/console/api"
var ServiceApiPrefix = "/v1"
var WholeConsoleApiPrefix = DifyHost + ConsoleApiPrefix
var WholeServiceApiPrefix = DifyHost + ServiceApiPrefix

func TestHttps() {
	var body struct {
		Query string `json:"query"`
	}

	body.Query = "seafile"

	headers := make(map[string]string)

	statusCode, respBody, _, _ := JSONWithResp("http://127.0.0.1:6317/api/query?index=Files",
		"POST",
		headers,
		body,
		time.Duration(time.Second*10))

	fmt.Println(statusCode, string(respBody))
}

func JSONWithResp(url string, method string, headers map[string]string, data interface{}, timeout time.Duration) (statusCode int, respBody []byte, respHeader map[string][]string, err error) {

	var jsonBytes []byte
	if bytes, ok := data.([]byte); ok {
		jsonBytes = bytes
	} else {
		if jsonBytes, err = json.Marshal(data); err != nil {
			return
		}
	}

	var req *http.Request
	if req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBytes)); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //不验证ca证书，否则卡在这里
	client := &http.Client{Timeout: timeout, Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode
	respBody, _ = ioutil.ReadAll(resp.Body)
	respHeader = resp.Header

	return
}

func FileWithResp(url string, method string, headers map[string]string, filename string, content string, timeout time.Duration) (statusCode int, respBody []byte, respHeader map[string][]string, err error) {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	//paths := strings.Split(filename, "/")
	//_, err := body_writer.CreateFormFile("file", paths[len(paths)-1])
	_, err = body_writer.CreateFormFile("file", path.Base(filename))
	if err != nil {
		fmt.Println("error writing to buffer")
		return
	}

	var fh *os.File
	var fi os.FileInfo

	if content == "" {
		// the file data will be the second part of the body
		fh, err = os.Open(filename)
		if err != nil {
			fmt.Println("error opening file")
			return
		}
	}

	// need to know the boundary to properly close the part myself.
	boundary := body_writer.Boundary()

	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	var request_reader io.Reader
	if content == "" {
		request_reader = io.MultiReader(body_buf, fh, close_buf)
	} else {
		request_reader = io.MultiReader(body_buf, strings.NewReader(content), close_buf)
	}
	if content == "" {
		fi, err = fh.Stat()
		if err != nil {
			fmt.Printf("Error Stating file: %s", filename)
			return
		}
	}
	req, err := http.NewRequest(method, url, request_reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	if content == "" {
		req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())
	} else {
		req.ContentLength = int64(len(content)) + int64(body_buf.Len()) + int64(close_buf.Len())
	}
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	//fmt.Println(req.Header)
	//fmt.Println(req.Body)
	//resp, err := http.DefaultClient.Do(req)
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //不验证ca证书，否则卡在这里
	client := &http.Client{Timeout: timeout, Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode
	respBody, _ = ioutil.ReadAll(resp.Body)
	respHeader = resp.Header
	return
}

func DifySetup() {
	fmt.Println(DifyHost)
	statusCode, respBody, respHeader, _ := JSONWithResp(
		DifyHost+"/console/api/setup",
		"GET",
		nil,
		nil,
		time.Duration(time.Second*10))
	fmt.Println(statusCode, respHeader, string(respBody))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
	}
	if myRespBody["step"] == "finished" {
		fmt.Println("Dify already setup!")
	} else {
		fmt.Println("Dify not setup! Will setup automatically!")

		var body struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Password string `json:"password"`
		}
		body.Email = os.Getenv("DIFY_ADMIN_USER_EMAIL")
		body.Name = "admin"
		body.Password = os.Getenv("DIFY_ADMIN_USER_PASSWORD")

		headers := make(map[string]string)

		statusCode, respBody, respHeader, err := JSONWithResp(DifyHost+"/console/api/setup",
			"POST",
			headers,
			body,
			time.Duration(time.Second*10))
		fmt.Println(statusCode, respHeader, string(respBody))
		if err != nil {
			fmt.Println("Dify setup failed!")
			return
		}
		fmt.Println("Dify setup success!")
	}
}

func GetDifyHeaders() {
	// 如果还未setup，就先用admin账号setup；由于不会重复setup，所以不同用户空间的程序调用此函数没有问题
	DifySetup()

	var body struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		RememberMe bool   `json:"remember_me"`
	}

	body.Email = os.Getenv("DIFY_ADMIN_USER_EMAIL")
	body.Password = os.Getenv("DIFY_ADMIN_USER_PASSWORD")
	body.RememberMe = true

	fmt.Println("body=", body)

	headers := make(map[string]string)

	statusCode, respBody, respHeader, _ := JSONWithResp(DifyHost+"/console/api/login",
		"POST",
		headers,
		body,
		time.Duration(time.Second*10))

	fmt.Println(statusCode, respHeader, string(respBody))

	type LoginData struct {
		Result string `json:"result"`
		Data   string `json:"data"`
	}
	var loginData LoginData
	err := json.Unmarshal([]byte(respBody), &loginData)
	if err != nil {
		fmt.Println(err)
		return
	}
	if statusCode == 200 {
		//remember_token := strings.Split(strings.Split(respHeader["Set-Cookie"][0], ";")[0], "=")[1]
		//session := strings.Split(strings.Split(respHeader["Set-Cookie"][1], ";")[0], "=")[1]
		DifyHeaders = map[string]string{
			//"Cookie":        "remember_token=" + remember_token + "; session=" + session,
			//"Cookie":        "session=" + session,
			"Authorization": "Bearer " + loginData.Data,
		}
		fmt.Println(DifyHeaders)
	} else {
		fmt.Println(statusCode, respHeader, string(respBody))
		fmt.Println("Login failed!")
	}
	return
}

func GetDifyHeadersV2(bflName string) {
	// 如果还未setup，就先用admin账号setup；由于不会重复setup，所以不同用户空间的程序调用此函数没有问题
	DifySetup()

	// 在配置里配置成空间用户名，如果用户已经建立，则什么事也不做，否则，会建立并激活用户（防止自动建立用户延迟）
	CallbackCreate("test")
	CallbackCreate(bflName)

	var body struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		RememberMe bool   `json:"remember_me"`
	}

	body.Email = bflName + "@dify.ai"
	body.Password = os.Getenv("DIFY_USER_PASSWORD")
	body.RememberMe = true

	fmt.Println(body)

	headers := make(map[string]string)

	statusCode, respBody, respHeader, _ := JSONWithResp(DifyHost+"/console/api/login",
		"POST",
		headers,
		body,
		time.Duration(time.Second*10))

	//fmt.Println(statusCode, respHeader, string(respBody))

	type LoginData struct {
		Result string `json:"result"`
		Data   string `json:"data"`
	}
	var loginData LoginData
	err := json.Unmarshal([]byte(respBody), &loginData)
	if err != nil {
		fmt.Println(err)
		return
	}

	if statusCode == 200 {
		//remember_token := strings.Split(strings.Split(respHeader["Set-Cookie"][0], ";")[0], "=")[1]
		//session := strings.Split(strings.Split(respHeader["Set-Cookie"][1], ";")[0], "=")[1]
		DifyHeaders = map[string]string{
			//"Cookie":        "remember_token=" + remember_token + "; session=" + session,
			//"Cookie":        "session=" + session,
			"Authorization": "Bearer " + loginData.Data,
		}
	} else {
		fmt.Println(statusCode, respHeader, string(respBody))
		fmt.Println("Login failed!")
	}
	return
}

func IsDocumentExist() {
	_, respBody, _, _ := JSONWithResp(DifyHost+"/console/api/datasets",
		"GET",
		DifyHeaders,
		nil,
		time.Duration(time.Second*10))
	//fmt.Println(string(respBody))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(myRespBody["data"])
	datasets := myRespBody["data"].([]interface{})
	//fmt.Println(datasets)
	for _, value := range datasets {
		valueTmp := value.(map[string]interface{})
		if valueTmp["name"].(string) == "Document" {
			DatasetId = valueTmp["id"].(string)
		}
	}
	return
}

func IsDocumentExistV2(bflName string) {
	_, respBody, _, _ := JSONWithResp(DifyHost+"/console/api/datasets",
		"GET",
		DifyHeaders,
		nil,
		time.Duration(time.Second*10))
	//fmt.Println(string(respBody))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(myRespBody["data"])
	datasets := myRespBody["data"].([]interface{})
	//fmt.Println(datasets)
	for _, value := range datasets {
		valueTmp := value.(map[string]interface{})
		if valueTmp["name"].(string) == bflName+"'s Document" {
			DatasetId = valueTmp["id"].(string)
		}
	}
	return
}

func IsAgentExist(agentName string) string {
	_, respBody, _, _ := JSONWithResp(DifyHost+"/console/api/apps?page=1&limit=100&name="+agentName,
		"GET",
		DifyHeaders,
		nil,
		time.Duration(time.Second*10))
	fmt.Println(string(respBody))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myRespBody["data"])
	agents := myRespBody["data"].([]interface{})
	fmt.Println(agents)
	for _, value := range agents {
		valueTmp := value.(map[string]interface{})
		if valueTmp["name"].(string) == agentName {
			tempId := valueTmp["id"].(string)
			if IsAppYours(tempId) {
				return tempId
			}
		}
	}
	return ""
}

func PostApps(body []byte) (int, []byte, error) {
	// 调用另一个服务的 createapp 接口
	url := WholeConsoleApiPrefix + "/apps"
	statusCode, respBody, err := GeneralPostForward(url, body)

	var myRespBody map[string]interface{}
	err = json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, respBody, err
	}

	appId, ok := myRespBody["id"].(string)
	if !ok {
		fmt.Println("无法获取id字段")
		return http.StatusInternalServerError, respBody, nil
	}

	// 创建 AccountApp 记录
	accountApp := database.AccountApp{
		AccountID:    CurrentAccountId,
		AccountEmail: CurrentAccountEmail,
		AppID:        appId,
		Status:       database.StatusActive,
	}

	// 创建 AccountApp 记录
	err = database.Database.Create(&accountApp).Error
	if err != nil {
		log.Fatal("Failed to create AccountApp record")
	}
	return statusCode, respBody, nil
}

func CreateAgent(agentName string) {
	// 构建请求体数据
	payload := map[string]interface{}{
		"name":            agentName,
		"icon":            "",
		"icon_background": "#FFEAD5",
		"mode":            "agent-chat",
		"description":     "",
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}
	_, _, err = PostApps(jsonPayload)
	if err != nil {
		fmt.Println(err)
		return
	}
	AshiaAgentId = IsAgentExist(agentName)
}

func GetAppIdsByAgentNameAndModelName(agentName string, modelName string, auth bool) []string {
	_, respBody, _, _ := JSONWithResp(DifyHost+"/console/api/apps?page=1&limit=100&name="+agentName,
		"GET",
		DifyHeaders,
		nil,
		time.Duration(time.Second*10))
	fmt.Println(string(respBody))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
	}

	agents := myRespBody["data"].([]interface{})
	appIds := make([]string, 0)

	for _, value := range agents {
		valueTmp := value.(map[string]interface{})
		if valueTmp["name"].(string) == agentName {
			tempId := valueTmp["id"].(string)
			modelConfig := valueTmp["model_config"].(map[string]interface{})
			model := modelConfig["model"].(map[string]interface{})
			tempModelName := model["name"].(string)
			if auth {
				if !IsAppYours(tempId) {
					continue
				}
			}
			if tempModelName != modelName {
				appIds = append(appIds, tempId)
			}
		}
	}

	return appIds
}

type Config struct {
	//PrePrompt                     string                       `json:"pre_prompt"`
	//PromptType                    string                       `json:"prompt_type"`
	//ChatPromptConfig              map[string]string            `json:"chat_prompt_config"`
	//CompletionPromptConfig        map[string]string            `json:"completion_prompt_config"`
	//UserInputForm                 []string                     `json:"user_input_form"`
	//DatasetQueryVariable          string                       `json:"dataset_query_variable"`
	//OpeningStatement              string                       `json:"opening_statement"`
	//SuggestedQuestions            []string                     `json:"suggested_questions"`
	//MoreLikeThis                  MoreLikeThisConfig           `json:"more_like_this"`
	//SuggestedQuestionsAfterAnswer SuggestedQuestions           `json:"suggested_questions_after_answer"`
	//SpeechToText                  SpeechToTextConfig           `json:"speech_to_text"`
	//TextToSpeech                  TextToSpeechConfig           `json:"text_to_speech"`
	//RetrieverResource             RetrieverResource            `json:"retriever_resource"`
	//SensitiveWordAvoidance        SensitiveWordAvoidanceConfig `json:"sensitive_word_avoidance"`
	//AgentMode                     AgentModeConfig              `json:"agent_mode"`
	Model ModelConfig `json:"model"`
	//DatasetConfigs                DatasetConfigs               `json:"dataset_configs"`
	//FileUpload                    FileUploadConfig             `json:"file_upload"`
}

type MoreLikeThisConfig struct {
	Enabled bool `json:"enabled"`
}

type SuggestedQuestions struct {
	Enabled bool `json:"enabled"`
}

type SpeechToTextConfig struct {
	Enabled bool `json:"enabled"`
}

type TextToSpeechConfig struct {
	Enabled bool `json:"enabled"`
}

type RetrieverResource struct {
	Enabled bool `json:"enabled"`
}

type SensitiveWordAvoidanceConfig struct {
	Enabled bool     `json:"enabled"`
	Type    string   `json:"type"`
	Configs []string `json:"configs"`
}

type AgentModeConfig struct {
	MaxIteration int      `json:"max_iteration"`
	Enabled      bool     `json:"enabled"`
	Strategy     string   `json:"strategy"`
	Tools        []string `json:"tools"`
	Prompt       *string  `json:"prompt"`
}

type ModelConfig struct {
	Provider         string                 `json:"provider"`
	Name             string                 `json:"name"`
	Mode             string                 `json:"mode"`
	CompletionParams CompletionParamsConfig `json:"completion_params"`
}

type CompletionParamsConfig struct {
	Temperature      float64 `json:"temperature"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	MaxTokens        int     `json:"max_tokens"`
}

type DatasetConfigs struct {
	RetrievalModel string `json:"retrieval_model"`
	Datasets       struct {
		Datasets []string `json:"datasets"`
	} `json:"datasets"`
}

type FileUploadConfig struct {
	Image struct {
		Enabled         bool     `json:"enabled"`
		NumberLimits    int      `json:"number_limits"`
		Detail          string   `json:"detail"`
		TransferMethods []string `json:"transfer_methods"`
	} `json:"image"`
}

func MakeSetAgentDefaultModelBody() []byte {
	config := Config{
		Model: ModelConfig{
			Provider:         "openai",
			Name:             "gpt-4",
			Mode:             "chat",
			CompletionParams: CompletionParamsConfig{},
		},
	}
	body, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

	fmt.Println(string(body))
	return body
}

func MakeSetAgentModelBody() []byte {
	config := Config{
		//PrePrompt:                     "",
		//PromptType:                    "simple",
		//ChatPromptConfig:              make(map[string]string),
		//CompletionPromptConfig:        make(map[string]string),
		//UserInputForm:                 []string{},
		//DatasetQueryVariable:          "",
		//OpeningStatement:              "",
		//SuggestedQuestions:            []string{},
		//MoreLikeThis:                  MoreLikeThisConfig{Enabled: false},
		//SuggestedQuestionsAfterAnswer: SuggestedQuestions{Enabled: false},
		//SpeechToText:                  SpeechToTextConfig{Enabled: false},
		//TextToSpeech:                  TextToSpeechConfig{Enabled: false},
		//RetrieverResource:             RetrieverResource{Enabled: false},
		//SensitiveWordAvoidance:        SensitiveWordAvoidanceConfig{Enabled: false, Type: "", Configs: []string{}},
		//AgentMode: AgentModeConfig{
		//	MaxIteration: 5,
		//	Enabled:      true,
		//	Strategy:     "react",
		//	Tools:        []string{},
		//	Prompt:       nil,
		//},
		Model: ModelConfig{
			Provider: "openai_api_compatible",
			Name:     "nitro",
			Mode:     "agent-chat",
			CompletionParams: CompletionParamsConfig{
				Temperature:      0.7,
				TopP:             1,
				FrequencyPenalty: 0,
				PresencePenalty:  0,
				MaxTokens:        512,
			},
		},
		//DatasetConfigs: DatasetConfigs{
		//	RetrievalModel: "single",
		//	Datasets: struct {
		//		Datasets []string `json:"datasets"`
		//	}{
		//		Datasets: []string{},
		//	},
		//},
		//FileUpload: FileUploadConfig{
		//	Image: struct {
		//		Enabled         bool     `json:"enabled"`
		//		NumberLimits    int      `json:"number_limits"`
		//		Detail          string   `json:"detail"`
		//		TransferMethods []string `json:"transfer_methods"`
		//	}{
		//		Enabled:         false,
		//		NumberLimits:    3,
		//		Detail:          "high",
		//		TransferMethods: []string{"remote_url", "local_file"},
		//	},
		//},
	}

	body, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}

	fmt.Println(string(body))
	return body
}

func SetAgentModel(appId string, body []byte) (int, []byte, error) {
	// 调用另一个服务的 createapp 接口
	url := WholeConsoleApiPrefix + "/apps/" + appId + "/model-config"
	statusCode, respBody, err := GeneralPostForward(url, body)

	var myRespBody map[string]interface{}
	err = json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, respBody, err
	}
	return statusCode, respBody, nil
}

func SetAshiaModel() error {
	appIds := GetAppIdsByAgentNameAndModelName("Ashia", "nitro", true)
	fmt.Println("Ashia Agent IDs: ", appIds)
	if len(appIds) > 0 {
		body := MakeSetAgentModelBody()
		//defaultBody := MakeSetAgentDefaultModelBody()
		for _, appId := range appIds {
			fmt.Println("Set Ashia ID: ", appId)
			statusCode, respBody, err := SetAgentModel(appId, body)
			fmt.Println("Set result: ", statusCode, respBody)
			if err != nil {
				return err
			}
			//if statusCode != 200 {
			//	fmt.Println("Set to default Ashia ID: ", appId)
			//	statusCode, respBody, err = SetAgentModel(appId, defaultBody)
			//	fmt.Println("Set result: ", statusCode, respBody)
			//	if err != nil {
			//		return err
			//	}
			//}
		}
	}
	return nil
}

func StartScheduledExecution() {
	// 使用定时器每隔一段时间执行逻辑
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		SetAshiaModel()
	}
}

func GetDifyProfile() {
	url := DifyHost + "/console/api/account/profile"
	_, respBody, _, _ := JSONWithResp(
		url,
		"GET",
		DifyHeaders,
		nil,
		time.Duration(time.Second*10))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	CurrentAccountId = myRespBody["id"].(string)
	CurrentAccountEmail = myRespBody["email"].(string)
	fmt.Println(CurrentAccountId, CurrentAccountEmail)
	return
}

func InitDify() {
	if DifyHeaders == nil {
		GetDifyHeaders()
	}
	IsDocumentExist()
	if DatasetId == "" {
		CreateDocument()
	}
	GetDifyProfile()
}

func InitDifyV2() {
	bflName := os.Getenv("DIFY_USER_NAME")
	if bflName == "" {
		return
	}
	if DifyHeaders == nil {
		GetDifyHeaders()
		fmt.Println("Get Basic Dify Headers done!")
		fmt.Println("Base DifyHeaders:\n", DifyHeaders)
		GetDifyHeadersV2(bflName)
		fmt.Println("Get Specific Dify Headers done!")
		fmt.Println("Specific DifyHeaders:\n", DifyHeaders)
	}
	GetDifyProfile()
	IsDocumentExistV2(bflName)
	if DatasetId == "" {
		fmt.Println("Dataset not found, create a new one.")
		CreateDocumentV2(bflName)
	}
	fmt.Println("Dataset got successfully! Dataset ID: ", DatasetId)
	AshiaAgentId = IsAgentExist("Ashia")
	if AshiaAgentId == "" {
		fmt.Println("Ashia agent not found, create a new one.")
		CreateAgent("Ashia")
	}
	fmt.Println("Ashia agent got successfully! Ashia agent ID: ", AshiaAgentId)
}

func IsAppYours(appId string) bool {
	// 查询符合条件的 AccountApps
	fmt.Println("CurrentAccountID: ", CurrentAccountId)
	fmt.Println("appId: ", appId)
	var accountApps []database.AccountApp
	result := database.Database.Where(
		"account_id = ? AND app_id = ? AND status = ?", CurrentAccountId, appId, database.StatusActive,
	).Find(&accountApps)
	if result.Error != nil {
		panic("Failed to retrieve AccountApps")
	}

	fmt.Println(accountApps)
	if len(accountApps) == 0 {
		return false
	}
	return true
}

func General404Response(code string, message string) []byte {
	if code == "" {
		code = "app_not_found"
	}
	if message == "" {
		message = "App not found."
	}
	type ErrorResponse struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Status  int    `json:"status"`
	}
	errorResponse := ErrorResponse{
		Code:    code,
		Message: message,
		Status:  http.StatusNotFound,
	}
	errorJSON, _ := json.Marshal(errorResponse)
	return errorJSON
}

func GeneralGetForward(url string) (int, []byte, error) {
	statusCode, respBody, _, err := JSONWithResp(
		url,
		"GET",
		DifyHeaders,
		nil,
		time.Duration(time.Second*10))
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, respBody, err
	}
	//fmt.Println(respBody)
	return statusCode, respBody, nil
}

func GeneralGetForwardWithHeader(url string, header http.Header) (int, []byte, error) {
	var h = make(map[string]string)
	for key, value := range header {
		h[key] = value[0]
	}
	for key, value := range DifyHeaders {
		if key != "Authorization" {
			h[key] = value
		}
	}

	statusCode, respBody, _, err := JSONWithResp(
		url,
		"GET",
		h,
		nil,
		time.Duration(time.Second*10))
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, respBody, err
	}
	//fmt.Println(respBody)
	return statusCode, respBody, nil
}

func GeneralSSEPrepare(w http.ResponseWriter, r *http.Request) (string, []byte, error) {
	// 读取请求的 JSON Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return "", nil, err
	}

	// 解析请求体
	var requestBody struct {
		ResponseMode string `json:"response_mode"`
	}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return "", nil, err
	}
	return requestBody.ResponseMode, body, nil
}

//func GeneralSSEForward(targetURL string, requestBody []byte, w http.ResponseWriter, r *http.Request) error {
//	// 创建目标请求
//	targetReq, err := http.NewRequest(r.Method, targetURL, bytes.NewReader(requestBody))
//	if err != nil {
//		return err
//	}
//
//	// 复制请求头
//	targetReq.Header = r.Header.Clone()
//
//	// 发起目标请求
//	client := &http.Client{}
//	resp, err := client.Do(targetReq)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	// 设置响应为服务器发送事件（SSE）
//	w.Header().Set("Content-Type", "text/event-stream")
//	w.Header().Set("Cache-Control", "no-cache")
//	w.Header().Set("Connection", "keep-alive")
//	flusher, _ := w.(http.Flusher)
//	flusher.Flush()
//
//	// 读取目标响应的数据并逐个发送给客户端
//	buf := make([]byte, 4096)
//	for {
//		n, err := resp.Body.Read(buf)
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//			log.Println("Failed to read response body:", err)
//			return err
//		}
//
//		// 发送数据给客户端
//		_, err = w.Write(buf[:n])
//		if err != nil {
//			log.Println("Failed to write response:", err)
//			return err
//		}
//
//		// 刷新数据到客户端
//		flusher.Flush()
//	}
//	return nil
//}

func GeneralSSEForward(targetURL string, requestBody []byte, w http.ResponseWriter, r *http.Request) error {
	// 创建目标请求
	targetReq, err := http.NewRequest(r.Method, targetURL, bytes.NewReader(requestBody))
	if err != nil {
		return err
	}

	// 复制请求头
	targetReq.Header = r.Header.Clone()

	// 添加自定义请求头
	for key, value := range DifyHeaders {
		targetReq.Header.Set(key, value)
	}

	// 发起目标请求
	client := &http.Client{
		Timeout: time.Second * 30, // 设置超时时间为30秒
	}

	// 发起目标请求，并获取响应
	resp, err := client.Do(targetReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("SSE Status Code:", resp.StatusCode)
	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		//w.Header().Set("Content-Type", "application/json")
		// 将目标请求的响应头设置到响应流的头部
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		// 将状态码原样转发
		w.WriteHeader(resp.StatusCode)
		// 将响应体的数据写入响应流
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			return err
		}
		return nil
	}

	// 设置响应为服务器发送事件（SSE）
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return errors.New("response writer does not support flushing")
	}
	flusher.Flush()

	// 将响应体的数据写入响应流
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func GeneralPostForward(url string, body []byte) (int, []byte, error) {
	// 发送 POST 请求
	statusCode, respBody, _, err := JSONWithResp(
		url,
		"POST",
		DifyHeaders,
		body,
		time.Duration(time.Second*10))
	if err != nil {
		return http.StatusInternalServerError, respBody, err
	}
	return statusCode, respBody, nil
}

func GeneralPostForwardWithHeader(url string, body []byte, header http.Header) (int, []byte, error) {
	var h = make(map[string]string)
	for key, value := range header {
		h[key] = value[0]
	}
	for key, value := range DifyHeaders {
		if key != "Authorization" {
			h[key] = value
		}
	}

	// 发送 POST 请求
	statusCode, respBody, _, err := JSONWithResp(
		url,
		"POST",
		h,
		body,
		time.Duration(time.Second*10))
	if err != nil {
		return http.StatusInternalServerError, respBody, err
	}
	return statusCode, respBody, nil
}

func FilePostForward(url string, filename string, content string) (int, []byte, error) {
	// 发送 POST 请求
	statusCode, respBody, _, err := FileWithResp(
		url,
		"POST",
		DifyHeaders,
		filename,
		content,
		time.Duration(time.Second*10))
	if err != nil {
		return http.StatusInternalServerError, respBody, err
	}
	return statusCode, respBody, nil
}

func GeneralPutForward(url string, body []byte) (int, []byte, error) {
	// 发送 POST 请求
	statusCode, respBody, _, err := JSONWithResp(
		url,
		"PUT",
		DifyHeaders,
		body,
		time.Duration(time.Second*10))
	if err != nil {
		return http.StatusInternalServerError, respBody, err
	}
	return statusCode, respBody, nil
}

func GeneralPutForwardWithHeader(url string, body []byte, header http.Header) (int, []byte, error) {
	var h = make(map[string]string)
	for key, value := range header {
		h[key] = value[0]
	}
	for key, value := range DifyHeaders {
		if key != "Authorization" {
			h[key] = value
		}
	}

	// 发送 POST 请求
	statusCode, respBody, _, err := JSONWithResp(
		url,
		"PUT",
		h,
		body,
		time.Duration(time.Second*10))
	if err != nil {
		return http.StatusInternalServerError, respBody, err
	}
	return statusCode, respBody, nil
}

func GeneralDeleteForward(url string) (int, []byte, error) {
	statusCode, respBody, _, err := JSONWithResp(
		url,
		"DELETE",
		DifyHeaders,
		nil,
		time.Duration(time.Second*10))
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, respBody, err
	}
	//fmt.Println(respBody)
	return statusCode, respBody, nil
}

func GeneralDeleteForwardWithHeader(url string, header http.Header) (int, []byte, error) {
	var h = make(map[string]string)
	for key, value := range header {
		h[key] = value[0]
	}
	for key, value := range DifyHeaders {
		if key != "Authorization" {
			h[key] = value
		}
	}

	statusCode, respBody, _, err := JSONWithResp(
		url,
		"DELETE",
		h,
		nil,
		time.Duration(time.Second*10))
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, respBody, err
	}
	//fmt.Println(respBody)
	return statusCode, respBody, nil
}

func GeneralPatchForward(url string, body []byte) (int, []byte, error) {
	// 发送 POST 请求
	statusCode, respBody, _, err := JSONWithResp(
		url,
		"PATCH",
		DifyHeaders,
		body,
		time.Duration(time.Second*10))
	if err != nil {
		return http.StatusInternalServerError, respBody, err
	}
	return statusCode, respBody, nil
}

func GeneralPatchForwardWithHeader(url string, body []byte, header http.Header) (int, []byte, error) {
	var h = make(map[string]string)
	for key, value := range header {
		h[key] = value[0]
	}
	for key, value := range DifyHeaders {
		if key != "Authorization" {
			h[key] = value
		}
	}

	// 发送 POST 请求
	statusCode, respBody, _, err := JSONWithResp(
		url,
		"PATCH",
		h,
		body,
		time.Duration(time.Second*10))
	if err != nil {
		return http.StatusInternalServerError, respBody, err
	}
	return statusCode, respBody, nil
}
