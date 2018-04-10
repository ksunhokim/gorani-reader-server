package config

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

type ConfigImpl interface {
	GetString(name string, initial string) string
}

var Impl ConfigImpl = EnvImpl{}

const DBName = "engbreaker"

var Debug = true

func GetString(name string, initial string) string {
	return Impl.GetString(name, initial)
}

func GetInt(name string, initial int) int {
	initialStr := strconv.Itoa(initial)
	value := Impl.GetString(name, initialStr)
	i, err := strconv.Atoi(value)
	if Debug {
		if err != nil {
			logrus.Error("Config int field parse error ", name, ":", value)
		}
	}
	return i
}
