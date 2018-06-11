package main

import (
	"net/http"
	"time"

	"github.com/sunho/gorani-reader-server/api/api"
	"github.com/sunho/gorani-reader-server/api/router"
	"github.com/sunho/gorani-reader-server/pkg/gorani"
	"github.com/sunho/gorani-reader-server/pkg/log"
	"github.com/sunho/gorani-reader-server/pkg/util"
)

func setup(conf gorani.Config, aconf api.Config) (*http.Server, error) {
	gorn, err := gorani.New(conf)
	if err != nil {
		return nil, err
	}

	ap, err := api.New(gorn, aconf)
	if err != nil {
		return nil, err
	}

	r := router.NewRouter(ap)
	hs := &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Addr:           aconf.Address,
	}

	return hs, nil
}

func main() {
	log.AppName = "api"

	conf, err := gorani.NewConfig("../config.yaml")
	if err != nil {
		panic(err)
	}

	aconf, err := api.NewConfig("aconfig.yaml")
	if err != nil {
		panic(err)
	}

	serv, err := setup(conf, aconf)
	if err != nil {
		panic(err)
	}

	log.Log(log.TopicSystem, util.M{
		"info":    "begin listening",
		"address": serv.Addr,
	})

	if err := serv.ListenAndServe(); err != nil {
		panic(err)
	}
}
