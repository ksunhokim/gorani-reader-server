package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sunho/engbreaker/pkg/auth"
)

type AuthContextKey string

func (k AuthContextKey) String() string {
	return "middleware auth " + string(k)
}

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(auth.CookieName)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, err)
			return
		}

		token := c.Value
		user, err := auth.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), AuthContextKey("user"), user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthOrRedirect(h http.Handler, url string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(auth.CookieName)
		if err != nil {
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			return
		}

		token := c.Value
		user, err := auth.ParseToken(token)
		if err != nil {
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			return
		}

		ctx := context.WithValue(r.Context(), AuthContextKey("user"), user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
