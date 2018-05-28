package router

import (
	"github.com/go-chi/chi"
	chimid "github.com/go-chi/chi/middleware"
	mymid "github.com/sunho/gorani-reader/server/pkg/middleware"
)

func (ro *Router) registerHandlers() {
	ro.Use(chimid.RealIP)
	ro.Use(mymid.RequestId)
	ro.Use(mymid.Logger)
	ro.Use(mymid.Recoverer)

	ro.Route("/user", func(r chi.Router) {
		r.Post("/withOauth", ro.UserWithOauth)
	})
}
