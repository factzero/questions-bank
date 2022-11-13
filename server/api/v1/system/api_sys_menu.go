package system

import (
	"fmt"
	"server/model/common/response"
	"server/model/system"
	systemRes "server/model/system/response"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type AuthorityMenuApi struct{}

func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c))
	if err != nil {
		fmt.Println("获取失败!", err.Error())
		response.FailWithMessage("获取失败", c)
	}
	if menus == nil {
		menus = []system.SysMenu{}
	}
	response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
}
