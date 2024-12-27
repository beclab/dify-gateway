package dify

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func CallbackCreateV14(name string) map[string]interface{} {
	email := name + "@dify.ai"
	url := DifyHost + "/console/api/workspaces/current/members"
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
		return nil
	}

	fmt.Println(myRespBody)

	accounts, ok := myRespBody["accounts"].([]interface{})
	if !ok {
		fmt.Println("Cannot get accounts field")
		return nil
	}

	for _, item := range accounts {
		itemAccount, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("Cannot get data of Account")
			continue
		}

		itemEmail, ok := itemAccount["email"].(string)
		if !ok {
			fmt.Println("Cannot get Email field")
			continue
		}

		if itemEmail == email {
			//timestamp := time.Now().Unix()
			//timeStr := time.Unix(timestamp, 0).Format("20060102150405")
			//email = name + timeStr + "@dify.ai"
			fmt.Println("User already exists!")
			return nil
		}
	}

	var body struct {
		Emails   []string `json:"emails"`
		Role     string   `json:"role"`
		Language string   `json:"language"`
	}

	body.Emails = []string{email}
	body.Role = "admin"
	body.Language = "zh-Hans"

	statusCode, respBody, _, _ := JSONWithResp(DifyHost+"/console/api/workspaces/current/members/invite-email",
		"POST",
		DifyHeaders,
		body,
		time.Duration(time.Second*10))

	if statusCode != 201 {
		fmt.Println(statusCode, string(respBody))
		fmt.Println("Invite Email Error!")
	}

	type InnerInvitationData struct {
		Status string `json:"status"`
		Email  string `json:"email"`
		Url    string `json:"url"`
	}

	type OuterInvitationData struct {
		Result            string                `json:"result"`
		InvitationResults []InnerInvitationData `json:"invitation_results"`
	}

	var invitationData OuterInvitationData
	err = json.Unmarshal([]byte(respBody), &invitationData)
	if err != nil {
		fmt.Println(err)
	}

	invitationUrl := invitationData.InvitationResults[0].Url
	tempSplit := strings.Split(invitationUrl, "=")
	token := tempSplit[len(tempSplit)-1]

	var bodyLogin struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		Language    string `json:"language"`
		RememberMe  bool   `json:"remember_me"`
		InviteToken string `json:"invite_token"`
	}
	bodyLogin.Email = email
	bodyLogin.Password = os.Getenv("DIFY_USER_PASSWORD")
	bodyLogin.Language = "zh-Hans"
	bodyLogin.RememberMe = true
	bodyLogin.InviteToken = token

	fmt.Println(bodyLogin)

	statusCodeLogin, respBodyLogin, _, _ := JSONWithResp(DifyHost+"/console/api/login",
		"POST",
		DifyHeaders,
		bodyLogin,
		time.Duration(time.Second*10))

	if statusCodeLogin != 200 {
		fmt.Println(statusCodeLogin, string(respBodyLogin))
		fmt.Println(respBodyLogin)
		return nil
	}

	var body2 struct {
		Email             string `json:"email"`
		Token             string `json:"token"`
		Name              string `json:"name"`
		Password          string `json:"password"`
		InterfaceLanguage string `json:"interface_language"`
		Timezone          string `json:"timezone"`
	}

	body2.Email = email
	body2.Token = token
	body2.Name = name
	body2.Password = os.Getenv("DIFY_USER_PASSWORD") //"abcd123456"
	body2.InterfaceLanguage = "zh-Hans"
	body2.Timezone = "Asia/Shanghai"

	fmt.Println(body2)

	statusCode2, respBody2, _, _ := JSONWithResp(DifyHost+"/console/api/activate",
		"POST",
		DifyHeaders,
		body2,
		time.Duration(time.Second*10))

	//if statusCode2 != 200 {
	//	fmt.Println(statusCode2)
	//	fmt.Println(myRespBody2)
	//	return nil
	//}

	var myRespBody2 map[string]interface{}

	err = json.Unmarshal([]byte(respBody2), &myRespBody2)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if statusCode2 != 200 {
		fmt.Println(statusCode2)
		fmt.Println(myRespBody2)
		fmt.Println("Waiting for 5 minutes to retry...")
		time.Sleep(5 * time.Minute)
		fmt.Println("Waiting end！")
		return nil
	}

	return myRespBody2
}

