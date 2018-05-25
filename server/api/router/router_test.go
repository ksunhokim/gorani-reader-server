package router_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/sunho/gorani-reader/server/api/config"
	"github.com/sunho/gorani-reader/server/api/gorani"
	"github.com/sunho/gorani-reader/server/api/router"
)

func prepareServer(t *testing.T) (*httpexpect.Expect, *httptest.Server) {
	bytes, err := ioutil.ReadFile("../config.yaml")
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
	router := router.NewRouter(gorn)
	server := httptest.NewServer(router)
	e := httpexpect.New(t, server.URL)
	return e, server
}
