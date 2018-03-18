package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sunho/engbreaker/pkg/auth"
	"github.com/sunho/engbreaker/pkg/models"
)

const tokenContextKey = "middleware auth user"

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := getUser(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, err)
			return
		}
		ctx := context.WithValue(r.Context(), tokenContextKey, user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUser(r *http.Request) (models.User, error) {
	c, err := r.Cookie(auth.CookieName)
	if err != nil {
		return models.User{}, err
	}

	token := c.Value
	user, err := auth.ParseToken(token)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUser(r *http.Request) (models.User, error) {
	val := r.Context().Value(tokenContextKey)
	if val == nil {
		return models.User{}, fmt.Errorf("no user in context")
	}

	user, ok := val.(models.User)
	if !ok {
		return models.User{}, fmt.Errorf("user type conversion failed")
	}

	return user, nil
}
