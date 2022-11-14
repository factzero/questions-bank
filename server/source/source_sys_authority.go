package source

import (
	"context"
	sysModel "server/model/system"
	"server/service"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initAuthority struct{}

// auto run
func init() {
	service.RegisterInit(service.InitOrderAuthority, &initAuthority{})
}

func (i initAuthority) InitializerName() string {
	return sysModel.SysAuthority{}.TableName()
}

func (i *initAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, service.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysAuthority{})
}

func (i *initAuthority) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysAuthority{})
}

func (i *initAuthority) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, service.ErrMissingDBContext
	}
	entities := []sysModel.SysAuthority{
		{AuthorityId: 9999, AuthorityName: "管理员", DefaultRouter: "dashboard"},
		{AuthorityId: 8888, AuthorityName: "普通用户", DefaultRouter: "dashboard"},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!", sysModel.SysAuthority{}.TableName())
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("authority_id = ?", "9999").
		First(&sysModel.SysAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
