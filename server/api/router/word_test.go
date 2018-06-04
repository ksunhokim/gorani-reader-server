package router_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/auth"
	"github.com/sunho/gorani-reader/server/pkg/middleware"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func TestUnknownWordbookGet(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	wordbook := e.
		GET("/word/unknown").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Object()

	wordbook.Keys().ContainsOnly("uuid", "name", "seen_date", "update_date")
	wordbook.Value("uuid").Equal(util.TestUnknownWordbookUuid)
}

func TestUnknownWordbookGetEntries(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	wordbooks := e.
		GET("/word/unknown/entries").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	wordbooks.Length().Equal(1)
	entry := wordbooks.Element(0).Object()
	entry.Keys().ContainsOnly("definition_id", "source_book", "source_sentence", "added_date", "word_index", "correct")
}
