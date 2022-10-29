package config

type CfgMysql struct {
	CfgGeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *CfgMysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
