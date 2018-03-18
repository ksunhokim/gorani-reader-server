package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunho/engbreaker/pkg/middlewares"
	"github.com/sunho/engbreaker/pkg/models"
)

var wordBookListHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, err := middlewares.GetUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	books, _ := user.GetWordBooks()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
})

var wordBookDetailHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, err := middlewares.GetUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	name := mux.Vars(r)["name"]

	wb, err := user.GetWordBook(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	defs, _ := wb.GetDefs()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(defs)
})

var wordBookAddHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, err := middlewares.GetUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	name := mux.Vars(r)["name"]

	wb := models.WordBook{}
	err = json.NewDecoder(r.Body).Decode(&wb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	wb.Name = name

	err = user.AddWordBook(wb)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
})

var wordBookRemoveHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, err := middlewares.GetUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	name := mux.Vars(r)["name"]

	wb, err := user.GetWordBook(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	err = wb.Remove()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
})
