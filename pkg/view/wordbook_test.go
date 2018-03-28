package view_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/antonholmquist/jason"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestWithoutUser(t *testing.T) {
	initDB()
	server, e := initServer(t)
	defer server.Close()

	e.GET("/wordbooks").
		Expect().
		Status(401)

	e.GET("/wordbooks/asd").
		Expect().
		Status(401)
}

func TestAddWordBook(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	obj := e.GET("/wordbooks").
		WithHeader("Authorization", token).
		Expect().
		Status(200).
		JSON().Array()

	obj.Length().Equal(1)
	obj.First().String().Equal("test")
}

func TestPostWordBook(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	e.POST("/wordbooks/asd").
		WithHeader("Authorization", token).
		Expect().
		Status(201)

	e.POST("/wordbooks/asd").
		WithHeader("Authorization", token).
		Expect().
		Status(400)

	obj := e.GET("/wordbooks").
		WithHeader("Authorization", token).
		Expect().
		Status(200).
		JSON().Array()
	obj.Length().Equal(2)
	obj.First().String().Equal("asd")
}

func TestGetWordBook(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	obj := e.GET("/wordbooks/test").
		WithHeader("Authorization", token).
		Expect().
		Status(200).
		JSON().Object()
	obj.Keys().NotContains("_id", "_created", "_modified").
		Contains("name", "entries", "created", "modified")
	obj.Value("name").String().Equal("test")
	obj.Value("entries").Array().First().Object().Value("star").Boolean().Equal(true)

	e.GET("/wordbooks/test2").
		WithHeader("Authorization", token).
		Expect().
		Status(404)
}

func TestAddEntryWordBook(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	e.POST("/wordbooks/test/words").
		WithHeader("Authorization", token).
		WithJSON([]gin.H{
			gin.H{
				"word":       "test",
				"definition": 1,
				"book":       "test",
			},
		}).
		Expect().
		Status(201)

	e.POST("/wordbooks/test/words").
		WithHeader("Authorization", token).
		WithJSON([]gin.H{
			gin.H{
				"word":       "test",
				"definition": 1,
				"book":       "test",
			},
		}).
		Expect().
		Status(400)

	e.POST("/wordbooks/test/words").
		WithHeader("Authorization", token).
		WithJSON([]gin.H{
			gin.H{
				"word":       "test",
				"definition": 1,
				"book":       "test2",
			},
		}).
		Expect().
		Status(400)

	e.POST("/wordbooks/test/words").
		WithHeader("Authorization", token).
		WithJSON([]gin.H{
			gin.H{
				"word":       "test",
				"definition": 2,
				"book":       "test",
			},
		}).
		Expect().
		Status(400)

	e.POST("/wordbooks/test/words").
		WithHeader("Authorization", token).
		WithJSON([]gin.H{
			gin.H{
				"word":       "test",
				"definition": 2,
				"book":       "test",
			},
		}).
		Expect().
		Status(400)

	e.POST("/wordbooks/test/words").
		WithHeader("Authorization", token).
		WithJSON([]gin.H{
			gin.H{
				"word": "test",
				"book": "test",
			},
		}).
		Expect().
		Status(400)

	e.POST("/wordbooks/test/words").
		WithHeader("Authorization", token).
		WithText(`{"words":}`).
		Expect().
		Status(400)

	obj := e.GET("/wordbooks/test/words").
		WithHeader("Authorization", token).
		Expect().
		Status(200).
		JSON().Object()
	obj.Value("entries").Array().Length().Equal(2)

}

func TestDeleteWordBook(t *testing.T) {
	token := initDB()
	a := assert.New(t)

	req, _ := http.NewRequest("DELETE", "/wordbooks/test", nil)
	w := testEndpoint(token, req)
	a.Equal(200, w.Code)
	w = testEndpoint(token, req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("GET", "/wordbooks/test", nil)
	w = testEndpoint(token, req)
	a.Equal(404, w.Code)

	req, _ = http.NewRequest("GET", "/wordbooks", nil)
	w = testEndpoint(token, req)
	obj, _ := jason.NewObjectFromReader(w.Body)
	entries, _ := obj.GetObjectArray("wordbooks")

	a.Equal(200, w.Code)
	a.Equal(0, len(entries))
}

func TestPutEntryWordBook(t *testing.T) {
	token := initDB()
	a := assert.New(t)

	req, _ := http.NewRequest("PUT", "/wordbooks/test/words", strings.NewReader(`{"words":[{"word":"test","definition":1,"book":"test"},{"word":"test","definition":0,"book":"test"}]}`))
	w := testEndpoint(token, req)
	a.Equal(200, w.Code)

	req, _ = http.NewRequest("GET", "/wordbooks/test", nil)
	w = testEndpoint(token, req)
	obj, _ := jason.NewObjectFromReader(w.Body)
	entries, _ := obj.GetObjectArray("entries")
	defOne, _ := entries[0].GetInt64("definition")
	defTwo, _ := entries[1].GetInt64("definition")

	a.Equal(200, w.Code)
	a.Equal(1, defOne)
	a.Equal(0, defTwo)

	req, _ = http.NewRequest("PUT", "/wordbooks/test/words", strings.NewReader(`[{"word":"test","definition":10,"book":"test"},{"word":"test","definition":0,"book":"test"}]`))
	w = testEndpoint(token, req)
	a.Equal(400, w.Code)
}
