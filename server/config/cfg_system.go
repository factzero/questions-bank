package config

type CfgSystem struct {
	Addr int `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
}
