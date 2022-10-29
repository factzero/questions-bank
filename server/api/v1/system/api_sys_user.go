package system

import (
	"server/global"
	"server/model/common/response"
	"server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBind(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if captchaStore.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username}
		b.TokenNext(c, *u)
	} else {
		response.FailWithMessage("验证码错误", c)
	}

	return
}

func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaimsRequest{
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysLoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)

	return
}
