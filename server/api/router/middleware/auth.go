package middleware

import (
	"context"
	"net/http"

	"github.com/sunho/gorani-reader/server/api/api"
	"github.com/sunho/gorani-reader/server/pkg/auth"
	"github.com/sunho/gorani-reader/server/pkg/models"
)

var UserKey = &contextKey{name: "user id"}

const ApiKeyHeader = "X-API-Key"

func Auth(ap *api.Api) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get(ApiKeyHeader)

			id, name, err := auth.UserByApiKey(ap.Config.SecretKey, key)
			if err != nil {
				http.Error(w, http.StatusText(403), 403)
				return
			}

			user, err := models.GetUser(ap.Gorn.Mysql, id)
			if err != nil {
				http.Error(w, http.StatusText(403), 403)
				return
			}

			if user.Name != name {
				http.Error(w, http.StatusText(403), 403)
				return
			}

			r = WithUser(r, user)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func WithUser(r *http.Request, user models.User) *http.Request {
	r = r.WithContext(context.WithValue(r.Context(), UserKey, user))
	return r
}

func GetUser(r *http.Request) models.User {
	user, _ := r.Context().Value(UserKey).(models.User)
	return user
}
