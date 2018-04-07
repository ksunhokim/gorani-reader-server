package view_test

import "testing"

func TestFetchBook(t *testing.T) {
	initDB()
	server, e := initServer(t)
	defer server.Close()

	obj := e.GET("/books/test").
		Expect().
		Status(200).
		JSON().
		Object()
	obj.Value("title").String().Equal("test")
	obj.Value("picture").String().Equal("test.png")
	obj.Value("view").Number().Equal(10)
	obj.Value("chapters").Array().First().String().Equal("hoi!호이")

	obj = e.GET("/books/test").
		Expect().
		Status(200).
		JSON().
		Object()
	obj.Value("view").Number().Equal(11)

	e.GET("/books/test2").
		Expect().
		Status(404)
}
func TestSearchBook(t *testing.T) {
	initDB()
	server, e := initServer(t)
	defer server.Close()

	obj := e.GET("/books/search?title=test&orderby=similarity").
		Expect().
		Status(200).
		JSON().
		Array()

	obj2 := obj.First().Object()
	obj2.Value("title").String().Equal("test")
	obj2.Value("picture").String().Equal("test.png")
	obj2.Value("view").Number().Equal(10)
	obj2.Value("chapters").Array().First().String().Equal("hoi!호이")
	obj.Length().Equal(3)

	obj = e.GET("/books/search?title=test").
		Expect().
		Status(200).
		JSON().
		Array()

	obj2 = obj.First().Object()
	obj2.Value("title").String().Equal("test")
	obj2.Value("picture").String().Equal("test.png")
	obj2.Value("view").Number().Equal(10)
	obj2.Value("chapters").Array().First().String().Equal("hoi!호이")
	obj.Length().Equal(3)

	obj = e.GET("books/search?author=호잇&orderby=abc").
		Expect().
		Status(200).
		JSON().
		Array()
	obj.First().Object().Value("title").String().Equal("테스트")
	obj.Element(1).Object().Value("title").String().Equal("테스트2")

	obj = e.GET("books/search?title=test&orderby=recent").
		Expect().
		Status(200).
		JSON().
		Array()
	obj.First().Object().Value("title").String().Equal("test0")

	e.GET("books/search?title=a&author=b").
		Expect().
		Status(400)

	e.GET("books/search").
		Expect().
		Status(400)

	e.GET("books/search&orderby=abc").
		Expect().
		Status(400)

	e.GET("books/search&title=").
		Expect().
		Status(400)

	e.GET("books/search&title=asdffadsf").
		Expect().
		Status(404)
}

func TestFetchBookContent(t *testing.T) {

}