func CallbackCreate(name string) map[string]interface{} {
	email := name + "@dify.ai"
	url := DifyHost + "/console/api/workspaces/current/members"
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
		return nil
	}

	fmt.Println(myRespBody)

	accounts, ok := myRespBody["accounts"].([]interface{})
	if !ok {
		fmt.Println("Cannot get accounts field")
		return nil
	}

	for _, item := range accounts {
		itemAccount, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("Cannot get data of Account")
			continue
		}

		itemEmail, ok := itemAccount["email"].(string)
		if !ok {
			fmt.Println("Cannot get Email field")
			continue
		}

		if itemEmail == email {
			//timestamp := time.Now().Unix()
			//timeStr := time.Unix(timestamp, 0).Format("20060102150405")
			//email = name + timeStr + "@dify.ai"
			fmt.Println("User already exists!")
			return nil
		}
	}

	var body struct {
		Emails string `json:"emails"`
		Role   string `json:"role"`
	}

	body.Emails = email
	body.Role = "admin"

	statusCode, respBody, _, _ := JSONWithResp(DifyHost+"/console/api/workspaces/current/members/invite-email",
		"POST",
		DifyHeaders,
		body,
		time.Duration(time.Second*10))

	if statusCode != 201 {
		fmt.Println(statusCode, string(respBody))
		fmt.Println("Invite Email Error!")
	}

	type InnerInvitationData struct {
		Status string `json:"status"`
		Email  string `json:"email"`
		Url    string `json:"url"`
	}

	type OuterInvitationData struct {
		Result            string                `json:"result"`
		InvitationResults []InnerInvitationData `json:"invitation_results"`
	}

	var invitationData OuterInvitationData
	err = json.Unmarshal([]byte(respBody), &invitationData)
	if err != nil {
		fmt.Println(err)
	}

	invitationUrl := invitationData.InvitationResults[0].Url
	tempSplit := strings.Split(invitationUrl, "=")
	token := tempSplit[len(tempSplit)-1]

	var body2 struct {
		Email             string `json:"email"`
		Token             string `json:"token"`
		Name              string `json:"name"`
		Password          string `json:"password"`
		InterfaceLanguage string `json:"interface_language"`
		Timezone          string `json:"timezone"`
	}

	body2.Email = email
	body2.Token = token
	body2.Name = name
	body2.Password = os.Getenv("DIFY_USER_PASSWORD") //"abcd123456"
	body2.InterfaceLanguage = "zh-Hans"
	body2.Timezone = "Asia/Shanghai"

	fmt.Println(body2)

	statusCode2, respBody2, _, _ := JSONWithResp(DifyHost+"/console/api/activate",
		"POST",
		DifyHeaders,
		body2,
		time.Duration(time.Second*10))

	//if statusCode2 != 200 {
	//	fmt.Println(statusCode2)
	//	fmt.Println(myRespBody2)
	//	return nil
	//}

	var myRespBody2 map[string]interface{}

	err = json.Unmarshal([]byte(respBody2), &myRespBody2)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if statusCode2 != 200 {
		fmt.Println(statusCode2)
		fmt.Println(myRespBody2)
		fmt.Println("Waiting for 5 minutes to retry...")
		time.Sleep(5 * time.Minute)
		fmt.Println("Waiting end！")
		return nil
	}

	return myRespBody2
}

func CallbackDelete(name string) map[string]interface{} {
	email := name + "@dify.ai"
	url := DifyHost + "/console/api/workspaces/current/members"
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
		return nil
	}

	accounts, ok := myRespBody["accounts"].([]interface{})
	if !ok {
		fmt.Println("Cannot get accounts field")
		return nil
	}

	for _, item := range accounts {
		itemAccount, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("Cannot get data of Account")
			continue
		}

		itemEmail, ok := itemAccount["email"].(string)
		if !ok {
			fmt.Println("Cannot get Email field")
			continue
		}

		if itemEmail == email {
			itemId, ok := itemAccount["id"].(string)
			if !ok {
				fmt.Println("Cannot get id field")
				return nil
			}
			statusCode, _, _, _ := JSONWithResp(DifyHost+"/console/api/workspaces/current/members/"+itemId,
				"DELETE",
				DifyHeaders,
				nil,
				time.Duration(time.Second*10))
			if statusCode != 204 {
				return nil
			}
			resultStr := `{"result": "success"}`

			var result map[string]interface{}
			err := json.Unmarshal([]byte(resultStr), &result)
			if err != nil {
				fmt.Println(err)
			}
			return result
		}
	}
	fmt.Println("User does not exist!")
	return nil
}
