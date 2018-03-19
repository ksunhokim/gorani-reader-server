package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunho/engbreaker/pkg/middlewares"
	"github.com/sunho/engbreaker/pkg/models"
)

type briefWordBook struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	SeenAt    time.Time `json:"seen_at"`
}

var wordBookListHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, err := middlewares.GetUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	books, _ := user.GetWordBooks()
	res := []briefWordBook{}
	for _, book := range books {
		res = append(res, briefWordBook{
			Name:      book.Name,
			CreatedAt: book.CreatedAt,
			SeenAt:    book.SeenAt,
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
})

type wordBookDef struct {
	No     int    `json:"no"`
	ID     int    `json:"id"`
	WordID int    `json:"word_id"`
	Star   bool   `json:"star"`
	Def    string `json:"def"`
	Part   string `json:"part"`
}

var wordBookDetailHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
		}
	}()

	user, err := middlewares.GetUser(r)
	if err != nil {
		return
	}
	name := mux.Vars(r)["name"]

	wb, err := user.GetWordBook(name)
	if err != nil {
		return
	}

	defs, _ := wb.GetDefs()
	res := []wordBookDef{}
	for _, def := range defs {
		d, _ := models.GetDef(def.DefID)
		word, _ := d.GetWord()
		res = append(res, wordBookDef{
			No:     def.No,
			ID:     def.DefID,
			Star:   def.Star,
			Def:    def.Def,
			Part:   def.Part.String,
			WordID: word.ID,
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
})

var wordBookAddHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
		}
	}()

	user, err := middlewares.GetUser(r)
	if err != nil {
		return
	}
	name := mux.Vars(r)["name"]

	wb := models.WordBook{}
	err = json.NewDecoder(r.Body).Decode(&wb)
	if err != nil {
		return
	}
	wb.Name = name

	err = user.AddWordBook(wb)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
})

var wordBookRemoveHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
		}
	}()

	user, err := middlewares.GetUser(r)
	if err != nil {
		return
	}
	name := mux.Vars(r)["name"]

	wb, err := user.GetWordBook(name)
	if err != nil {
		return
	}

	err = wb.Remove()
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
})

var wordBookAddDefHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
		}
	}()

	user, err := middlewares.GetUser(r)
	if err != nil {
		return
	}

	name := mux.Vars(r)["name"]
	def := mux.Vars(r)["def"]

	i, err := strconv.Atoi(def)
	if err != nil {
		return
	}

	wb, err := user.GetWordBook(name)
	if err != nil {
		return
	}

	err = wb.AddDef(i)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
})

var wordBookRemoveDefHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
		}
	}()

	user, err := middlewares.GetUser(r)
	if err != nil {
		return
	}

	name := mux.Vars(r)["name"]
	def := mux.Vars(r)["def"]

	i, err := strconv.Atoi(def)
	if err != nil {
		return
	}

	wb, err := user.GetWordBook(name)
	if err != nil {
		return
	}

	err = wb.RemoveDef(i)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
})
