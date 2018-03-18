package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/middlewares"
)

func (h *HTTPServer) registerRoutes() {
	h.route = mux.NewRouter()

	auth := h.route.PathPrefix("/auth").Subrouter()
	auth.
		Handle("/refresh", middlewares.Auth(refreshAuthHandler)).
		Methods("GET")
	auth.
		Handle("/{provider}", beginAuthHandler).
		Methods("GET")
	auth.
		Handle("/{provider}/callback", completeAuthHandler).
		Methods("GET")

	api := h.route.PathPrefix("/api").Subrouter()
	api.
		Handle("/wordbooks", middlewares.Auth(wordBookListHandler)).
		Methods("GET")
	api.
		Handle("/wordbooks/{name}", middlewares.Auth(wordBookDetailHandler)).
		Methods("GET")
	api.
		Handle("/wordbooks/{name}", middlewares.Auth(wordBookAddHandler)).
		Methods("POST")
	api.
		Handle("/wordbooks/{name}", middlewares.Auth(wordBookRemoveHandler)).
		Methods("DELETE")

	api.
		Handle("/words/{word}", wordsHandler).
		Methods("GET")
	api.
		Handle("/words/{word}/{index}", wordHandler).
		Methods("GET")

	resDir := config.GetString("RESOURCE", "../../public/dist/")
	h.route.PathPrefix("/resource/").
		Handler(http.StripPrefix("/resource/", http.FileServer(http.Dir(resDir)))).
		Methods("GET")

	h.route.PathPrefix("/").
		Handler(indexHandler).
		Methods("GET")
}
