package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/sunho/gorani-reader/server/api/log"
)

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				logger := GetLogger(r)
				if logger != nil {
					body, _ := ioutil.ReadAll(r.Body)
					logger.Log(log.TagRequest, log.M{
						"panic":  fmt.Sprintf("%v", rvr),
						"body":   body,
						"stack":  string(debug.Stack()),
						"req_id": GetRequestId(r).String(),
					})
				} else {
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)
					debug.PrintStack()
				}

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
