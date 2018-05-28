package api

import (
	"io/ioutil"

	"github.com/sunho/gorani-reader/server/api/apiconfig"
	"github.com/sunho/gorani-reader/server/pkg/auth"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
)

type Api struct {
	Gorn     *gorani.Gorani
	Config   apiconfig.Config
	Services auth.Services
}

func New(gorn *gorani.Gorani, conf apiconfig.Config) (*Api, error) {
	s, err := createServices(conf)
	if err != nil {
		return nil, err
	}

	return &Api{
		Gorn:     gorn,
		Config:   conf,
		Services: s,
	}, nil
}

func createServices(conf apiconfig.Config) (auth.Services, error) {
	bytes, err := ioutil.ReadFile(conf.ServicesUrl)
	if err != nil {
		return auth.Services{}, err
	}

	services, err := auth.NewServices(bytes)
	if err != nil {
		return auth.Services{}, err
	}

	return services, nil
}
