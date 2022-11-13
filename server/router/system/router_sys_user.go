package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register) // 管理员注册账号
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
	}
}
