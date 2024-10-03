package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port     int    `yaml:"port"`
	LogLevel string `yaml:"log_level"`
}

type TableConfig struct {
	Name    string   `yaml:"name"`
	Columns []string `yaml:"columns"`
}

type ConfluencePageConfig struct {
	ID            string        `yaml:"id"`
	URL           string        `yaml:"url"`
	CacheDuration int           `yaml:"cache_duration"`
	Tables        []TableConfig `yaml:"tables"`
}

type Config struct {
	Server     ServerConfig           `yaml:"server"`
	Confluence []ConfluencePageConfig `yaml:"confluence"`
}

var AppConfig *Config

func LoadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		return err
	}
	return nil
}
