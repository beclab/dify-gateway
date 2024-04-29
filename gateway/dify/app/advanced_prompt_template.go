package app

import (
	"wzinc/dify"
)

func GetAdvancedPromptTemplateList(rawQuery string) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/app/prompt-templates"
	if rawQuery != "" {
		url += "?" + rawQuery
	}
	return dify.GeneralGetForward(url)
}
