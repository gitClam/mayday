package config

type Jwt struct {
	JWTTimeout int    `yaml:"jwtTimeout"` //second
	Secret     string `yaml:"secret"`     //加密方式
}
