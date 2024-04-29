package datasets

import "wzinc/dify"

func GetFileApi() (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/files/upload"
	return dify.GeneralGetForward(url)
}

func PostFileApi(filename string, content string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/files/upload"
	return dify.FilePostForward(url, filename, content)
}

func FilePreviewApi(fileId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/files/" + fileId + "/preview"
	return dify.GeneralGetForward(url)
}
