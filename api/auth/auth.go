package auth

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/markbates/goth"
	"github.com/sunho/engbreaker/pkg/dbs"
)

func generateAuthSession(provider goth.Provider) (goth.Session, error) {
	st := dbs.GenerateRedisNonce("authstate:")

	s, err := provider.BeginAuth(st)
	if err != nil {
		return nil, err
	}
	payload, _ := json.Marshal(s)

	err = dbs.RDB.Set("authstate:"+st, payload, time.Minute*10).Err()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func GetAuthURL(provider string) (string, error) {
	if provider == "" {
		return "", fmt.Errorf("provider is empty")
	}

	p, err := goth.GetProvider(provider)
	if err != nil {
		return "", err
	}

	s, err := generateAuthSession(p)
	if err != nil {
		return "", err
	}

	url, err := s.GetAuthURL()
	if err != nil {
		return "", err
	}

	return url, nil
}

func CompleteAuth(provider string, url *url.URL) (goth.User, error) {
	if provider == "" {
		return goth.User{}, fmt.Errorf("provider is empty")
	}

	p, err := goth.GetProvider(provider)
	if err != nil {
		return goth.User{}, err
	}

	name := "authstate:" + url.Query().Get("state")
	payload, err := dbs.RDB.Get(name).Result()
	if err != nil {
		return goth.User{}, err
	}

	dbs.RDB.Del(name)

	sess, err := p.UnmarshalSession(payload)
	if err != nil {
		return goth.User{}, err
	}

	_, err = sess.Authorize(p, url.Query())
	if err != nil {
		return goth.User{}, err
	}

	user, err := p.FetchUser(sess)
	if err != nil {
		return goth.User{}, err
	}

	return user, nil
}
