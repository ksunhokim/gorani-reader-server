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

	entries := e.
		GET("/word/unknown/entries").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	entries.Length().Equal(1)
	entry := entries.Element(0).Object()
	entry.Keys().ContainsOnly("definition_id", "source_book", "source_sentence", "added_date", "word_index", "correct")
}

func TestUnknownWordbookAddEntry(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	entry := util.M{
		"definition_id": 1,
		"added_date":    "2018-06-04T11:03:49.859Z",
	}

	e.
		POST("/word/unknown/entries").
		WithHeader(middleware.ApiKeyHeader, key).
		WithJSON(entry).
		Expect().
		Status(200)

	entries := e.
		GET("/word/unknown/entries").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	entries.Length().Equal(2)
}

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

	entries := e.
		GET("/word/known").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Object().
		Value("word_ids").
		Array()

	entries.Length().Equal(1)
	entries.Element(0).Equal(1)
}
