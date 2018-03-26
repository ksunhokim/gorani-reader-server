package config

import (
	"os"
	"regexp"

	"github.com/sirupsen/logrus"
)

type EnvImpl struct {
}

func (e EnvImpl) GetString(name string, initial string) string {
	value := os.Getenv(name)
	if value == "" {
		value = initial
	}
	if Debug {
		pat := regexp.MustCompile(`[a-zA-Z_]+[a-zA-Z0-9_]*`)
		if !pat.MatchString(name) {
			logrus.Error("Config field name is invalid:", name)
			return ""
		}
		logrus.Info("Evn Config field fetched ", name, ":", value)
	}
	return value
}
