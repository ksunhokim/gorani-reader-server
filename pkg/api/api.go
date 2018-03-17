package api

import (
	"github.com/gorilla/mux"
	"github.com/sunho/engbreaker/pkg/middlewares"
)

func (h *HTTPServer) registerRoutes() {
	h.route = mux.NewRouter()
	h.route.Handle("/auth/{provider}", beginAuthHandler)
	h.route.Handle("/auth/{provider}/callback", completeAuthHandler)
	h.route.Handle("/dummy", middlewares.Auth(dummy))
	h.route.Handle("/", middlewares.AuthOrRedirect(dummy, "/login"))
}
