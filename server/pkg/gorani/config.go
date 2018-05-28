package gorani

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	MysqlURL                string `yaml:"mysql_url"`
	GoMaxProcs              int    `yaml:"go_max_procs"`
	MysqlConnectionPoolSize int    `yaml:"mysql_connection_pool_size"`
	Debug                   bool   `yaml:"debug"`
}

func NewConfig(yamlBytes []byte) (Config, error) {
	conf := Config{}

	err := yaml.Unmarshal(yamlBytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
