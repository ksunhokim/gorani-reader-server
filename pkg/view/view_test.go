package view_test

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/markbates/goth"
	"github.com/sunho/engbreaker/pkg/auth"
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/dbs"
	"github.com/sunho/engbreaker/pkg/model"
	"github.com/sunho/engbreaker/pkg/router"
)

func initWordDB() {
	os.Setenv("MONGO_DB", "wordtest") // should preparej
	config.Debug = true
	dbs.Init()
}

func initDB() string {
	os.Setenv("MONGO_DB", "viewtest")
	config.Debug = true
	dbs.Init()
	dbs.MDB.Session.DB("viewtest").DropDatabase()
	model.MigrateIndex()
	word := model.Word{
		Word:          "test",
		Pronunciation: "test",
		Definitions: []model.Definition{
			model.Definition{
				Definition: "hello",
				Part:       "verb",
				Examples: []model.Example{
					model.Example{
						First:  "hello",
						Second: "안녕",
					},
				},
			},
			model.Definition{
				Definition: "world",
				Part:       "noun",
				Examples: []model.Example{
					model.Example{
						First:  "world",
						Second: "월드",
					},
				},
			},
		},
	}
	model.Save(&word)
	token := auth.GetTokenOrRegister(
		goth.User{
			Provider: "admin",
			UserID:   "hohoho",
			NickName: "test",
			Email:    "asd@asd.asdf",
		},
	)
	user, _ := auth.ParseToken(token)
	user.Wordbooks = []string{"test"}
	model.Save(&user)
	unkown := model.Unkown{
		UserID: user.GetId(),
		Words: []model.UnkownWord{
			model.UnkownWord{
				Word:       "test",
				Definition: 0,
				Book:       "test",
			},
		},
	}
	model.Save(&unkown)
	book := model.Wordbook{
		UserID: user.GetId(),
		Name:   "test",
		Entries: []model.WordbookEntry{
			model.WordbookEntry{
				Word:       "test",
				Definition: 0,
				Book:       "test",
				Star:       true,
			},
		},
	}
	model.Save(&book)
	return token
}

func testEndpoint(token string, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r := router.New()
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	return w
}
