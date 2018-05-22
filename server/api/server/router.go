package server

import (
	"github.com/go-chi/chi"
	"github.com/sunho/gorani-reader/server/api/gorani"
)

type Router struct {
	chi.Router
	gorn *gorani.Gorani
}

func NewRouter(gorn *gorani.Gorani) *Router {
	r := &Router{
		Router: chi.NewRouter(),
		gorn:   gorn,
	}
	r.registerHandlers()

	return r
}
