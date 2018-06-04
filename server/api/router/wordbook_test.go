package router_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/auth"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/middleware"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func TestWordbookGet(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	wordbooks := e.
		GET("/wordbook").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	wordbooks.Length().Equal(1)
	wordbook := wordbooks.Element(0).Object()
	wordbook.Keys().ContainsOnly("uuid", "name", "seen_date", "update_date")
	wordbook.Value("uuid").Equal(util.TestWordbookUuid)
	wordbook.Value("name").Equal("test")
}

func TestWordbookCreate(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	wordbook := util.M{
		"uuid":      "123e4567-e89b-12d3-a456-426655440000",
		"seen_date": "2018-06-04T11:03:49.859Z",
		"name":      "test2",
	}

	e.
		POST("/wordbook").
		WithHeader(middleware.ApiKeyHeader, key).
		WithJSON(wordbook).
		Expect().
		Status(201)

	wordbooks := e.
		GET("/wordbook").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	wordbooks.Length().Equal(2)
}

func TestWordbookPut(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	wordbook := util.M{
		"uuid":      util.TestWordbookUuid,
		"seen_date": "2018-06-04T11:03:49.859Z",
		"name":      "test3",
	}

	e.
		PUT("/wordbook/"+util.TestWordbookUuid).
		WithHeader(middleware.ApiKeyHeader, key).
		WithJSON(wordbook).
		Expect().
		Status(200)

	wordbooks := e.
		GET("/wordbook").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	wordbooks.Element(0).Object().Value("name").Equal("test3")
}

func TestWordbookDelete(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	e.
		DELETE("/wordbook/"+util.TestWordbookUuid).
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200)

	wordbooks := e.
		GET("/wordbook").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	wordbooks.Length().Equal(0)
}

func TestWordbookGetEntries(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	entries := e.
		GET("/wordbook/"+util.TestWordbookUuid+"/entries").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	entries.Length().Equal(1)
	entry := entries.Element(0).Object()
	entry.Keys().ContainsOnly("definition_id", "source_book", "source_sentence", "added_date", "word_index", "correct")
}

func TestWordbookPutEntries(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	entries := []dbh.WordbookEntry{
		dbh.WordbookEntry{
			DefinitionId: 1,
			AddedDate:    util.RFCTime{time.Now()},
		},
		dbh.WordbookEntry{
			DefinitionId: 1,
			AddedDate:    util.RFCTime{time.Now()},
		},
	}

	req := util.M{
		"entries": entries,
		"time":    util.RFCTime{time.Now()},
	}

	e.
		PUT("/wordbook/"+util.TestWordbookUuid+"/entries").
		WithJSON(req).
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200)

	entries2 := e.
		GET("/wordbook/"+util.TestWordbookUuid+"/entries").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array()

	entries2.Length().Equal(2)
}
