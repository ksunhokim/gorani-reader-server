package router_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/pkg/auth"
	"github.com/sunho/gorani-reader-server/pkg/middleware"
	"github.com/sunho/gorani-reader-server/pkg/util"
)

func TestGenre(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	arr := e.
		GET("/genre").
		Expect().
		Status(200).
		JSON().
		Array()

	arr.Length().Equal(1)
	arr.Element(0).String().Equal("test")

	entries := []string{"test"}

	e.
		PUT("/genre/prefer").
		WithHeader(middleware.ApiKeyHeader, key).
		WithJSON(entries).
		Expect().
		Status(200)

	arr = e.
		GET("/genre/prefer").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	arr.Length().Equal(1)
	arr.Element(0).String().Equal("test")
}
