package system

type RouterGroup struct {
	BaseRouter
	MenuRouter
	UserRouter
	AuthorityRouter
	OperationRecordRouter
	FileUploadAndDownloadRouter
}
