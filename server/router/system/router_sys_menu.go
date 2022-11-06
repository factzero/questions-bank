package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.AuthorityMenuApi
	{
		menuRouter.POST("getMenu", authorityMenuApi.GetMenu)
	}

	return menuRouter
}
