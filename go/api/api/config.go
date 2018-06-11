package api

import (
	"io/ioutil"

	"github.com/sunho/gorani-reader-server/go/pkg/gorani"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	// gorani config is initialized by api.New
	gorani.Config
	Address     string `yaml:"address"`
	ServicesUrl string `yaml:"services_url"`
	SecretKey   string `yaml:"secret_key"`
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
