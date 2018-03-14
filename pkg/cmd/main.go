package main

import (
	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/api"
	_ "github.com/sunho/engbreaker/pkg/db"
)

func main() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	formatter.DisableColors = true
	logrus.SetFormatter(formatter)
	server := api.NewHTTPServer()
	err := server.Start()
	if err != nil {
		logrus.Panic(err)
	}
}
