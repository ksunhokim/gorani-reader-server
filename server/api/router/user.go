package router

import (
	"encoding/json"
	"net/http"

	"github.com/sunho/gorani-reader/server/api/auth"
	"github.com/sunho/gorani-reader/server/api/models"
)

func (ro *Router) UserWithOauth(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Token   string `json:"token"`
		Service string `json:"service"`
	}{}

	json.NewDecoder(r.Body).Decode(&req)

	gorn := ro.gorn

	ouser, err := gorn.Services.FetchUser(req.Service, req.Token)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user, err := models.CreateOrGetUserWithOauth(gorn.Mysql, ouser)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	key, err := auth.ApiKeyByUser(gorn.Config.SecretKey, user.Id, user.Name)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp := struct {
		ApiKey string `json:"api_key"`
	}{
		ApiKey: key,
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
