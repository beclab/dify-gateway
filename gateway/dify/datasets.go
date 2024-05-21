package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
	"wzinc/parser"
)

func CreateDocument() {
	var body struct {
		Name string `json:"name"`
	}

	body.Name = "Document"

	statusCode, respBody, respHeader, _ := JSONWithResp(DifyHost+"/console/api/datasets",
		"POST",
		DifyHeaders,
		body,
		time.Duration(time.Second*10))

	fmt.Println(statusCode, respHeader, string(respBody))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
	}
	DatasetId = myRespBody["id"].(string)
}

func CreateDocumentV2(bflName string) {
	var body struct {
		Name string `json:"name"`
	}

	body.Name = bflName + "'s Document"

	statusCode, respBody, respHeader, _ := JSONWithResp(DifyHost+"/console/api/datasets",
		"POST",
		DifyHeaders,
		body,
		time.Duration(time.Second*10))

	fmt.Println(statusCode, respHeader, string(respBody))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
	}
	DatasetId = myRespBody["id"].(string)
}

func postFile(filename string, content string, target_url string, headers map[string]string) (*http.Response, error) {

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	//paths := strings.Split(filename, "/")
	//_, err := body_writer.CreateFormFile("file", paths[len(paths)-1])
	_, err := body_writer.CreateFormFile("file", path.Base(filename))
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return nil, err
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
	fi, err := fh.Stat()
	if err != nil {
		fmt.Printf("Error Stating file: %s", filename)
		return nil, err
	}
	req, err := http.NewRequest("POST", target_url, request_reader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	//fmt.Println(req.Header)
	//fmt.Println(req.Body)
	return http.DefaultClient.Do(req)
}

func UploadFile(fileName string, content string) (fileId string) {
	resp, err := postFile( // "/Users/wangrongxiang/Documents/seafile-ce-env-latest镜像使用方法.pdf",
		fileName,
		content,
		DifyHost+"/console/api/files/upload",
		DifyHeaders)
	if err != nil {
		fmt.Println(err)
		return
	}

	statusCode := resp.StatusCode
	respBody, _ := ioutil.ReadAll(resp.Body)
	//respHeader := resp.Header
	//fmt.Println(statusCode, string(respBody), respHeader)

	if statusCode != 201 {
		return ""
	}
	var myRespBody map[string]interface{}
	err = json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
	}
	fileId = myRespBody["id"].(string)
	return
}

func IndexerEstimate(fileId string) {
	// body example:
	//{
	//	"info_list": {
	//		"data_source_type": "upload_file",
	//		"file_info_list": {
	//			"file_ids": ["89e6269a-86f7-41aa-8546-7c574bf14bdd"]
	//		}
	//	},
	//	"indexing_technique": "economy",
	//	"process_rule": {
	//		"rules": {
	//			"pre_processing_rules": [
	//				{"id":"remove_extra_spaces","enabled":true},
	//				{"id":"remove_urls_emails","enabled":true}
	//			],
	//			"segmentation": {
	//				"separator":"\n",
	//				"max_tokens":500
	//			}
	//		},
	//		"mode": "custom"
	//	}
	//}

	var body struct {
		IndexingTechnique string `json:"indexing_technique"`
		InfoList          struct {
			DataSourceType string `json:"data_source_type"`
			FileInfoList   struct {
				FileIds []string `json:"file_ids"`
			} `json:"file_info_list"`
		} `json:"info_list"`
		ProcessRule struct {
			Mode  string   `json:"mode"`
			Rules struct{} `json:"rules"`
		} `json:"process_rule"`
	}

	body.IndexingTechnique = "economy"
	body.InfoList.DataSourceType = "upload_file"
	body.InfoList.FileInfoList.FileIds = []string{fileId}
	body.ProcessRule.Mode = "automatic"

	statusCode, respBody, respHeader, _ := JSONWithResp(DifyHost+"/console/api/datasets/indexing-estimate",
		"POST",
		DifyHeaders,
		body,
		time.Duration(time.Second*10))

	fmt.Println(statusCode, respHeader, string(respBody))
}

