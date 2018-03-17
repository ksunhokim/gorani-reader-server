package middlewares

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tomasen/realip"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		crw := newCustomResponseWriter(w)
		defer logrus.WithFields(logrus.Fields{
			"status":   crw.status,
			"method":   r.Method,
			"path":     r.URL.Path,
			"ip":       realip.FromRequest(r),
			"size":     crw.size,
			"duration": time.Since(start).Nanoseconds() / int64(time.Millisecond),
		}).Info("http request")
		h.ServeHTTP(crw, r)
	})
}

//https://github.com/unrolled/logger/blob/master/logger.go
type customResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (c *customResponseWriter) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

func (c *customResponseWriter) Write(b []byte) (int, error) {
	size, err := c.ResponseWriter.Write(b)
	c.size += size
	return size, err
}

func newCustomResponseWriter(w http.ResponseWriter) *customResponseWriter {
	return &customResponseWriter{
		ResponseWriter: w,
		status:         200,
	}
}
