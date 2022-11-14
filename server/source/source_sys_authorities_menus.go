package source

import (
	"context"

	sysModel "server/model/system"
	"server/service"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initMenuAuthority struct{}

// auto run
func init() {
	service.RegisterInit(service.InitOrderMenuAuthority, &initMenuAuthority{})
}

func (i *initMenuAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil // do nothing
}

func (i *initMenuAuthority) TableCreated(ctx context.Context) bool {
	return false // always replace
}

func (i initMenuAuthority) InitializerName() string {
	return "sys_menu_authorities"
}

func (i *initMenuAuthority) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, service.ErrMissingDBContext
	}
	authorities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return ctx, errors.Wrap(service.ErrMissingDependentContext, "创建 [菜单-权限] 关联失败, 未找到权限表初始化数据")
	}
	menus, ok := ctx.Value(initMenu{}.InitializerName()).([]sysModel.SysBaseMenu)
	if !ok {
		return next, errors.Wrap(errors.New(""), "创建 [菜单-权限] 关联失败, 未找到菜单表初始化数据")
	}
	next = ctx
	// 9999 管理员
	if err = db.Model(&authorities[0]).Association("SysBaseMenus").Replace(menus[:5]); err != nil {
		return next, err
	}
	if err = db.Model(&authorities[0]).Association("SysBaseMenus").Append(menus[6:]); err != nil {
		return next, err
	}

	// 8888 普通用户
	menu8888 := menus[:2]
	menu8888 = append(menu8888, menus[9])
	if err = db.Model(&authorities[1]).Association("SysBaseMenus").Replace(menu8888); err != nil {
		return next, err
	}
	return next, nil
}

func (i *initMenuAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	auth := &sysModel.SysAuthority{}
	if ret := db.Model(auth).
		Where("authority_id = ?", 9999).Preload("SysBaseMenus").Find(auth); ret != nil {
		if ret.Error != nil {
			return false
		}
		return len(auth.SysBaseMenus) > 0
	}
	return false
}
