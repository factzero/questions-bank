package config

type CfgJWT struct {
	SigningKey  string `mapstructure:"signingkey" json:"signingkey" yaml:"signingkey"`    // jwt签名
	ExpiresTime string `mapstructure:"expirestime" json:"expirestime" yaml:"expirestime"` // 过期时间
	BufferTime  string `mapstructure:"buffertime" json:"buffertime" yaml:"buffertime"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                // 签发者
}
