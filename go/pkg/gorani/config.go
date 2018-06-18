package gorani

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	MysqlConnectionPoolSize = 5
	MysqlConnectionLimit    = 10
	RedisConnectionPoolSize = 5
)

type Config struct {
	Debug              bool   `yaml:"debug"`
	MysqlURL           string `yaml:"mysql_url"`
	RedisURL           string `yaml:"redis_url"`
	GoMaxProcs         int    `yaml:"go_max_procs"`
	S3Id               string `yaml:"s3_id"`
	S3Secret           string `yaml:"s3_secret"`
	S3EndPoint         string `yaml:"s3_endpoint"`
	S3Ssl              bool   `yaml:"s3_ssl"`
	IsGarbageCollector bool   `yaml:"is_garbage_collector"`
	IsConsumer         bool   `yaml:"is_consumer"`
}

func NewConfig(path string) (Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	conf := Config{}
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
