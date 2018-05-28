package router

import (
	"github.com/go-chi/chi"
	"github.com/sunho/gorani-reader/server/pkg/gorani"
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

	return r
}
