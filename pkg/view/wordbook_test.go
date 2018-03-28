package view_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/antonholmquist/jason"
	"github.com/stretchr/testify/assert"
)

func TestWithoutUser(t *testing.T) {
	initDB()
	a := assert.New(t)

	req, _ := http.NewRequest("GET", "/wordbooks", nil)
	w := testEndpoint("dummy", req)
	a.Equal(401, w.Code)

	req, _ = http.NewRequest("POST", "/wordbooks/asd", nil)
	w = testEndpoint("dummy", req)
	a.Equal(401, w.Code)
}

func TestAddWordBook(t *testing.T) {
	token := initDB()
	a := assert.New(t)

	req, _ := http.NewRequest("GET", "/wordbooks", nil)
	w := testEndpoint(token, req)
	obj, _ := jason.NewObjectFromReader(w.Body)
	books, _ := obj.GetStringArray("wordbooks")

	a.Equal(200, w.Code)
	a.Equal(1, len(books))
	a.Equal("test", books[0])
}

func TestPostWordBook(t *testing.T) {
	token := initDB()
	a := assert.New(t)

	req, _ := http.NewRequest("POST", "/wordbooks/asd", nil)
	w := testEndpoint(token, req)
	a.Equal(201, w.Code)

	w = testEndpoint(token, req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("GET", "/wordbooks", nil)
	w = testEndpoint(token, req)
	obj, _ := jason.NewObjectFromReader(w.Body)
	books, _ := obj.GetStringArray("wordbooks")

	a.Equal(200, w.Code)
	a.Equal(2, len(books))
	a.Equal("asd", books[0])
}

func TestGetWordBook(t *testing.T) {
	token := initDB()
	a := assert.New(t)

	req, _ := http.NewRequest("GET", "/wordbooks/test", nil)
	w := testEndpoint(token, req)
	obj, _ := jason.NewObjectFromReader(w.Body)
	_, err := obj.GetString("_id")
	name, _ := obj.GetString("name")
	entries, _ := obj.GetObjectArray("entries")
	firstStar, _ := entries[0].GetBoolean("star")
	created, _ := obj.GetString("created")
	modified, _ := obj.GetString("modified")

	a.Equal(200, w.Code)
	a.NotEqual(nil, err)
	a.Equal("test", name)
	a.Equal(1, len(entries))
	a.Equal(true, firstStar)
	a.NotEqual("", created)
	a.NotEqual("", modified)

	req, _ = http.NewRequest("GET", "/wordbooks/test2", nil)
	w = testEndpoint(token, req)
	a.Equal(404, w.Code)
}

func TestAddEntryWordBook(t *testing.T) {
	token := initDB()
	a := assert.New(t)

	req, _ := http.NewRequest("POST", "/wordbooks/test/words", strings.NewReader(`{"words":[{"word":"test","definition":1,"book":"test"}]`))
	w := testEndpoint(token, req)
	a.Equal(200, w.Code)

	w = testEndpoint(token, req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("POST", "/wordbooks/test/words", strings.NewReader(`{"words":[{"word":"test","definition":1,"book":"test2"}]}`))
	w = testEndpoint(token, req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("POST", "/wordbooks/test/words", strings.NewReader(`{"words":[{"word":"test","definition":2,"book":"test"}]}`))
	w = testEndpoint(token, req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("POST", "/wordbooks/test/words", strings.NewReader(`{"words":[{"word":"test","book":"test2"}]}`))
	w = testEndpoint(token, req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("POST", "/wordbooks/test/words", strings.NewReader(`{"words":]}`))
	w = testEndpoint(token, req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("GET", "/wordbooks/test", nil)
	w = testEndpoint(token, req)
	obj, _ := jason.NewObjectFromReader(w.Body)
	entries, _ := obj.GetObjectArray("entries")

	a.Equal(200, w.Code)
	a.Equal(2, len(entries))
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
