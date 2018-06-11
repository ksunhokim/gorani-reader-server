package main

import (
	"os"
	"os/signal"

	"github.com/sunho/gorani-reader-server/etl/etl"
	"github.com/sunho/gorani-reader-server/etl/service"
	"github.com/sunho/gorani-reader-server/pkg/gorani"
	"github.com/sunho/gorani-reader-server/pkg/log"
	"github.com/sunho/gorani-reader-server/pkg/util"
)

func setup(conf gorani.Config, econf etl.Config) (*service.Service, error) {
	gorn, err := gorani.New(conf)
	if err != nil {
		return nil, err
	}

	e, err := etl.New(gorn, econf)
	if err != nil {
		return nil, err
	}

	serv := service.New(e)

	return serv, nil
}

func main() {
	log.AppName = "etl"

	conf, err := gorani.NewConfig("../config.yaml")
	if err != nil {
		panic(err)
	}

	econf, err := etl.NewConfig("econfig.yaml")
	if err != nil {
		panic(err)
	}

	serv, err := setup(conf, econf)
	if err != nil {
		panic(err)
	}

	log.Log(log.TopicSystem, util.M{
		"info":    "begin listening",
		"address": serv.Addr,
	})

	err = serv.Open()
	if err != nil {
		panic(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	serv.Close()
}
