package initialize

import (
	"fmt"
	"server/global"
	"server/model/system"

	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysAuthority{},
	)
	if err != nil {
		fmt.Println("register table failed")
		return
	}
	fmt.Println("register table success")
}

func InitDB() {
	// dbInfo := request.InitDB{
	// 	UserName: "root",
	// 	Password: "123456",
	// 	DBName:   "questionsbank",
	// }

	// if err := service.ServiceGroupApp.InitDBService.InitDB(dbInfo); err != nil {
	// 	fmt.Println("自动创建数据库失败! ", err.Error())
	// }
}
