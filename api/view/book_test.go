package view_test

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFetchBook(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	obj := e.GET("/books/test").
		WithHeader("Authorization", token).
		Expect().
		Status(200).
		JSON().
		Object()
	obj.Value("title").String().Equal("test")
	obj.Value("picture").String().Equal("test.png")
	obj.Value("view").Number().Equal(10)
	obj.Value("chapters").Array().First().String().Equal("hoi!호이")

	obj = e.GET("/books/test").
		WithHeader("Authorization", token).
		Expect().
		Status(200).
		JSON().
		Object()
	obj.Value("view").Number().Equal(11)

	e.GET("/books/test2").
		Expect().
		Status(404)
}
func TestFetchBookContent(t *testing.T) {
	initDB()
	server, e := initServer(t)
	defer server.Close()

	obj := e.GET("/books/test/chpaters/0").
		Expect().
		Status(200).
		JSON().
		Object()

	obj.Value("content").String().Equal("<div>호이</div>")
}
func TestBookProgress(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	e.GET("/books/test/progress").
		WithHeader("Authorization", token).
		Expect().
		Status(404)

	e.PUT("/books/test/progress").
		WithHeader("Authorization", token).
		WithJSON(
			gin.H{
				"chapter":  0,
				"sentence": 1,
			},
		).
		Expect().
		Status(200)

	obj := e.GET("/books/test/progress").
		WithHeader("Authorization", token).
		Expect().
		Status(200).
		JSON().
		Object()

	obj.Value("chapter").Number().Equal(0)
	obj.Value("sentence").Number().Equal(0)

	e.PUT("/books/test/progress").
		WithHeader("Authorization", token).
		WithJSON(
			gin.H{
				"chapter":  1,
				"sentence": 1,
			},
		).
		Expect().
		Status(400)
}

func TestMyBook(t *testing.T) {
	token := initDB()
	server, e := initServer(t)
	defer server.Close()

	e.PUT("/books/test/progress").
		WithHeader("Authorization", token).
		WithJSON(
			gin.H{
				"chapter":  0,
				"sentence": 1,
			},
		).
		Expect().
		Status(200)

	obj := e.GET("/books/my").
		Expect().
		Status(200).
		JSON().
		Array().
		First().
		Object()

	obj.Value("title").String().Equal("test")
	obj.Value("picture").String().Equal("test.png")
	obj.Value("view").Number().Equal(10)
	obj2 := obj.Value("progress").Object()
	obj2.Value("chapter").Number().Equal(0)
	obj2.Value("senetence").Number().Equal(1)
	obj2.Value("chpatertitle").String().Equal("hoi!호이")
}
