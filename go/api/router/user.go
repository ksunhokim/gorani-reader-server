package router

import (
	"encoding/json"
	"net/http"

	"github.com/sunho/gorani-reader-server/go/pkg/util"
)

func (ro *Router) UserWithOauth(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Token   string `json:"token"`
		Service string `json:"service"`
	}{}

	json.NewDecoder(r.Body).Decode(&req)

	key, err := ro.ap.RegisterUser(req.Service, req.Token)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp := util.M{
		"api_key": key,
	}

	util.JSON(w, resp)
}
