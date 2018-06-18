package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sunho/gorani-reader-server/go/api/api"
	"github.com/sunho/gorani-reader-server/go/api/router"
	"github.com/sunho/gorani-reader-server/go/pkg/gorani"
)

func main() {
	conf, err := gorani.NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	aconf, err := api.NewConfig("aconfig.yaml")
	if err != nil {
		panic(err)
	}

	gorn, err := gorani.New(conf)
	if err != nil {
		panic(err)
	}

	ap, err := api.New(gorn, aconf)
	if err != nil {
		panic(err)
	}

	handler := router.New(ap)
	gorn.Start(":80", handler)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	gorn.End()

}
