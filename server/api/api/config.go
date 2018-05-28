package api

import (
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	RedisURL                string `yaml:"redis_url"`
	RedisConnectionPoolSize int    `yaml:"redis_connection_pool_size"`
	ApiAddress              string `yaml:"api_address"`
	ServicesUrl             string `yaml:"services_url"`
	SecretKey               string `yaml:"secret_key"`
}

func NewConfig(yamlBytes []byte) (Config, error) {
	conf := Config{}

	err := yaml.Unmarshal(yamlBytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
