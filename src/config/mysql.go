package config

import (
	"fmt"
)

type DB struct {
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

func (info *DB) GetConnURL() (url string) {
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		info.User,
		info.Password,
		info.Host,
		info.Port,
		info.Database,
		info.Charset)
	return
}
