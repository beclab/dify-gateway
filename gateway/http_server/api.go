package http_server

import (
	"github.com/gorilla/mux"
	"os"
	"wzinc/api"
	"wzinc/api/app"
	"wzinc/api/datasets"
	"wzinc/api/service_api_app"
	"wzinc/api/service_api_dataset"
)

var NginxPrefix = os.Getenv("PREFIX")
var ServiceApiPrefix = NginxPrefix + "/v1"
var ConsoleApiPrefix = NginxPrefix + "/console/api"

// func NewRouter() *http.ServeMux {
func NewRouter() *mux.Router {
	//router := http.NewServeMux()
	router := mux.NewRouter()

	// 添加路由处理程序
	// account
	router.HandleFunc("/callback/create", api.CallbackCreateHandler)
	router.HandleFunc("/callback/delete", api.CallbackDeleteHandler)

	// provider
	//router.HandleFunc(NginxPrefix+"/dify_gateway_base_provider", inotify.DifyGatewayBaseProviderHandler)
	//router.HandleFunc(NginxPrefix+"/update_dataset_folder_paths", inotify.UpdateDatasetFolderPathsHandler) // unused for now

	// service api
	// app
	router.HandleFunc(ServiceApiPrefix+"/parameters", service_api_app.AppParameterApiHandler)

	router.HandleFunc(ServiceApiPrefix+"/audio-to-text", service_api_app.AudioApiHandler)
	router.HandleFunc(ServiceApiPrefix+"/text-to-audio", service_api_app.TextApiHandler)

	router.HandleFunc(ServiceApiPrefix+"/completion-messages", service_api_app.CompletionApiHandler)
	router.HandleFunc(ServiceApiPrefix+"/completion-messages/{task_id}/stop", service_api_app.CompletionStopApiHandler)
	router.HandleFunc(ServiceApiPrefix+"/chat-messages", service_api_app.ChatApiHandler)
	router.HandleFunc(ServiceApiPrefix+"/chat-messages/{task_id}/stop", service_api_app.ChatStopApiHandler)

	router.HandleFunc(ServiceApiPrefix+"/conversations/{c_id}/name", service_api_app.ConversationRenameApiHandler)
	router.HandleFunc(ServiceApiPrefix+"/conversations", service_api_app.ConversationApiHandler)
	router.HandleFunc(ServiceApiPrefix+"/conversations/{c_id}", service_api_app.ConversationDetailApiHandler)

	router.HandleFunc(ServiceApiPrefix+"/messages", service_api_app.MessageListApiHandler)
	router.HandleFunc(ServiceApiPrefix+"/messages/{message_id}/feedbacks", service_api_app.MessageFeedbackApiHandler)
	router.HandleFunc(ServiceApiPrefix+"/messages/{message_id}/suggested", service_api_app.MessageSuggestedApiHandler)

	// dataset
	router.HandleFunc(ServiceApiPrefix+"/datasets", service_api_dataset.DatasetApiHandler)

	router.HandleFunc(ServiceApiPrefix+"/datasets/{dataset_id}/document/create_by_text",
		service_api_dataset.DocumentAddByTextApiHandler)
	router.HandleFunc(ServiceApiPrefix+"/datasets/{dataset_id}/document/create_by_file",
		service_api_dataset.DocumentAddByFileApiHandler) // 本身接口都有问题，不可用，因此，这个实现未能真正完成，回头版本再考虑

	// console api
	// app
	router.HandleFunc("/console/api/app/prompt-templates", app.AdvancedPromptTemplateListHandler)

	router.HandleFunc(NginxPrefix+"/listapp", app.ListAppHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps", app.AppListApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/app-templates", app.AppTemplateApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}", app.AppApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/copy", app.AppCopyHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/name", app.AppNameApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/icon", app.AppIconApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/site-enable", app.AppSiteStatusHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/api-enable", app.AppApiStatusHandler)

	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/audio-to-text", app.ChatMessageAudioApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/text-to-audio", app.ChatMessageTextApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/text-to-audio/voices", app.TextModesApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/completion-messages", app.CompletionMessageApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/completion-messages/{task_id}/stop",
		app.CompletionMessageStopApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/chat-messages", app.ChatMessageApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/chat-messages/{task_id}/stop", app.ChatMessageStopApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/completion-conversations", app.CompletionConversationApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/completion-conversations/{conversation_id}",
		app.CompletionConversationDetailApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/chat-conversations", app.ChatConversationApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/chat-conversations/{conversation_id}",
		app.ChatConversationDetailApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/rule-generate", app.RuleGenerateApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/completion-messages/{message_id}/more-like-this",
		app.MessageMoreLikeThisApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/chat-messages/{message_id}/suggested-questions",
		app.MessageSuggestedQuestionApiHandler)
	//router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/chat-messages", app.ChatMessageListApiHandler)	//URL与前面相同
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/feedbacks", app.MessageFeedbackApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/annotations", app.MessageAnnotationApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/annotations/count", app.MessageAnnotationCountApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/messages/{message_id}", app.MessageApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/model-config", app.ModelConfigResourceHandler)

	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/site", app.AppSiteHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/site/access-token-reset", app.AppSiteAccessTokenResetHandler)

	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/statistics/daily-conversations",
		app.DailyConversationStatisticHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/statistics/daily-end-users",
		app.DailyTerminalsStatisticHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/statistics/token-costs",
		app.DailyTokenCostStatisticHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/statistics/average-session-interactions",
		app.AverageSessionInteractionStatisticHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/statistics/user-satisfaction-rate",
		app.UserSatisfactionRateStatisticHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/statistics/average-response-time",
		app.AverageResponseTimeStatisticHandler)
	router.HandleFunc(ConsoleApiPrefix+"/apps/{app_id}/statistics/tokens-per-second",
		app.TokensPerSecondStatisticHandler)

	//datasets
	router.HandleFunc(ConsoleApiPrefix+"/data-source/integrates", datasets.GetDataSourceApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/data-source/integrates/{binding_id}/{action}", datasets.PostDataSourceApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/notion/pre-import/pages", datasets.DataSourceNotionListApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/notion/workspaces/{workspace_id}/pages/{page_id}/{page_type}/preview",
		datasets.GetDataSourceNotionApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/notion-indexing-estimate", datasets.PostDataSourceNotionApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/notion/sync",
		datasets.DataSourceNotionDatasetSyncApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/notion/sync",
		datasets.DataSourceNotionDocumentSyncApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/datasets", datasets.DatasetListApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}", datasets.DatasetApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/queries", datasets.DatasetQueryApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/indexing-estimate", datasets.DatasetIndexingEstimateApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/related-apps", datasets.DatasetRelatedAppListApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/indexing-status",
		datasets.DatasetIndexingStatusApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/api-keys", datasets.DatasetApiKeyApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/api-keys/{api_key_id}", datasets.DatasetApiDeleteApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/api-base-info", datasets.DatasetApiBaseUrlApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/datasets/process-rule", datasets.GetProcessRuleApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents", datasets.DatasetDocumentListApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/init", datasets.DatasetInitApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/indexing-estimate",
		datasets.DocumentIndexingEstimateApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/batch/{batch}/indexing-estimate",
		datasets.DocumentBatchIndexingEstimateApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/batch/{batch}/indexing-status",
		datasets.DocumentBatchIndexingStatusApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/indexing-status",
		datasets.DocumentIndexingStatusApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}",
		datasets.DocumentDetailApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/processing/{action}",
		datasets.DocumentProcessingApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}",
		datasets.DocumentDeleteApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/metadata",
		datasets.DocumentMetadataApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/status/{action}",
		datasets.DocumentStatusApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/processing/pause",
		datasets.DocumentPauseApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/processing/resume",
		datasets.DocumentRecoverApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/limit", datasets.DocumentLimitApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/segments",
		datasets.DatasetDocumentSegmentListApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/segments/{segment_id}/{action}",
		datasets.DatasetDocumentSegmentApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/segment",
		datasets.DatasetDocumentSegmentAddApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/segments/{segment_id}",
		datasets.DatasetDocumentSegmentUpdateApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/documents/{document_id}/segments/batch_import",
		datasets.PostDatasetDocumentSegmentBatchImportApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/datasets/batch_import_status/{job_id}",
		datasets.GetDatasetDocumentSegmentBatchImportApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/files/upload", datasets.FileApiHandler)
	router.HandleFunc(ConsoleApiPrefix+"/files/{file_id}/preview", datasets.FilePreviewApiHandler)

	router.HandleFunc(ConsoleApiPrefix+"/datasets/{dataset_id}/hit-testing", datasets.HitTestingApiHandler)

	return router
}
