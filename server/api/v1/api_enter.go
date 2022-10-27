package v1

import "server/api/v1/system"

type ApiGroup struct {
	BaseApi system.BaseApi
}

var ApiGroupApp = new(ApiGroup)
