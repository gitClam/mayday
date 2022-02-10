package middleware

import (
	"github.com/kataras/iris/v12/context"
	"mayday/src/global"
	"regexp"
)

func ServeHTTP(ctx context.Context) {

	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(ctx.Path()) {
		ctx.Next()
		return
	}

	// jwt token拦截
	if !Serve(ctx) {
		return
	}

	// 系统菜单不进行权限拦截
	/*if !strings.Contains(path, "/sysMenu") {
		// casbin权限拦截
		ok := casbins.CheckPermissions(ctx)
		if !ok {
			return
		}
	}*/

	ctx.Next()
}

func checkURL(reqPath string) bool {

	for _, v := range global.GVA_CONFIG.System.IgnoreURLs {
		if yes, _ := regexp.Match(v, []byte(reqPath)); yes {
			return true
		}
	}
	return false
}
