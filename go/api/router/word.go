package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sunho/gorani-reader-server/go/pkg/dbh"
	"github.com/sunho/gorani-reader-server/go/pkg/middleware"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
)

var uwordbookContextKey = contextKey{"uwordbook"}

type KnownWordsRequest struct {
	WordIds []int `json:"word_ids"`
}

func (ro *Router) AddKnownWords(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)

	req := KnownWordsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	for _, id := range req.WordIds {
		err = user.AddKnownWord(ro.ap.Mysql, id)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	w.WriteHeader(201)
}

func (ro *Router) GetUnknownWords(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)

	words, err := user.GetUnknownWordWithQuizs(ro.ap.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	util.JSON(w, words)
}

func (ro *Router) PutUnknownWord(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)
	raw := chi.URLParam(r, "word_id")
	wordid, _ := strconv.Atoi(raw)

	word := dbh.UnknownWord{}
	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	word.WordId = wordid

	err = user.PutUnknownWord(ro.ap.Mysql, &word)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (ro *Router) DeleteUnknownWord(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)

	raw := chi.URLParam(r, "word_id")
	wordid, _ := strconv.Atoi(raw)

	word, err := user.GetUnknownWord(ro.ap.Mysql, wordid)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = word.Delete(ro.ap.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
