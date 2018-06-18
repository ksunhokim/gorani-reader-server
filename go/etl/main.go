package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sunho/gorani-reader-server/go/etl/etl"
	"github.com/sunho/gorani-reader-server/go/etl/router"
	"github.com/sunho/gorani-reader-server/go/pkg/gorani"
)

func main() {
	conf, err := gorani.NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	econf, err := etl.NewConfig("econfig.yaml")
	if err != nil {
		panic(err)
	}

	gorn, err := gorani.New(conf)
	if err != nil {
		panic(err)
	}

	e, err := etl.New(gorn, econf)
	if err != nil {
		panic(err)
	}

	handler := router.New(e)
	gorn.Start(":2018", handler)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	gorn.End()
}
