package initialize

import "mayday/src/utils"

func RegisterLocalTimeDecoder() {
	utils.TimeDecoder.RegisterLocalTimeDecoder()
}
