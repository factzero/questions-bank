package main

import (
	"server/core"
	"server/global"
)

func main() {
	global.GVA_VP = core.Viper()

	core.RunWindowsServer()
}
