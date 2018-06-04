package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/middleware"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

var wordbookContextKey = contextKey{"wordbook"}

func (ro *Router) WordbookCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "uuid")
		u, err := uuid.Parse(uid)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		user := middleware.GetUser(r)
		wb, err := user.GetWordbook(ro.ap.Mysql, dbh.UUID{u})
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), wordbookContextKey, wb)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ro *Router) GetWordbook(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)
	wordbooks, err := user.GetWordbooks(ro.ap.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	util.JSON(w, wordbooks)
}

func (ro *Router) AddWordbook(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)

	wordbook := dbh.Wordbook{}
	err := json.NewDecoder(r.Body).Decode(&wordbook)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = user.AddWordbook(ro.ap.Mysql, &wordbook)
	if err != nil {
		http.Error(w, err.Error(), 409)
		return
	}

	w.WriteHeader(201)
}

func (ro *Router) DeleteWordbook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	wb := ctx.Value(wordbookContextKey).(dbh.Wordbook)
	err := wb.Delete(ro.ap.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}

func (ro *Router) GetWordbookEntries(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// wb := ctx.Value(wordbookContextKey).(dbh.Wordbook)
	// entries, err := wb.GetEntries(ro.ap.Mysql)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

}
