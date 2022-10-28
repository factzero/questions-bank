package core

import "server/initialize"

func RunWindowsServer() {
	r := initialize.Routers()

	r.Run(":6000")
}
