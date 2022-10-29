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
	)
	if err != nil {
		fmt.Println("register table failed")
		return
	}
	fmt.Println("register table success")
}
