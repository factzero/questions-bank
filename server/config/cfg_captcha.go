package config

type CfgCaptcha struct {
	KeyLong   int `mapstructure:"keylong" json:"keylong" yaml:"keylong"`       // 验证码长度
	ImgWidth  int `mapstructure:"imgwidth" json:"imgwidth" yaml:"imgwidth"`    // 验证码宽度
	ImgHeight int `mapstructure:"imgheight" json:"imgheight" yaml:"imgheight"` // 验证码高度
}
