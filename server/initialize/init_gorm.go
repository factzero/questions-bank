package initialize

import (
	"fmt"
	"server/global"
	"server/model/system"
	"server/model/system/request"
	"server/service"

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
		system.SysOperationRecord{},
		system.FileUploadAndDownload{},
	)
	if err != nil {
		fmt.Println("register table failed")
		return
	}
	fmt.Println("register table success")
}

func InitDB() {
	m := global.GVA_CONFIG.Mysql
	dbInfo := request.InitDB{
		DBType:   "mysql",
		Host:     m.Path,
		Port:     m.Port,
		UserName: m.Username,
		Password: m.Password,
		DBName:   m.Dbname,
	}

	if err := service.ServiceGroupApp.InitDBService.InitDB(dbInfo); err != nil {
		fmt.Println("自动创建数据库失败! ", err.Error())
	}
}
