package source

import (
	"context"
	. "server/model/system"
	"server/service"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initMenu struct{}

func init() {
	service.RegisterInit(service.InitOrderMenu, &initMenu{})
}

func (i initMenu) InitializerName() string {
	return SysBaseMenu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, service.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&SysBaseMenu{})
}

func (i *initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&SysBaseMenu{})
}

func (i *initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, service.ErrMissingDBContext
	}
	entities := []SysBaseMenu{
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: Meta{Title: "个人信息", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/dashboard.vue", Sort: 1, Meta: Meta{Title: "仪表盘", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: Meta{Title: "超级管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: Meta{Title: "用户管理", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: Meta{Title: "操作历史", Icon: "pie-chart"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 8, Meta: Meta{Title: "服务器状态", Icon: "cloudy"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "exam", Name: "exam", Component: "view/exam/exam.vue", Sort: 1, Meta: Meta{Title: "试卷管理", Icon: "document"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "item", Name: "itemBank", Component: "view/itemBank/itemBank.vue", Sort: 1, Meta: Meta{Title: "试题管理", Icon: "management"}},
		{MenuLevel: 0, Hidden: false, ParentId: "8", Path: "summary", Name: "summary", Component: "view/itemBank/summary/summary.vue", Sort: 2, Meta: Meta{Title: "题库汇总", Icon: "notebook"}},
		{MenuLevel: 0, Hidden: false, ParentId: "8", Path: "upload", Name: "upload", Component: "view/itemBank/upload/upload.vue", Sort: 2, Meta: Meta{Title: "试题上传", Icon: "upload"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "workbook", Name: "workbook", Component: "view/itemBank/workbook/workbook.vue", Sort: 2, Meta: Meta{Title: "练习册", Icon: "notebook"}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ?", "dashboard").First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
