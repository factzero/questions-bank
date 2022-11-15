package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.AuthorityMenuApi
	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu)
	}

	return menuRouterWithoutRecord
}
