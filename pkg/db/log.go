package db

import (
	"github.com/sirupsen/logrus"
)

type gormLogger struct{}

func (*gormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		logrus.WithFields(logrus.Fields{"module": "gorm", "type": "sql"}).Print(v[3])
	} else if v[0] == "log" {
		logrus.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
