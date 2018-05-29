package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"

	"github.com/sunho/gorani-reader/server/pkg/log"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				body, _ := ioutil.ReadAll(r.Body)
				log.Log(log.TopicError, util.M{
					"panic":  fmt.Sprintf("%v", rvr),
					"body":   body,
					"stack":  string(debug.Stack()),
					"req_id": GetRequestId(r).String(),
				})

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
