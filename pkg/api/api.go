package api

import (
	"github.com/gorilla/mux"
)

func (h *HTTPServer) registerRoutes() {
	h.route = mux.NewRouter()
	
}
