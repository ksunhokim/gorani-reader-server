package router_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/auth"
	"github.com/sunho/gorani-reader/server/pkg/middleware"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func TestKnownWord(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	entry := util.M{
		"word_id": 1,
	}

	e.
		POST("/word/known").
		WithHeader(middleware.ApiKeyHeader, key).
		WithJSON(entry).
		Expect().
		Status(200)
}
