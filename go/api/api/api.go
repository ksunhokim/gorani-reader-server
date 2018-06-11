package api

import (
	"github.com/sunho/gorani-reader-server/go/api/services"
	"github.com/sunho/gorani-reader-server/go/pkg/auth"
	"github.com/sunho/gorani-reader-server/go/pkg/gorani"
)

type Api struct {
	*gorani.Gorani
	Config   Config
	Services auth.Services
}

func New(gorn *gorani.Gorani, conf Config) (*Api, error) {
	ap := &Api{
		Gorani:   gorn,
		Config:   conf,
		Services: services.New(),
	}
	ap.Config.Config = gorn.Config

	return ap, nil
}
