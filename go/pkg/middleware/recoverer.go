package middleware

import (
	"io/ioutil"
	"net/http"
	"runtime/debug"

	"github.com/yanyiwu/simplelog"
)

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				body, _ := ioutil.ReadAll(r.Body)
				simplelog.Error("panic recovered | cause: %v \n body: %s \n stack: %s", rvr, body, string(debug.Stack()))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
