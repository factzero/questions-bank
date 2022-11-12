package service

import (
	"server/global"
	"server/model/common/request"
	"server/model/system"
)

type AuthorityService struct{}

func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAuthority{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Find(&authority).Error
	return authority, total, err
}
