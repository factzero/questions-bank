package config

type Local struct {
	Path      string `mapstructure:"path" json:"path" yaml:"path"`                // 本地文件访问路径
	StorePath string `mapstructure:"storepath" json:"storepath" yaml:"storepath"` // 本地文件存储路径
}
