package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sunho/engbreaker/pkg/auth"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(auth.CookieName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		token := c.Value
		user, err := auth.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
