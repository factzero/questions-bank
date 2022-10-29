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
		response.FailWithMessage("��֤�����", c)
	}

	return
}

func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // Ψһǩ��
	claims := j.CreateClaims(systemReq.BaseClaimsRequest{
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("��ȡtokenʧ��", c)
		return
	}
	response.OkWithDetailed(systemRes.SysLoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "��¼�ɹ�", c)

	return
}
