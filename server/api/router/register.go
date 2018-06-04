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

	ro.Route("/wordbook", func(r chi.Router) {
		r.Use(mymid.Auth(ro.ap.Mysql, ro.ap.Config.SecretKey))

		r.Get("/", ro.GetWordbook)
		r.Post("/", ro.AddWordbook)

		r.Route("/{uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", func(r chi.Router) {
			r.Use(ro.WordbookCtx)

			r.Put("/", ro.PutWordbook)
			r.Delete("/", ro.DeleteWordbook)

			r.Get("/entries", ro.GetWordbookEntries)
			r.Put("/entries", ro.PutWordbookEntries)
		})
	})

	ro.Route("/word", func(r chi.Router) {
		r.Use(mymid.Auth(ro.ap.Mysql, ro.ap.Config.SecretKey))

		r.Route("/unknown", func(r chi.Router) {
			r.Use(ro.UnknownWordbookCtx)

			r.Get("/", ro.GetUnknownWordbook)
			r.Get("/entries", ro.GetUnknownWordbookEntries)
			r.Post("/entries", ro.AddUnknownWordbookEntry)
		})

		r.Route("/known", func(r chi.Router) {
			r.Get("/", ro.GetKnownWords)
			r.Post("/", ro.AddKnownWord)
		})
	})
}
