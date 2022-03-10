package initialize

import (
	"mayday/src/model/common/timedecoder"
)

func RegisterLocalTimeDecoder() {
	timedecoder.TimeDecoder.RegisterLocalTimeDecoder()
}
