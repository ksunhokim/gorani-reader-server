package main

import (
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/sunho/gorani-reader/server/etl/etl"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
	"github.com/sunho/gorani-reader/server/pkg/log"
)

func main() {
	log.AppName = "etl"

	bytes, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		panic(err)
	}

	conf, err := gorani.NewConfig(bytes)
	if err != nil {
		panic(err)
	}

	gorn, err := gorani.New(conf)
	if err != nil {
		panic(err)
	}

	ebytes, err := ioutil.ReadFile("econfig.yaml")
	if err != nil {
		panic(err)
	}

	econf, err := etl.NewConfig(ebytes)
	if err != nil {
		panic(err)
	}

	e, err := etl.New(gorn, econf)
	if err != nil {
		panic(err)
	}

	err = e.Open()
	if err != nil {
		panic(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	e.Close()
}
