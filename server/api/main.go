package main

import (
	"io/ioutil"

	"github.com/sunho/gorani-reader/server/api/apiconfig"
	"github.com/sunho/gorani-reader/server/pkg/config"
	"github.com/sunho/gorani-reader/server/pkg/log"
)

func main() {
	bytes, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		panic(err)
	}

	conf, err := config.New(bytes)
	if err != nil {
		panic(err)
	}

	abytes, err := ioutil.ReadFile("aconfig.yaml")
	if err != nil {
		panic(err)
	}

	aconf, err := apiconfig.New(abytes)
	if err != nil {
		panic(err)
	}

	serv, err := setup(conf, aconf)
	if err != nil {
		panic(err)
	}

	log.Log(log.TopicSystem.Api(), log.M{
		"info":    "begin listening",
		"address": serv.Addr,
	})

	if err := serv.ListenAndServe(); err != nil {
		panic(err)
	}
}
