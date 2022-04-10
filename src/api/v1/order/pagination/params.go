package pagination

import (
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"mayday/src/global"
)

func RequestParams(c iris.Context) map[string]interface{} {
	params := make(map[string]interface{}, 10)

	if c.Request().Form == nil {
		if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
			global.GVA_LOG.Error("", zap.Error(err))
		}
	}

	if len(c.Request().Form) > 0 {
		for key, value := range c.Request().Form {
			if key == "page" || key == "per_page" || key == "sort" {
				continue
			}
			params[key] = value[0]
		}
	}

	return params
}