func DatasetsUploadDocument(fileId string, targetDatasetID string) {
	//body example:
	//{
	//	"data_source": {
	//		"type": "upload_file",
	//		"info_list": {
	//			"data_source_type": "upload_file",
	//			"file_info_list": {
	//				"file_ids": ["89e6269a-86f7-41aa-8546-7c574bf14bdd"]
	//			}
	//		}
	//	},
	//	"indexing_technique": "economy",
	//	"process_rule": {
	//		"rules": {},
	//		"mode": "automatic"
	//	}
	//}

	var body struct {
		IndexingTechnique string `json:"indexing_technique"`
		DataSource        struct {
			Type     string `json:"type"`
			InfoList struct {
				DataSourceType string `json:"data_source_type"`
				FileInfoList   struct {
					FileIds []string `json:"file_ids"`
				} `json:"file_info_list"`
			} `json:"info_list"`
		} `json:"data_source"`
		ProcessRule struct {
			Mode  string   `json:"mode"`
			Rules struct{} `json:"rules"`
		} `json:"process_rule"`
	}

	body.IndexingTechnique = "economy"
	body.DataSource.Type = "upload_file"
	body.DataSource.InfoList.DataSourceType = "upload_file"
	body.DataSource.InfoList.FileInfoList.FileIds = []string{fileId}
	body.ProcessRule.Mode = "automatic"

	// statusCode, respBody, respHeader, _ :=
	_, _, _, _ = JSONWithResp(DifyHost+"/console/api/datasets/"+targetDatasetID+"/documents",
		"POST",
		DifyHeaders,
		body,
		time.Duration(time.Second*10))

	//fmt.Println(statusCode, respHeader, string(respBody))

	//var myRespBody map[string]interface{}
	//err := json.Unmarshal([]byte(respBody), &myRespBody)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//DatasetId = myRespBody["id"].(string)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func DatasetsSearchDocument(fileName string, targetDatasetID string) (fileIds []string) {
	// statusCode, respBody, respHeader, _ :=
	_, respBody, _, _ := JSONWithResp(DifyHost+"/console/api/datasets/"+targetDatasetID+"/documents",
		"GET",
		DifyHeaders,
		nil,
		time.Duration(time.Second*10))

	//fmt.Println(statusCode, respHeader, string(respBody))

	var myRespBody map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &myRespBody)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(myRespBody["data"])
	datasets := myRespBody["data"].([]interface{})
	//fmt.Println(datasets)
	fileIds = []string{}[:]
	for _, value := range datasets {
		valueTmp := value.(map[string]interface{})
		if valueTmp["name"].(string) == fileName {
			fileIds = append(fileIds, valueTmp["id"].(string))
		}
	}
	return
}

func DatasetsAddDocument(fileName string, targetDatasetID string) error {
	if !IsDir(fileName) && IsFileExist(fileName) {
		fileIds := DatasetsSearchDocument(path.Base(fileName), targetDatasetID)
		if len(fileIds) == 0 {
			content := ""
			fileType := parser.GetTypeFromName(fileName)
			if fileType != ".txt" && fileType != ".md" && fileType != ".markdown" && fileType != ".html" && fileType != ".htm" && fileType != ".docx" && fileType != ".doc" && fileType != ".pdf" && fileType != ".xls" && fileType != ".xlsx" && fileType != ".csv" {
				return nil
				//f, err := os.Open(fileName)
				//if err != nil {
				//	return err
				//}
				//data, _ := ioutil.ReadAll(f)
				//f.Close()
				//r := bytes.NewReader(data)
				//content, err = parser.ParseDoc(r, fileName)
				//if err != nil {
				//	return err
				//}
				//fmt.Println(content)
			}
			uploadFileId := UploadFile(fileName, content)
			if uploadFileId != "" {
				DatasetsUploadDocument(uploadFileId, targetDatasetID)
			}
		}
	}
	return nil
}

func DatasetsDeleteDocument(fileName string, targetDatasetID string) {
	fileIds := DatasetsSearchDocument(path.Base(fileName), targetDatasetID)
	for _, fileId := range fileIds {
		if fileId != "" {
			//statusCode, respBody, respHeader, _ :=
			_, _, _, _ = JSONWithResp(DifyHost+"/console/api/datasets/"+targetDatasetID+"/documents/"+fileId,
				"DELETE",
				DifyHeaders,
				nil,
				time.Duration(time.Second*10))

			//fmt.Println(statusCode, respHeader, string(respBody))
		}
	}
	return
}
