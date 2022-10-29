package system

type SysUser struct {
	CommonModel
	UUID     string `json:"uuid" gorm:"index;comment:用户UUID"`
	Username string `json:"username" gorm:"index;comment:用户登录名"` // 用户登录名
	Password string `json:"-"  gorm:"comment:用户登录密码"`            // 用户登录密码
}
