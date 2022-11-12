package service

type ServiceGroup struct {
	UserService
	InitDBService
	AuthorityService
}

var ServiceGroupApp = new(ServiceGroup)
