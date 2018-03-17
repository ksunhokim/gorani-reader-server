package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunho/engbreaker/pkg/middlewares"
)

func (h *HTTPServer) registerRoutes() {
	h.route = mux.NewRouter()
	h.route.HandleFunc("/auth/{provider}", beginAuthHandler)
	h.route.HandleFunc("/auth/{provider}/callback", completeAuthHandler)
	h.route.Handle("/dummy", middlewares.Auth(http.HandlerFunc(dummy)))
}
