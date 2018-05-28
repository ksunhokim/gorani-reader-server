package middleware

import (
	"fmt"
	"net/http"
	"time"

	chmid "github.com/go-chi/chi/middleware"
	"github.com/sunho/gorani-reader/server/pkg/log"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		ww := chmid.NewWrapResponseWriter(w, r.ProtoMajor)
		defer func() {
			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}
			url := fmt.Sprintf("%s://%s%s %s", scheme, r.Host, r.RequestURI, r.Proto)
			log.Log(log.TopicRequest.Api(), log.M{
				"status":      ww.Status(),
				"bytes":       ww.BytesWritten(),
				"url":         url,
				"spent":       int64(time.Since(t) / time.Millisecond),
				"req_id":      GetRequestId(r).String(),
				"remote_addr": r.RemoteAddr,
			})
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}
