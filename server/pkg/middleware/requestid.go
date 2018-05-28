package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

var RequestIdCtxKey = &contextKey{name: "request id"}

func RequestId(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r = WithRequestId(r)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func WithRequestId(r *http.Request) *http.Request {
	r = r.WithContext(context.WithValue(r.Context(), RequestIdCtxKey, uuid.New()))
	return r
}

func GetRequestId(r *http.Request) uuid.UUID {
	id, _ := r.Context().Value(RequestIdCtxKey).(uuid.UUID)
	return id
}
