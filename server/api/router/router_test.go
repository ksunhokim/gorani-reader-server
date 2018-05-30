package router_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/sunho/gorani-reader/server/api/api"
	"github.com/sunho/gorani-reader/server/api/router"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func prepareServer(t *testing.T) (*httpexpect.Expect, *httptest.Server) {
	gorn := util.SetupTestGorani()
	ap, err := api.New(gorn, api.Config{})
	if err != nil {
		panic(err)
	}

	router := router.NewRouter(ap)
	server := httptest.NewServer(router)
	e := httpexpect.New(t, server.URL)

	return e, server
}
