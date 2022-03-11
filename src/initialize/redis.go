package initialize

import (
	"mayday/src/global"
	"os"
)

func Redis() {
	if global.GVA_REDIS != nil {
		os.Exit(0)
	}
}
