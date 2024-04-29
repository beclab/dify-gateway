package app

import (
	"net/http"
	"wzinc/api"
	"wzinc/dify/app"
)

func AdvancedPromptTemplateListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	code, response, err := app.GetAdvancedPromptTemplateList(r.URL.RawQuery)
	api.GeneralGetResponse(w, code, response, err, "Failed to call app advanced prompt template list API")
	return
}
