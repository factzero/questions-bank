package v1

import "server/api/v1/system"

type ApiGroup struct {
	BaseApi          system.BaseApi
	AuthorityApi     system.AuthorityApi
	AuthorityMenuApi system.AuthorityMenuApi
}

var ApiGroupApp = new(ApiGroup)
