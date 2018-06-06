package services

import (
	"github.com/sunho/gorani-reader/server/pkg/auth"
)

var (
	naverService = auth.Service{
		Name:             "naver",
		BaseUrl:          "https://openapi.naver.com",
		UserEndPoint:     "/v1/nid/me",
		UsernameSelector: "response.nickname",
		AvatorSelector:   "response.profile_image",
		IdSelector:       "response.id",
	}
)

func New() (srv auth.Services) {
	srv.AddService(naverService)
	return
}
