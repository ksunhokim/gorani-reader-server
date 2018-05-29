package router_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/sunho/gorani-reader/server/api/router"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
)

func prepareServer(t *testing.T) (*httpexpect.Expect, *httptest.Server) {
	bytes, err := ioutil.ReadFile("../config_test.yaml")
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
	router := router.NewRouter(gorn)
	server := httptest.NewServer(router)
	e := httpexpect.New(t, server.URL)

	return e, server
}
