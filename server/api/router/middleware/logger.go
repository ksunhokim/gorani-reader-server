package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	chmid "github.com/go-chi/chi/middleware"
	"github.com/sunho/gorani-reader/server/api/log"
)

var LoggerCtxKey = &contextKey{name: "logger"}

func Logger(logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()
			r = WithLogger(r, logger)
			ww := chmid.NewWrapResponseWriter(w, r.ProtoMajor)
			defer func() {
				scheme := "http"
				if r.TLS != nil {
					scheme = "https"
				}

				url := fmt.Sprintf("%s://%s%s %s", scheme, r.Host, r.RequestURI, r.Proto)
				logger.Log(log.TagRequest, log.M{
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
}

func WithLogger(r *http.Request, logger log.Logger) *http.Request {
	r = r.WithContext(context.WithValue(r.Context(), LoggerCtxKey, logger))
	return r
}

func GetLogger(r *http.Request) log.Logger {
	logger, _ := r.Context().Value(LoggerCtxKey).(log.Logger)
	return logger
}
