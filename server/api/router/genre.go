package router

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/middleware"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func genresToStrs(genres []dbh.Genre) []string {
	strs := make([]string, 0, len(genres))
	for _, genre := range genres {
		strs = append(strs, genre.Name)
	}
	return strs
}

func strsToGenres(db *gorm.DB, strs []string) (genres []dbh.Genre, err error) {
	var genre dbh.Genre
	for _, str := range strs {
		genre, err = dbh.GetGenreByName(db, str)
		if err != nil {
			return
		}
		genres = append(genres, genre)
	}
	return
}

func (ro *Router) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := dbh.GetGenres(ro.ap.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	strs := genresToStrs(genres)
	util.JSON(w, strs)
}

func (ro *Router) GetUserPreferGenres(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)
	genres, err := user.GetPreferGenres(ro.ap.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	strs := genresToStrs(genres)
	util.JSON(w, strs)
}

func (ro *Router) PutUserPreferGenres(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)

	strs := []string{}
	err := json.NewDecoder(r.Body).Decode(&strs)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	genres, err := strsToGenres(ro.ap.Mysql, strs)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = user.PutPreferGenres(ro.ap.Mysql, genres)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
