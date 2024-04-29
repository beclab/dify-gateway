package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"wzinc/database"
	"wzinc/dify"
)

func ListApp(host string, pageInt int, limitInt int) map[string]interface{} {
	url := fmt.Sprintf("%s/apps?page=%d&limit=%d", dify.WholeConsoleApiPrefix, pageInt, limitInt)
	_, respBody, _, _ := dify.JSONWithResp(
		url,
		"GET",
		dify.DifyHeaders,
		nil,
		time.Duration(time.Second*10))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data, ok := myRespBody["data"].([]interface{})
	if !ok {
		fmt.Println("无法获取data字段")
		return nil
	}

	// 提取所有的id值
	for i, item := range data {
		itemData, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("无法获取item的数据")
			continue
		}

		id, ok := itemData["id"].(string)
		if !ok {
			fmt.Println("无法获取id字段")
			continue
		}

		_, subRespBody, _, _ := dify.JSONWithResp(dify.DifyHost+"/console/api/apps/"+id,
			"GET",
			dify.DifyHeaders,
			nil,
			time.Duration(time.Second*10))

		var mySubRespBody map[string]interface{}
		err := json.Unmarshal([]byte(subRespBody), &mySubRespBody)
		if err != nil {
			fmt.Println(err)
			continue
		}

		appBaseURL, _ := mySubRespBody["site"].(map[string]interface{})["app_base_url"].(string)
		mode, _ := mySubRespBody["mode"].(string)
		code, _ := mySubRespBody["site"].(map[string]interface{})["code"].(string)

		appURL := appBaseURL + "/" + mode + "/" + code
		apiBaseURL, _ := mySubRespBody["api_base_url"].(string)

		if !strings.Contains(host, "localhost") {
			appURL = strings.Replace(appURL, "http://localhost", "https://"+host, -1)
			apiBaseURL = strings.Replace(apiBaseURL, "http://localhost", "https://"+host, -1)
		}
		//fmt.Println("app_base_url:", appBaseURL)
		//fmt.Println("mode:", mode)
		//fmt.Println("code:", code)
		//fmt.Println("app_url:", appURL)
		//fmt.Println("api_base_url:", apiBaseURL)

		// 将appURL和apiBaseURL添加到原始的myRespBody["data"]的对应id位置
		itemData["app_url"] = appURL
		itemData["api_base_url"] = apiBaseURL

		data[i] = itemData
	}

	myRespBody["data"] = data

	// 输出更新后的myRespBody
	output, err := json.Marshal(myRespBody)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(string(output))

	// 返回更新后的myRespBody
	return myRespBody
}

func GetApps(host string, page int, limit int) (int, []byte, error) {
	// 查询符合条件的 AccountApps
	var accountApps []database.AccountApp
	offset := (page - 1) * limit
	result := database.Database.Where(
		"account_id = ? AND status = ?", dify.CurrentAccountId, database.StatusActive,
	).Offset(offset).Limit(limit).Find(&accountApps)
	if result.Error != nil {
		panic("Failed to retrieve AccountApps")
	}

	// 判断是否还有后续数据
	var totalCount int64
	countResult := database.Database.Model(&database.AccountApp{}).Where(
		"account_id = ? AND status = ?", dify.CurrentAccountId, database.StatusActive).Count(&totalCount)
	if countResult.Error != nil {
		panic("Failed to retrieve total count")
	}

	hasMore := totalCount > int64(page*limit)

	// 打印分页结果和是否还有后续数据
	for _, app := range accountApps {
		fmt.Printf("AppID: %d\n", app.AppID)
	}
	fmt.Printf("Has more data: %v\n", hasMore)

	// 构建 appIds 数组
	var appIds []string
	for _, app := range accountApps {
		appIds = append(appIds, app.AppID)
	}
	appIdsString := strings.Join(appIds, ", ")

	// 调用另一个服务的接口
	difyHasMore := true
	difyPage := 1
	difyLimit := 100
	var statusCode int
	var respBody []byte
	var err error
	var difyData []interface{}
	var myRespBody map[string]interface{}
	for difyHasMore {
		url := fmt.Sprintf("%s/apps?page=%d&limit=%d", dify.WholeConsoleApiPrefix, difyPage, difyLimit)
		statusCode, respBody, err = dify.GeneralGetForward(url)
		if err != nil {
			return statusCode, respBody, err
		}

		err = json.Unmarshal([]byte(respBody), &myRespBody)
		if err != nil {
			fmt.Println(err)
			return http.StatusInternalServerError, nil, err
		}

		data, ok := myRespBody["data"].([]interface{})
		if !ok {
			fmt.Println("无法获取data字段")
			return http.StatusInternalServerError, nil, nil
		}

		difyHasMore = myRespBody["has_more"].(bool)
		difyPage += 1

		// 提取所有的id值
		for _, item := range data {
			itemData, subOk := item.(map[string]interface{})
			if !subOk {
				fmt.Println("无法获取item的数据")
				continue
			}

			id, subOk := itemData["id"].(string)
			if !subOk {
				fmt.Println("无法获取id字段")
				continue
			}

			if strings.Contains(appIdsString, id) {
				subUrl := dify.WholeConsoleApiPrefix + "/apps/" + id
				_, subRespBody, subErr := dify.GeneralGetForward(subUrl)
				if subErr != nil {
					fmt.Println(subErr)
					continue
				}

				var mySubRespBody map[string]interface{}
				subErr = json.Unmarshal([]byte(subRespBody), &mySubRespBody)
				if subErr != nil {
					fmt.Println(subErr)
					continue
				}

				appBaseURL, _ := mySubRespBody["site"].(map[string]interface{})["app_base_url"].(string)
				mode, _ := mySubRespBody["mode"].(string)
				code, _ := mySubRespBody["site"].(map[string]interface{})["code"].(string)

				appURL := appBaseURL + "/" + mode + "/" + code
				apiBaseURL, _ := mySubRespBody["api_base_url"].(string)

				if !strings.Contains(host, "localhost") {
					appURL = strings.Replace(appURL, "http://localhost", "https://"+host, -1)
					apiBaseURL = strings.Replace(apiBaseURL, "http://localhost", "https://"+host, -1)
				}
				//fmt.Println("app_base_url:", appBaseURL)
				//fmt.Println("mode:", mode)
				//fmt.Println("code:", code)
				//fmt.Println("app_url:", appURL)
				//fmt.Println("api_base_url:", apiBaseURL)

				// 将appURL和apiBaseURL添加到原始的myRespBody["data"]的对应id位置
				itemData["app_url"] = appURL
				itemData["api_base_url"] = apiBaseURL
				// data[i] = itemData

				difyData = append(difyData, itemData)
			}
		}
	}

	myRespBody["page"] = page
	myRespBody["limit"] = limit
	myRespBody["total"] = len(appIds)
	myRespBody["has_more"] = hasMore
	myRespBody["data"] = difyData

	// 输出更新后的myRespBody
	output, err := json.Marshal(myRespBody)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, nil, err
	}

	fmt.Println(string(output))

	// 返回更新后的myRespBody
	return statusCode, output, nil
}

