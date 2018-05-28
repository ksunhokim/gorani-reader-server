package router

import (
	"github.com/go-chi/chi"
	"github.com/sunho/gorani-reader/server/api/api"
)

type Router struct {
	chi.Router
	ap *api.Api
}

func NewRouter(ap *api.Api) *Router {
	r := &Router{
		Router: chi.NewRouter(),
		ap:     ap,
	}
	r.registerHandlers()

	return r
}
