package router

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

var wordContextKey = contextKey{"word"}

func (ro *Router) WordCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		word, err := dbh.GetWord(ro.gorn.Mysql, id)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), wordContextKey, word)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ro *Router) GetWords(w http.ResponseWriter, r *http.Request) {
	words, err := dbh.GetWords(ro.gorn.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	util.JSON(w, words)
}

func (ro *Router) GetWord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	word := ctx.Value(wordContextKey).(dbh.Word)

	defs, err := word.GetDefinitions(ro.gorn.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for i := range defs {
		examples, err := defs[i].GetExamples(ro.gorn.Mysql)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defs[i].Examples = examples
	}

	word.Definitions = defs

	util.JSON(w, word)
}

func (ro *Router) DeleteWord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	word := ctx.Value(wordContextKey).(dbh.Word)
	err := word.Delete(ro.gorn.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}
