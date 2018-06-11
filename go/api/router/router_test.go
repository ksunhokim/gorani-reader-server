package router_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/sunho/gorani-reader-server/go/api/api"
	"github.com/sunho/gorani-reader-server/go/api/router"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
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
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewRequireReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	return e, server, ap
}
