package parse

import (
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v2"
	"log"
	"mayday/src/supports/fio"
)

var (
	//配置文件位置
	confPath = "config/app.yml"

	// app.conf配置项key定义
	ignoreURLs string = "IgnoreURLs"
	port       string = "Port"
	jwtTimeout string = "JWTTimeout"
	logLevel   string = "LogLevel"
	secret     string = "Secret"
)
var (
	// Conf strut
	Conf iris.Configuration

	// O 自定义配置
	O Other
)

type (
	Other struct {
		IgnoreURLs []string
		Port       string
		JWTTimeout int64
		LogLevel   string
		Secret     string
	}
)

func AppOtherParse() {

	log.Print("### Init app conf")

	appData, err := file_io.Load(confPath)
	if err != nil {
		log.Print("err : ", err)
	}

	c := iris.DefaultConfiguration()
	if err1 := yaml.Unmarshal(appData, &c); err1 != nil {
		log.Print("err : ", err1)
	}

	Conf = c

	iURLs := c.GetOther()[ignoreURLs].([]interface{})
	for _, v := range iURLs {
		O.IgnoreURLs = append(O.IgnoreURLs, v.(string))
	}

	O.Port = c.GetOther()[port].(string)

	jTimeout := c.GetOther()[jwtTimeout].(int)
	O.JWTTimeout = int64(jTimeout)
	O.LogLevel = c.GetOther()[logLevel].(string)
	O.Secret = c.GetOther()[secret].(string)

}
