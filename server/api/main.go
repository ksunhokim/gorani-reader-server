package main

import (
	"io/ioutil"

	"github.com/sunho/gorani-reader/server/api/config"
	"github.com/sunho/gorani-reader/server/api/gorani"
	"github.com/sunho/gorani-reader/server/api/server"
)

func main() {
	bytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	conf, err := config.NewConfig(bytes)
	if err != nil {
		panic(err)
	}

	gorn, err := gorani.NewGorani(conf)
	if err != nil {
		panic(err)
	}

	serv := server.NewServer(gorn)
	serv.ListenAndServe()
}
