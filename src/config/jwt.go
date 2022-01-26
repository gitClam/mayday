package config

type JWT struct {
	JWTTimeout        int    `yaml:"jwtTimeout"` //second
	Secret            string `yaml:"secret"`     //加密方式
	DefaultContextKey string `yaml:"defaultContextKey"`
}
