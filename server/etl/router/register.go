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

	ro.Route("/word", func(r chi.Router) {
		r.Get("/", ro.GetWords)

		r.Route("/{id:[0-9]+}", func(r chi.Router) {
			r.Use(ro.WordCtx)
			r.Get("/", ro.GetWord)
			r.Delete("/", ro.DeleteWord)
		})
	})
}
