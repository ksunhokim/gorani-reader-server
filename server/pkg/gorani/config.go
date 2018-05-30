package gorani

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	MysqlURL                string `yaml:"mysql_url"`
	GoMaxProcs              int    `yaml:"go_max_procs"`
	MysqlConnectionPoolSize int    `yaml:"mysql_connection_pool_size"`
	Debug                   bool   `yaml:"debug"`
	S3Id                    string `yaml:"s3_id"`
	S3Secret                string `yaml:"s3_secret"`
	S3EndPoint              string `yaml:"s3_endpoint"`
	S3Ssl                   bool   `yaml:"s3_ssl"`
}

func NewConfig(yamlBytes []byte) (Config, error) {
	conf := Config{}

	err := yaml.Unmarshal(yamlBytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
