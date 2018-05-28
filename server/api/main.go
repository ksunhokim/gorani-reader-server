package main

import (
	"io/ioutil"

	"github.com/sunho/gorani-reader/server/api/api"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
	"github.com/sunho/gorani-reader/server/pkg/log"
)

func main() {
	log.AppName = "api"

	bytes, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		panic(err)
	}

	conf, err := gorani.NewConfig(bytes)
	if err != nil {
		panic(err)
	}

	abytes, err := ioutil.ReadFile("aconfig.yaml")
	if err != nil {
		panic(err)
	}

	aconf, err := api.NewConfig(abytes)
	if err != nil {
		panic(err)
	}

	serv, err := setup(conf, aconf)
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
