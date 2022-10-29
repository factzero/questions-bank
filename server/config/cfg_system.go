package config

type CfgSystem struct {
	Addr   int    `mapstructure:"addr" json:"addr" yaml:"addr"`       // 端口值
	DbType string `mapstructure:"dbtype" json:"dbtype" yaml:"dbtype"` // 数据库类型:mysql
}
