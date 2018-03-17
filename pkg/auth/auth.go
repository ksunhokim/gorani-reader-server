package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/markbates/goth"
	"github.com/sirupsen/logrus"
	"github.com/sunho/engbreaker/pkg/dbs"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func state() string {
	nonceBytes := make([]byte, 64)
	for i := 0; i < 64; i++ {
		nonceBytes[i] = byte(random.Int63() % 256)
	}
	return base64.URLEncoding.EncodeToString(nonceBytes)
}
func checkState(s string) bool {
	b, err := dbs.RDB.Exists("authstate:" + s).Result()
	if err != nil {
		logrus.Error(err)
	}
	return b == 1
}
func GenerateAuthSession(provider goth.Provider) (goth.Session, error) {
	st := state()
	for checkState(st) {
		st = state()
	}

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

	s, err := GenerateAuthSession(p)
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

	payload, err := dbs.RDB.Get("authstate:" + url.Query().Get("state")).Result()
	if err != nil {
		return goth.User{}, err
	}

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
