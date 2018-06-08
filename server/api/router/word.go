package router

import (
	"encoding/json"
	"net/http"

	"github.com/sunho/gorani-reader/server/pkg/middleware"
)

var uwordbookContextKey = contextKey{"uwordbook"}

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
