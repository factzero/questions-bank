package system

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.AuthorityApi
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
