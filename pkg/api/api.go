package api

import (
	"github.com/gorilla/mux"
	"github.com/sunho/engbreaker/pkg/middlewares"
)

func (h *HTTPServer) registerRoutes() {
	h.route = mux.NewRouter()
	h.route.
		Handle("/auth/{provider}", beginAuthHandler).
		Methods("GET")
	h.route.
		Handle("/auth/{provider}/callback", completeAuthHandler).
		Methods("GET")

	h.route.
		Handle("/wordbooks", middlewares.Auth(wordBookListHandler)).
		Methods("GET")
	h.route.
		Handle("/wordbooks/{name}", middlewares.Auth(wordBookDetailHandler)).
		Methods("GET")
	h.route.
		Handle("/wordbooks/{name}", middlewares.Auth(wordBookAddHandler)).
		Methods("POST")
	h.route.
		Handle("/wordbooks/{name}", middlewares.Auth(wordBookRemoveHandler)).
		Methods("DELETE")

	h.route.
		Handle("/", middlewares.AuthOrRedirect(wordBookListHandler, "/login")).
		Methods("GET")
}
