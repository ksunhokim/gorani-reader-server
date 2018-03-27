package view_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/engbreaker/pkg/model"
	"github.com/sunho/engbreaker/pkg/router"
)

func TestAddWordBook(t *testing.T) {
	token := initDB()
	w := httptest.NewRecorder()

	a := assert.New(t)
	r := router.New()

	req, _ := http.NewRequest("GET", "/wordbooks", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)

	a.Equal(200, w.Code)
	st := struct {
		Books []string `json:"wordbooks"`
	}{
		Books: []string{},
	}

	json.NewDecoder(w.Body).Decode(&st)
	a.Equal(1, len(st.Books))
	a.Equal("test", st.Books[0])
}

func TestPostWordBook(t *testing.T) {
	token := initDB()
	time.Sleep(time.Second)
	w := httptest.NewRecorder()

	a := assert.New(t)
	r := router.New()

	req, _ := http.NewRequest("POST", "/wordbooks/asd", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	a.Equal(201, w.Code)

	req, _ = http.NewRequest("POST", "/wordbooks/asd", nil)
	req.Header.Add("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	a.Equal(400, w.Code)

}

func TestGetWordBook(t *testing.T) {
	rr := view.Reponse{}
	fmt.Println(rr)
	token := initDB()
	w := httptest.NewRecorder()

	a := assert.New(t)
	r := router.New()
	req, _ := http.NewRequest("GET", "/wordbooks/test", nil)
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	a.Equal(200, w.Code)

	book := model.Wordbook{}
	json.NewDecoder(w.Body).Decode(&book)
	a.Equal("test", book.Name)
}
