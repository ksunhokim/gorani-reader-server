package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"

	"github.com/sunho/engbreaker/pkg/auth"
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/middlewares"
	"github.com/sunho/goth/providers/naver"
)

func init() {
	url := config.GetString("URL", "http://localhost:3000")
	goth.UseProviders(
		naver.New(config.GetString("NAVER_KEY", "naver key"), config.GetString("NAVER_SECRET", "naver secret"), url+"/auth/naver/callback"),
	)
}

var beginAuthHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]
	url, err := auth.GetAuthURL(provider)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
})

var completeAuthHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]
	user, err := auth.CompleteAuth(provider, r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	cookie := http.Cookie{Path: "/", Name: auth.CookieName, Value: auth.GetTokenOrRegister(user), HttpOnly: true}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
})

var refreshAuthHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, err := middlewares.GetUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	cookie := http.Cookie{Path: "/", Name: auth.CookieName, Value: auth.CreateToken(user.Email), HttpOnly: true}
	http.SetCookie(w, &cookie)
})
