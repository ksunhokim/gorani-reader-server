package util

import (
	"net/http"
	"time"
)

func CreateClient() http.Client {
	return http.Client{
		Timeout: time.Second * 10,
	}
}
