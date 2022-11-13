package service

type ServiceGroup struct {
	UserService
	InitDBService
	AuthorityService
	MenuService
}

var ServiceGroupApp = new(ServiceGroup)
