package main

import (
	"net/http"
	"time"

	"github.com/sunho/gorani-reader/server/api/api"
	"github.com/sunho/gorani-reader/server/api/router"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
)

func setup(conf gorani.Config, aconf api.Config) (*http.Server, error) {
	gorn, err := gorani.New(conf)

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
		Addr:           aconf.ApiAddress,
	}

	return hs, nil
}
