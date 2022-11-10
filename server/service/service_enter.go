package service

type ServiceGroup struct {
	UserService
	InitDBService
}

var ServiceGroupApp = new(ServiceGroup)
