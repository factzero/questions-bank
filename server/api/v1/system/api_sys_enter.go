package system

import "server/service"

var (
	userService            = service.ServiceGroupApp.UserService
	authorityService       = service.ServiceGroupApp.AuthorityService
	menuService            = service.ServiceGroupApp.MenuService
	operationRecordService = service.ServiceGroupApp.OperationRecordService
)
