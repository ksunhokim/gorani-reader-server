package router_test

import (
	"testing"
)

func TestGenre(t *testing.T) {
	// a := assert.New(t)
	e, s, _ := prepareServer(t)
	defer s.Close()

	// key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	// a.Nil(err)

	arr := e.
		GET("/genre").
		Expect().
		Status(200).
		JSON().
		Array()

	arr.Length().Equal(1)
	arr.Element(0).String().Equal("test")
}
