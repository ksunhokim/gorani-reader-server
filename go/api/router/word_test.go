package router_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/pkg/auth"
	"github.com/sunho/gorani-reader-server/pkg/middleware"
	"github.com/sunho/gorani-reader-server/pkg/util"
)

func TestKnownWord(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	entry := util.M{
		"word_ids": []int{
			1,
		},
	}

	e.
		POST("/word/known").
		WithHeader(middleware.ApiKeyHeader, key).
		WithJSON(entry).
		Expect().
		Status(201)
}

func TestUnknownWord(t *testing.T) {
	a := assert.New(t)
	e, s, ap := prepareServer(t)
	defer s.Close()

	key, err := auth.ApiKeyByUser(ap.Config.SecretKey, util.TestUserId, "test")
	a.Nil(err)

	entry := util.M{
		"added_date":      "2018-06-09T12:29:47Z",
		"memory_sentence": "string",
		"sources": []util.M{
			util.M{
				"definition_id":     1,
				"source_book":       "string",
				"source_sentence":   "string",
				"source_word_index": 0,
			},
		},
	}

	e.
		PUT("/word/unknown/1").
		WithHeader(middleware.ApiKeyHeader, key).
		WithJSON(entry).
		Expect().
		Status(200)

	obj := e.
		GET("/word/unknown").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array().
		Element(0).
		Object()

	obj.Keys().ContainsOnly("added_date", "word_id", "sources", "memory_sentence")
	obj.Value("added_date").Equal(entry["added_date"])
	arr := obj.Value("sources").Array()
	arr.Length().Equal(1)
	obj2 := arr.Element(0).Object()
	obj2.Keys().ContainsOnly("definition_id", "source_book", "source_sentence", "source_word_index")
	obj2.Value("definition_id").Equal(1)

	e.
		DELETE("/word/unknown/1").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200)

	e.
		GET("/word/unknown").
		WithHeader(middleware.ApiKeyHeader, key).
		Expect().
		Status(200).
		JSON().
		Array().
		Length().
		Equal(0)
}
