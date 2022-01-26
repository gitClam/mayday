package parse

import (
	"gopkg.in/yaml.v2"
	"log"
	"mayday/src/utils"
)

var DbConfigPath = "config/db.yml"
var DBConfig DB

func DBSettingParse() {
	log.Print("### Init db conf")

	dbData, err := utils.IO.Load(DbConfigPath)
	if err != nil {
		log.Print("err : ", err)
	}

	if err1 := yaml.Unmarshal(dbData, &DBConfig); err1 != nil {
		log.Print("err : ", err1)
	}

}

type DB struct {
	Master DBConfigInfo
	//Slave  DBConfigInfo
}

type DBConfigInfo struct {
	Dialect      string `yaml:"dialect"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Database     string `yaml:"database"`
	Charset      string `yaml:"charset"`
	ShowSql      bool   `yaml:"showSql"`
	LogLevel     string `yaml:"logLevel"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`

	//ParseTime       bool   `yaml:"parseTime"`
	//MaxIdleConns    int    `yaml:"maxIdleConns"`
	//MaxOpenConns    int    `yaml:"maxOpenConns"`
	//ConnMaxLifetime int64  `yaml:"connMaxLifetime: 10"`
	//Sslmode         string `yaml:"sslmode"`
}
