package source

import (
	"context"
	sysModel "server/model/system"
	"server/service"
	"server/utils"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type initUser struct{}

// auto run
func init() {
	service.RegisterInit(service.InitOrderUser, &initUser{})
}

func (i initUser) InitializerName() string {
	return sysModel.SysUser{}.TableName()
}

func (i *initUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, service.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysUser{})
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, service.ErrMissingDBContext
	}
	password := utils.BcryptHash("123456")
	adminPassword := utils.BcryptHash("123456")

	entities := []sysModel.SysUser{
		{
			UUID:          uuid.NewV4(),
			Username:      "admin",
			Password:      adminPassword,
			NickName:      "Mr.zero",
			HeaderImg:     "https://qmplusimg.henrongyi.top/gva_header.jpg",
			AuthorityName: "管理员",
			AuthorityId:   9999,
			Phone:         "1111111111",
			Email:         "1111111111@qq.com",
		},
		{
			UUID:          uuid.NewV4(),
			Username:      "test",
			Password:      password,
			NickName:      "Mr.test",
			HeaderImg:     "https://qmplusimg.henrongyi.top/1572075907logo.png",
			AuthorityName: "普通用户",
			AuthorityId:   8888,
			Phone:         "2222222222",
			Email:         "2222222222@qq.com"},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysUser{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	authorityEntities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return next, errors.Wrap(service.ErrMissingDependentContext, "创建 [用户-权限] 关联失败, 未找到权限表初始化数据")
	}
	if err = db.Model(&entities[0]).Association("Authorities").Replace(&authorityEntities[0]); err != nil {
		return next, err
	}
	if err = db.Model(&entities[1]).Association("Authorities").Replace(&authorityEntities[1]); err != nil {
		return next, err
	}
	return next, err
}

func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysUser
	if errors.Is(db.Where("username = ?", "test").
		Preload("Authorities").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Authorities) > 0 && record.Authorities[0].AuthorityId == 6666
}
