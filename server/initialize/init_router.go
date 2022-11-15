package initialize

import (
	"server/middleware"
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
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitMenuRouter(PrivateGroup)               // 注册菜单路由
		systemRouter.InitUserRouter(PrivateGroup)               // 注册用户路由
		systemRouter.InitAuthorityRouter(PrivateGroup)          // 注册角色路由
		systemRouter.InitSysOperationRecordRouter(PrivateGroup) // 操作记录
	}

	return r
}
