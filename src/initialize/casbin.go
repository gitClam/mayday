package initialize

import (
	"mayday/src/global"
	"mayday/src/middleware"
	"os"
)

func Casbin() {
	casbin, err := middleware.Casbin()
	if err != nil {
		os.Exit(0)
	}
	global.GVA_CASBIN = casbin
}
