package middlewares

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tomasen/realip"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		crw := newCustomResponseWriter(w)
		start := time.Now()
		h.ServeHTTP(crw, r)
		ip := realip.FromRequest(r)
		path := r.URL.Path
		method := r.Method
		duration := time.Since(start)
		size := crw.size
		status := crw.status
		logrus.WithFields(logrus.Fields{
			"status":   status,
			"method":   method,
			"path":     path,
			"ip":       ip,
			"size":     size,
			"duration": duration.Nanoseconds() / int64(time.Millisecond),
		}).Info("http request")
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
