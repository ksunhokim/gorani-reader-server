package api

import (
	"github.com/sunho/gorani-reader-server/pkg/auth"
	"github.com/sunho/gorani-reader-server/pkg/dbh"
)

func (ap *Api) RegisterUser(service string, token string) (string, error) {
	ouser, err := ap.Services.FetchUser(service, token)
	if err != nil {
		return "", err
	}

	user, err := dbh.CreateOrGetUserWithOauth(ap.Mysql, ouser)
	if err != nil {
		return "", err
	}

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, user.Id, user.Name)
	if err != nil {
		return "", err
	}

	return key, nil
}
