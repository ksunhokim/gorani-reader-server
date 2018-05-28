package router

import (
	chimid "github.com/go-chi/chi/middleware"
	mymid "github.com/sunho/gorani-reader/server/pkg/middleware"
)

func (ro *Router) registerHandlers() {
	ro.Use(chimid.RealIP)
	ro.Use(mymid.RequestId)
	ro.Use(mymid.Logger)

}
