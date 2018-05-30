package router

import (
	"encoding/json"
	"net/http"

	"github.com/sunho/gorani-reader/server/pkg/auth"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func (ro *Router) UserWithOauth(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Token   string `json:"token"`
		Service string `json:"service"`
	}{}

	json.NewDecoder(r.Body).Decode(&req)

	ap := ro.ap

	ouser, err := ap.Services.FetchUser(req.Service, req.Token)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := dbh.CreateOrGetUserWithOauth(ap.Gorn.Mysql, ouser)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, user.Id, user.Name)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp := util.M{
		"api_key": key,
	}

	util.JSON(w, resp)
}
