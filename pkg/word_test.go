package view_test

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/antonholmquist/jason"
	"github.com/stretchr/testify/assert"
	"github.com/sunho/engbreaker/pkg/model"
)

func TestWordGet(t *testing.T) {
	initDB()
	a := assert.New(t)

	req, _ := http.NewRequest("GET", "/words/test", nil)
	w := testEndpoint("dummy", req)
	obj := model.Word{}
	json.NewDecoder(w.Body).Decode(&obj)

	a.Equal(200, w.Code)
	a.Equal("test", obj.Word)
	a.Equal("test", obj.Pronunciation)
	a.Equal(2, len(obj.Definitions))
	a.Equal("hello", obj.Definitions[0].Definition)
	a.Equal("안녕", obj.Definitions[0].Examples[0].Second)

	req, _ = http.NewRequest("GET", "/words/test/0", nil)
	w = testEndpoint("dummy", req)
	obj2, _ := jason.NewObjectFromReader(w.Body)
	word2, _ := obj2.GetString("word")
	pron, _ := obj2.GetString("pronunciation")
	def, _ := obj2.GetString("definition")
	examples, _ := obj2.GetObjectArray("examples")
	exampleSecond, _ := examples[0].GetString("second")

	a.Equal(200, w.Code)
	a.Equal("test", word2)
	a.Equal("test", pron)
	a.Equal("hello", def)
	a.Equal(1, len(examples))
	a.Equal("안녕", exampleSecond)

	req, _ = http.NewRequest("GET", "/words/test2", nil)
	w = testEndpoint("dummy", req)
	a.Equal(404, w.Code)

	req, _ = http.NewRequest("GET", "/words/", nil)
	w = testEndpoint("dummy", req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("GET", "/words/test2/0", nil)
	w = testEndpoint("dummy", req)
	a.Equal(404, w.Code)

	req, _ = http.NewRequest("GET", "/words/언오랴나/0", nil)
	w = testEndpoint("dummy", req)
	a.Equal(404, w.Code)

	req, _ = http.NewRequest("GET", "/words/test/40123", nil)
	w = testEndpoint("dummy", req)
	a.Equal(404, w.Code)

	maxUint := ^uint64(0)
	maxi := strconv.FormatUint(maxUint, 10)
	req, _ = http.NewRequest("get", "/words/test/"+maxi, nil)
	w = testEndpoint("dummy", req)
	a.Equal(404, w.Code)

	req, _ = http.NewRequest("get", "/words/test/"+maxi+"9", nil)
	w = testEndpoint("dummy", req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("get", "/words/test/-1", nil)
	w = testEndpoint("dummy", req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("get", "/words/test/호이!", nil)
	w = testEndpoint("dummy", req)
	a.Equal(400, w.Code)

	req, _ = http.NewRequest("get", "/words/test/!/dsd/", nil)
	w = testEndpoint("dummy", req)
	a.Equal(404, w.Code)
}

func TestWordSearch(t *testing.T) {
	initWordDB()
	a := assert.New(t)

	req, _ := http.NewRequest("GET", "/words?q=go", nil)
	w := testEndpoint("dummy", req)
	obj := model.Word{}
	json.NewDecoder(w.Body).Decode(&obj)

	a.Equal(200, w.Code)
	a.Equal("go", obj.Word)

	req, _ = http.NewRequest("GET", "/words?q=went", nil)
	w = testEndpoint("dummy", req)
	obj = model.Word{}
	json.NewDecoder(w.Body).Decode(&obj)

	a.Equal(200, w.Code)
	a.Equal("go", obj.Word)
}
