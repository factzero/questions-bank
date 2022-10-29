package service

import (
	"fmt"
	"server/global"
	"server/model/system"
)

type UserService struct{}

func (s *UserService) Login(u *system.SysUser) (*system.SysUser, error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err := global.GVA_DB.Where("username = ?", u.Username).First(&user).Error
	fmt.Println(err)

	return &user, err
}
