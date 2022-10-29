package config

type CfgGeneralDB struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                         // 服务器地址:端口
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                         //:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                   // 高级配置
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                   // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`             // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`             // 数据库密码
	MaxIdleConns int    `mapstructure:"maxidleconns" json:"maxidleconns" yaml:"maxidleconns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"maxopenconns" json:"maxopenconns" yaml:"maxopenconns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"logmode" json:"logmode" yaml:"logmode"`                // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"logzap" json:"logzap" yaml:"logzap"`                   // 是否通过zap写入日志文件
}
