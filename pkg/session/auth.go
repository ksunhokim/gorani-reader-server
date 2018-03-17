package session

import (
	"encoding/base64"
	"fmt"
	"math/rand"

	"github.com/sunho/goth"
)

var random *rand.Rand

func state() string {
	nonceBytes := make([]byte, 64)
	for i := 0; i < 64; i++ {
		nonceBytes[i] = byte(random.Int63() % 256)
	}
	return base64.URLEncoding.EncodeToString(nonceBytes)
}

func GenerateAuthSession(provider goth.Provider) (goth.Session, error) {
	s, err := provider.BeginAuth(state())
	if err != nil {
		return nil, err
	}

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

func StoreAuthSession(sess goth.Session) {

}
