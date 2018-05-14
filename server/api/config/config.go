package config

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	MysqlURL           string
	RedisURL           string
	MLQueuePort        int
	MLQueueSize        int
	ApiPort            int
	GoMaxProcs         int
	ConnectionPoolSize int
}

func NewConfig(bytes []byte) (Config, error) {
	conf := Config{}

	err := yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
