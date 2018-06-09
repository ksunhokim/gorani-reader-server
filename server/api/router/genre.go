package router

import (
	"net/http"

	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func (ro *Router) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := dbh.GetGenres(ro.ap.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	strs := make([]string, 0, len(genres))

	for _, g := range genres {
		strs = append(strs, g.Name)
	}

	util.JSON(w, strs)
}
