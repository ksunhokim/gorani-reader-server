package config

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	MysqlURL                string `yaml:"mysql_url"`
	RedisURL                string `yaml:"redis_url"`
	ApiAddress              string `yaml:"api_address"`
	GoMaxProcs              int    `yaml:"go_max_procs"`
	MysqlConnectionPoolSize int    `yaml:"mysql_connection_pool_size"`
	RedisConnectionPoolSize int    `yaml:"redis_connection_pool_size"`
	LoggerType              string `yaml:"logger_type"`
	FluentHost              string `yaml:"fluent_host"`
	FluentPort              int    `yaml:"fluent_port"`
	SecretKey               string `yaml:"secret_key"`
}

const (
	LoggerTypeFluent = "fluentd"
	LoggerTypeStdout = "stdout"
	LoggerTypeBoth   = "both"
)

func NewConfig(yamlBytes []byte) (Config, error) {
	conf := Config{}

	err := yaml.Unmarshal(yamlBytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
