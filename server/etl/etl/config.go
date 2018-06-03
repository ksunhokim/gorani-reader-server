package etl

import yaml "gopkg.in/yaml.v2"

type Config struct {
	Address string `yaml:"address"`
}

func NewConfig(yamlBytes []byte) (Config, error) {
	conf := Config{}

	err := yaml.Unmarshal(yamlBytes, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
