package config

type System struct {
	IgnoreURLs []string `yaml:"ignoreURLs"`
	Port       string   `yaml:"port"`
	PhotoPath  string   `yaml:"photoPath"`
	Env        string   `yaml:"env"`
}
