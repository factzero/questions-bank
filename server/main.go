package main

import (
	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB)
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	core.RunWindowsServer()
}
