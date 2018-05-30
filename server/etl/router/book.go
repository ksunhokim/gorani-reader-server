package router

import (
	"net/http"

	"github.com/sunho/gorani-reader/server/etl/book"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
)

func (ro *Router) PutBook(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	defer file.Close()

	words, err := dbh.GetWords(ro.gorn.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	size := handler.Size

	d := book.NewDictionary(words)
	_, err = book.Parse(d, file, size)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

}
