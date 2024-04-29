package datasets

import "wzinc/dify"

func GetDataSourceApi() (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/data-source/integrates"
	return dify.GeneralGetForward(url)
}

func PostDataSourceApi(bindingId string, action string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/data-source/integrates/" + bindingId + "/" + action
	return dify.GeneralPostForward(url, nil)
}

func DataSourceNotionListApi(rawQuery string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/notion/pre-import/pages"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}

func GetDataSourceNotionApi(workspaceId string, pageId string, pageType string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/notion/workspaces/" + workspaceId + "/pages/" + pageId + "/" + pageType + "/preview"
	return dify.GeneralGetForward(url)
}

func PostDataSourceNotionApi(body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/notion-indexing-estimate"
	return dify.GeneralPostForward(url, body)
}

func DataSourceNotionDatasetSyncApi(datasetId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/notion/sync"
	return dify.GeneralGetForward(url)
}

func DataSourceNotionDocumentSyncApi(datasetId string, documentId string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/datasets/" + datasetId + "/documents/" + documentId + "/notion/sync"
	return dify.GeneralGetForward(url)
}
