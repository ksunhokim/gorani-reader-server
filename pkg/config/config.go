package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Debug = false

func init() {
	if GetString("DEBUG", "false") == "true" {
		Debug = true
	}
}

func GetString(name, defaul string) string {
	s := os.Getenv(name)
	if s == "" {
		s = defaul
	}
	if Debug {
		logrus.WithFields(
			logrus.Fields{
				"name":  name,
				"value": s,
			},
		).Info("config fetched")
	}
	return s
}
