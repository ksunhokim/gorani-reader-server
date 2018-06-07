package router

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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

func (ro *Router) GetUnknownWordbookEntries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	wb := ctx.Value(uwordbookContextKey).(dbh.Wordbook)
	entries, err := wb.GetEntries(ro.ap.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	util.JSON(w, entries)
}

func (ro *Router) AddUnknownWordbookEntry(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entry := dbh.WordbookEntry{}
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	wb := ctx.Value(uwordbookContextKey).(dbh.Wordbook)
	err = wb.AddEntry(ro.ap.Mysql, time.Now().UTC(), &entry)
	if err != nil {
		http.Error(w, err.Error(), 409)
		return
	}

	w.WriteHeader(200)
}

type KnwonWordRequest struct {
	WordId int `json:"word_id"`
}

func (ro *Router) AddKnownWord(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)

	req := KnwonWordRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = user.AddKnownWord(ro.ap.Mysql, req.WordId)
	if err != nil {
		http.Error(w, err.Error(), 409)
		return
	}

	w.WriteHeader(200)
}

type KnwonWordResponse struct {
	WordIds []int `json:"word_ids"`
}

func (ro *Router) GetKnownWords(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)

	wids, err := user.GetKnownWords(ro.ap.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	resp := KnwonWordResponse{}
	resp.WordIds = wids

	util.JSON(w, resp)
}
