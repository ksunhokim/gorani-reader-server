package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	minio "github.com/minio/minio-go"
	"github.com/sunho/gorani-reader/server/etl/dict"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

var wordContextKey = contextKey{"word"}

func (ro *Router) WordCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		word, err := dbh.GetWord(ro.gorn.Mysql, id)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), wordContextKey, word)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ro *Router) AddWord(w http.ResponseWriter, r *http.Request) {
	word := dbh.Word{}

	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = dbh.AddWord(ro.gorn.Mysql, &word)
	if err != nil {
		http.Error(w, err.Error(), 409)
		return
	}

	resp := util.M{
		"id": word.Id,
	}

	util.JSON(w, resp)
}

func (ro *Router) GetWords(w http.ResponseWriter, r *http.Request) {
	words, err := dbh.GetWords(ro.gorn.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	util.JSON(w, words)
}

func (ro *Router) GetWord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	word := ctx.Value(wordContextKey).(dbh.Word)
	util.JSON(w, word)
}

func (ro *Router) DeleteWord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	word := ctx.Value(wordContextKey).(dbh.Word)
	err := word.Delete(ro.gorn.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (ro *Router) CreateSqliteDict(w http.ResponseWriter, r *http.Request) {
	url2, err := dict.Create(ro.gorn.Mysql)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = ro.gorn.S3.FPutObject("dict", "dict_sqlite.db", url2, minio.PutObjectOptions{
		ContentType: "application/x-sqlite3",
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	reqParams := url.Values{}
	reqParams.Set("response-content-disposition", `attachment; filename="dict.db"`)

	obj, err := ro.gorn.S3.PresignedGetObject("dict", "dict_sqlite.db", time.Hour*24, reqParams)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintln(w, obj.String())
}
