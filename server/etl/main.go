package main

import (
	"io/ioutil"

	"github.com/sunho/gorani-reader/server/pkg/config"
	"github.com/sunho/gorani-reader/server/pkg/log"
)

func main() {
	log.AppName = "etl"

	bytes, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		panic(err)
	}

	conf, err := config.New(bytes)
	if err != nil {
		panic(err)
	}

	serv, err := setup(conf)
	if err != nil {
		panic(err)
	}

	log.Log(log.TopicSystem, log.M{
		"info":    "begin listening",
		"address": serv.Addr,
	})

	if err := serv.ListenAndServe(); err != nil {
		panic(err)
	}
}
