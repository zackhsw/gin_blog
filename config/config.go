package config

import (
	"encoding/json"
	"github.com/go-ini/ini"
	"io/ioutil"
)

type AppConfig struct {
	*ServerConfig `json:"server" ini:"server"`
	*MySQLConfig  `json:"mysql" ini:"mysql"`
	*RedisConfig  `json:"redis" ini:"redis"`
	*LogConfig    `json:"log" ini:"log"`
}

type ServerConfig struct {
	Port int `json:"port" ini:"port"`
}

type MySQLConfig struct {
	Host     string `json:"host" ini:"host"`
	Username string `json:"username" ini:"username"`
	Password string `json:"password" ini:"password"`
	Port     int    `json:"port" ini:"port"`
	DB       string `json:"db" ini:"db"`
}
type RedisConfig struct {
	Host     string `json:"host" ini:"host"`
	Password string `json:"password" ini:"password"`
	Port     int    `json:"port" ini:"port"`
	DB       int    `json:"db" ini:"db"`
}

type LogConfig struct {
	Level      string `json:"level" ini:"level"`
	Filename   string `json:"filename" ini:"filename"`
	MaxSize    int    `json:"maxsize" ini:"maxsize"`
	MaxAge     int    `json:"max_age" ini:"max_age"`
	MaxBackups int    `json:"max_backups" ini:"max_backups"`
}

var Conf = new(AppConfig)

func Init(file string) error {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, Conf); err != nil {
		return err
	}
	return nil
}

func InitFromStr(str string) error {
	if err := json.Unmarshal([]byte(str), Conf); err != nil {
		return err
	}
	return nil
}

func InitFromIni(filename string) error {
	err := ini.MapTo(Conf, filename)
	if err != nil {
		panic(err)
	}
	return nil
}
