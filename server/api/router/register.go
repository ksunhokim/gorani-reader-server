package router

import (
	"fmt"

	"github.com/go-chi/chi"
	chimid "github.com/go-chi/chi/middleware"
	"github.com/sunho/gorani-reader/server/pkg/auth"
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
			r.Post("/", ro.AddKnownWord)
		})
	})
	fmt.Println(auth.ApiKeyByUser(ro.ap.Config.SecretKey, 1, "asdf"))
}
