package router_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/sunho/gorani-reader/server/etl/router"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func prepareServer(t *testing.T) (*httpexpect.Expect, *httptest.Server) {
	gorn := util.SetupTestGorani()

	router := router.NewRouter(gorn)
	server := httptest.NewServer(router)
	e := httpexpect.New(t, server.URL)

	return e, server
}
