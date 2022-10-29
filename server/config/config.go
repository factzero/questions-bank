package config

type CfgServer struct {
	System  CfgSystem  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha CfgCaptcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
