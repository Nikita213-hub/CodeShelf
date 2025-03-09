package config

import (
	"github.com/Nikita213-hub/CodeShelf/db"
	"github.com/Nikita213-hub/CodeShelf/httpserver"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	HttpServerCfg *httpserver.HttpCfgRoot
	DbCfg         *db.CfgRoot
	//...
}

func NewCfg() *Config {
	return &Config{
		HttpServerCfg: &httpserver.HttpCfgRoot{
			HttpServerCfg: &httpserver.Config{},
		},
		DbCfg: &db.CfgRoot{
			DbConfig: &db.Config{},
		},
	}
}

func (c *Config) LoadCfg() error {
	yamlData, err := os.ReadFile("../config/config.global.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlData, c.HttpServerCfg)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlData, c.DbCfg)
	if err != nil {
		return err
	}
	return nil
}
