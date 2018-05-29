package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sunho/gorani-reader/server/etl/router"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
	"github.com/sunho/gorani-reader/server/pkg/log"
)

const Addr = "localhost:5982"

func setup(conf gorani.Config) (*http.Server, error) {
	gorn, err := gorani.New(conf)
	if err != nil {
		return nil, err
	}

	r := router.NewRouter(gorn)
	hs := &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Addr:           Addr,
	}

	return hs, nil
}

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
