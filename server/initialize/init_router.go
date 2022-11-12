package initialize

import (
	"server/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	systemRouter := router.RouterGroupApp.System
	PublicGroup := r.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := r.Group("/api/v1")
	{
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)      // 注册用户路由
		systemRouter.InitAuthorityRouter(PrivateGroup) // 注册角色路由
	}

	return r
}
