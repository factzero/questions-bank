package system

import (
	"server/global"
	"server/model/common/response"
	systemRes "server/model/system/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var captchaStore = base64Captcha.DefaultMemStore

func (b *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.8, 80)
	cp := base64Captcha.NewCaptcha(driver, captchaStore)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
	}, "验证码获取成功", c)
}
