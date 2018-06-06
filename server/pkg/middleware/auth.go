package middleware

import (
	"context"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader/server/pkg/auth"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
)

var UserKey = &contextKey{name: "user id"}

const ApiKeyHeader = "X-API-Key"

func Auth(db *gorm.DB, secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get(ApiKeyHeader)

			id, name, err := auth.UserByApiKey(secret, key)
			if err != nil {
				http.Error(w, http.StatusText(403), 403)
				return
			}

			user, err := dbh.GetUser(db, id)
			if err != nil {
				http.Error(w, http.StatusText(403), 403)
				return
			}

			// invalid api key
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

func WithUser(r *http.Request, user dbh.User) *http.Request {
	r = r.WithContext(context.WithValue(r.Context(), UserKey, user))
	return r
}

func GetUser(r *http.Request) dbh.User {
	user, _ := r.Context().Value(UserKey).(dbh.User)
	return user
}
