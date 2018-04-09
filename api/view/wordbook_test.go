package view_test

import (
	"testing"

	"github.com/gin-gonic/gin"
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
	obj.First().Object().Value("name").String().Equal("test")
	obj.First().Object().Value("entries").Number().Equal(1)
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
	obj.First().Object().Value("name").String().Equal("asd")
	obj.First().Object().Value("entries").Number().Equal(1)
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
	obj.Value("entries").Array().First().Object().Value("definition_text").String().Equal("hello")

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

	obj := e.GET("/wordbooks/test").
		WithHeader("Authorization", token).
		Expect().
		Status(200).
		JSON().Object()
	obj.Value("entries").Array().Length().Equal(2)
}

func TestDeleteWordBook(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	e.DELETE("/wordbooks/test").
		WithHeader("Authorization", token).
		Expect().
		Status(200)

	e.DELETE("/wordbooks/test").
		WithHeader("Authorization", token).
		Expect().
		Status(404)

	e.GET("/wordbooks/test").
		WithHeader("Authorization", token).
		Expect().
		Status(404)

	obj := e.GET("/wordbooks").
		WithHeader("Authorization", token).
		Expect().
		Status(200).
		JSON().Array()
	obj.Length().Equal(0)
}

func TestPutEntryWordBook(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	e.PUT("/wordbooks/test/words").
		WithHeader("Authorization", token).
		WithJSON([]gin.H{
			gin.H{
				"word":       "test",
				"definition": 0,
				"book":       "test",
			},
			gin.H{
				"word":       "test",
				"definition": 1,
				"book":       "test",
			},
		}).
		Expect().
		Status(200)
}
