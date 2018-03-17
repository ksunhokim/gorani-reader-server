package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/api"
)

func main() {
	server := api.NewHTTPServer()
	err := server.Start()
	if err != nil {
		logrus.Panic(err)
	}
}
