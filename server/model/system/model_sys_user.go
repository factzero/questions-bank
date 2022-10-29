package system

type SysUser struct {
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}
