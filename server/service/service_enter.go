package service

type ServiceGroup struct {
	UserService
	InitDBService
	AuthorityService
	MenuService
	OperationRecordService
	FileUploadAndDownloadService
}

var ServiceGroupApp = new(ServiceGroup)
