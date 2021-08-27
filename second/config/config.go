package config

import "github.com/jinzhu/configor"

var Config = struct {
	App struct {
		Address  string
		APIKey   string
		HTTPPort string
	}
	DB struct {
		Driver   string
		Host     string
		Port     string
		Name     string
		Username string
		Password string
		Locale   string
	}
}{}

func init() {
	configor.Load(&Config, "config.yaml")
}
