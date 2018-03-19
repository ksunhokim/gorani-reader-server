package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunho/engbreaker/pkg/models"
)

type detailWord struct {
	No   int    `json:"no"`
	Pron string `json:"pron"`
	Word string `json:"word"`
	Type string `json:"type"`
	Defs []def  `json:"defs"`
}

type def struct {
	ID   int    `json:"id"`
	Def  string `json:"def"`
	Part string `json:"part"`
}

var wordHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
		}
	}()

	id := mux.Vars(r)["id"]
	i, _ := strconv.Atoi(id)

	wo, err := models.GetWord(i)
	if err != nil {
		return
	}

	defs, err := wo.GetDefs()
	if err != nil {
		return
	}

	rdefs := []def{}
	for _, de := range defs {
		rdefs = append(rdefs, def{
			Def:  de.Def,
			ID:   de.ID,
			Part: de.Part.String,
		})
	}

	res := detailWord{
		Word: wo.Word,
		Type: wo.Type,
		Pron: wo.Pron.String,
		No:   i,
		Defs: rdefs,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
})
