package router

import (
	"github.com/go-chi/chi"
	"github.com/sunho/gorani-reader-server/go/etl/etl"
)

type Router struct {
	chi.Router
	e *etl.Etl
}

func New(e *etl.Etl) *Router {
	r := &Router{
		Router: chi.NewRouter(),
		e:      e,
	}

	return r
}
