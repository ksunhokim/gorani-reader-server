package config

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	MysqlURL                string `yaml:"mysql_url"`
	RedisURL                string `yaml:"redis_url"`
	GoMaxProcs              int    `yaml:"go_max_procs"`
	MysqlConnectionPoolSize int    `yaml:"mysql_connection_pool_size"`
	RedisConnectionPoolSize int    `yaml:"redis_connection_pool_size"`
	Debug                   bool   `yaml:"debug"`
}

func New(yamlBytes []byte) (Config, error) {
	conf := Config{}

	err := yaml.Unmarshal(yamlBytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
