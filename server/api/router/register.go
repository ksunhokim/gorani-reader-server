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

	ro.Route("/word", func(r chi.Router) {
		r.Use(mymid.Auth(ro.ap.Mysql, ro.ap.Config.SecretKey))

		r.Route("/known", func(r chi.Router) {
			r.Post("/", ro.AddKnownWords)
		})

		r.Route("/unknown", func(r chi.Router) {
			r.Get("/", ro.GetUnknownWords)
			r.Put("/{word_id:[0-9+]}", ro.PutUnknownWord)
			r.Delete("/{word_id:[0-9+]}", ro.DeleteUnknownWord)
		})
	})

	ro.Route("/genre", func(r chi.Router) {
		r.Get("/", ro.GetGenres)
	})
}
