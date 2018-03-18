package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunho/engbreaker/pkg/models"
)

type briefWord struct {
	No   int    `json:"no"`
	Pron string `json:"pron"`
	Word string `json:"word"`
	Type string `json:"type"`
}

var wordsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	word := mux.Vars(r)["word"]

	ws, err := models.GetWords(word)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	res := []briefWord{}
	for index, word := range ws {
		res = append(res, briefWord{
			No:   index,
			Pron: word.Pron.String,
			Word: word.Word,
			Type: word.Type.String,
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
})

type detailWord struct {
	No   int          `json:"no"`
	Pron string       `json:"pron"`
	Word string       `json:"word"`
	Type string       `json:"type"`
	Defs []models.Def `json:"defs"`
}

var wordHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	word := mux.Vars(r)["word"]
	index := mux.Vars(r)["index"]
	i, _ := strconv.Atoi(index)

	wo, err := models.GetWord(word, i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	defs, err := wo.GetDefs()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	res := detailWord{
		Word: wo.Word,
		Type: wo.Type.String,
		Pron: wo.Pron.String,
		No:   i,
		Defs: defs,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
})
