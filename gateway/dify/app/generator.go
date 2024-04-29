package app

import (
	"wzinc/dify"
)

func RuleGenerateApi(body []byte) (int, []byte, error) {
	url := dify.WholeConsoleApiPrefix + "/rule-generate"
	return dify.GeneralPostForward(url, body)
}