func PostApps(body []byte) (int, []byte, error) {
	return dify.PostApps(body)
	//// 调用另一个服务的 createapp 接口
	//url := dify.WholeConsoleApiPrefix + "/apps"
	//statusCode, respBody, err := dify.GeneralPostForward(url, body)
	//
	//var myRespBody map[string]interface{}
	//err = json.Unmarshal([]byte(respBody), &myRespBody)
	//if err != nil {
	//	fmt.Println(err)
	//	return http.StatusInternalServerError, respBody, err
	//}
	//
	//appId, ok := myRespBody["id"].(string)
	//if !ok {
	//	fmt.Println("无法获取id字段")
	//	return http.StatusInternalServerError, respBody, nil
	//}
	//
	//// 创建 AccountApp 记录
	//accountApp := database.AccountApp{
	//	AccountID:    dify.CurrentAccountId,
	//	AccountEmail: dify.CurrentAccountEmail,
	//	AppID:        appId,
	//	Status:       database.StatusActive,
	//}
	//
	//// 创建 AccountApp 记录
	//err = database.Database.Create(&accountApp).Error
	//if err != nil {
	//	log.Fatal("Failed to create AccountApp record")
	//}
	//return statusCode, respBody, nil
}

func GetAppTemplate() (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/app-templates"
	return dify.GeneralGetForward(url)
}

func GetAppDetail(appId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}
	url := dify.WholeConsoleApiPrefix + "/apps/" + appId
	return dify.GeneralGetForward(url)
}

func DeleteApp(appId string) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId
	statusCode, respBody, err := dify.GeneralDeleteForward(url)

	result := database.Database.Model(&database.AccountApp{}).Where(
		"account_id = ? AND app_id = ? AND status = ?", dify.CurrentAccountId, appId, database.StatusActive,
	).Update("status", database.StatusDelete)
	if result.Error != nil {
		log.Fatal("Failed to delete AccountApps relation!")
	}
	return statusCode, respBody, err
}

func AppCopy(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/copy"
	statusCode, respBody, err := dify.GeneralPostForward(url, body)

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
		AccountID:    dify.CurrentAccountId,
		AccountEmail: dify.CurrentAccountEmail,
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

func AppName(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/name"
	return dify.GeneralPostForward(url, body)
}

func AppIcon(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/icon"
	return dify.GeneralPostForward(url, body)
}

func AppSiteStatus(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/site-status"
	return dify.GeneralPostForward(url, body)
}

func AppApiStatus(appId string, body []byte) (int, []byte, error) {
	if !dify.IsAppYours(appId) {
		return http.StatusNotFound, dify.General404Response("", ""), nil
	}

	url := dify.WholeConsoleApiPrefix + "/apps/" + appId + "/api-status"
	return dify.GeneralPostForward(url, body)
}
