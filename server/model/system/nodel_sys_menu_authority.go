package system

type SysMenu struct {
	SysBaseMenu
	MenuId      string          `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId uint            `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu       `json:"children" gorm:"-"`
	Btns        map[string]uint `json:"btns" gorm:"-"`
}
