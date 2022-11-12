package service

import (
	"errors"
	"fmt"
	"server/global"
	"server/model/common/request"
	"server/model/system"
	"server/utils"
)

type UserService struct{}

func (s *UserService) Login(u *system.SysUser) (*system.SysUser, error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err := global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("√‹¬Î¥ÌŒÛ")
		}
	}

	return &user, err
}

func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}
