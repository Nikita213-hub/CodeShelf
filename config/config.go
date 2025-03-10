package config

import (
	"github.com/Nikita213-hub/CodeShelf/db"
	"github.com/Nikita213-hub/CodeShelf/httpserver"
	"github.com/Nikita213-hub/CodeShelf/httpservice"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	HttpServerCfg  *httpserver.Config  `yaml:"http_server"`
	DbCfg          *db.Config          `yaml:"database"`
	HttpServiceCfg *httpservice.Config `yaml:"http_service"`
	//...
}

func NewCfg() *Config {
	return &Config{
		HttpServerCfg:  &httpserver.Config{},
		DbCfg:          &db.Config{},
		HttpServiceCfg: &httpservice.Config{},
	}
}

func (c *Config) LoadCfg() error {
	yamlData, err := os.ReadFile("../config/config.global.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlData, c)
	if err != nil {
		return err
	}
	//err = yaml.Unmarshal(yamlData, c.DbCfg)
	//if err != nil {
	//	return err
	//}
	return nil
}
