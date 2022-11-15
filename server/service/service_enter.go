package service

type ServiceGroup struct {
	UserService
	InitDBService
	AuthorityService
	MenuService
	OperationRecordService
}

var ServiceGroupApp = new(ServiceGroup)
