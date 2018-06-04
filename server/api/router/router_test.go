package router_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/sunho/gorani-reader/server/api/api"
	"github.com/sunho/gorani-reader/server/api/router"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func prepareServer(t *testing.T) (*httpexpect.Expect, *httptest.Server, *api.Api) {
	gorn := util.SetupTestGorani()

	aconf, err := api.NewConfig("../aconfig_test.yaml")
	if err != nil {
		panic(err)
	}

	ap, err := api.New(gorn, aconf)
	if err != nil {
		panic(err)
	}

	router := router.NewRouter(ap)
	server := httptest.NewServer(router)
	e := httpexpect.New(t, server.URL)

	return e, server, ap
}
