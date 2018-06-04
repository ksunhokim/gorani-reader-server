package router

import (
	"context"
	"net/http"

	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/middleware"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

var uwordbookContextKey = contextKey{"uwordbook"}

func (ro *Router) UnknownWordbookCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := middleware.GetUser(r)
		wb, err := user.GetUnknownWordbook(ro.ap.Mysql)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), uwordbookContextKey, wb)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ro *Router) GetUnknownWordbook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	wb := ctx.Value(uwordbookContextKey).(dbh.Wordbook)
	util.JSON(w, wb)
}
