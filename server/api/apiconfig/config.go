package apiconfig

import (
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ApiAddress  string `yaml:"api_address"`
	ServicesUrl string `yaml:"services_url"`
	SecretKey   string `yaml:"secret_key"`
}

func New(yamlBytes []byte) (Config, error) {
	conf := Config{}

	err := yaml.Unmarshal(yamlBytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
