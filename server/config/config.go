package config

type CfgServer struct {
	System  CfgSystem  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha CfgCaptcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	JWT     CfgJWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysql   CfgMysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
}
