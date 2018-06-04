package etl

import (
	"io/ioutil"

	"github.com/sunho/gorani-reader/server/pkg/gorani"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	gorani.Config
	Address string `yaml:"address"`
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
