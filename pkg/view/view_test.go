package view_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/markbates/goth"
	"github.com/sunho/engbreaker/pkg/auth"
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/dbs"
	"github.com/sunho/engbreaker/pkg/model"
	"github.com/sunho/engbreaker/pkg/router"
	httpexpect "gopkg.in/gavv/httpexpect.v1"
)

func initWordDB() {
	os.Setenv("MONGO_DB", "wordtest") // should preparej
}

func initServer(t *testing.T) (*httptest.Server, *httpexpect.Expect) {
	handler := router.New()
	server := httptest.NewServer(handler)
	e := httpexpect.New(t, server.URL)
	return server, e
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
				WordRef: model.WordRef{
					Word:       "test",
					Definition: 0,
				},
				Book: "test",
				Star: true,
			},
		},
	}
	model.Save(&book)
	chapter := model.ChapterContent{
		Content: "<div>호이</div>",
	}
	model.Save(&chapter)
	book2 := model.Book{
		Title:     "test",
		Picture:   "test.png",
		View:      10,
		Completed: 1,
		Chapters: []model.Chapter{
			model.Chapter{
				Title:     "hoi!호이",
				ContentID: chapter.Id,
			},
		},
	}
	model.Save(&book2)
	book3 := []model.Book{
		model.Book{
			Title:     "test2",
			View:      5,
			Completed: 3,
		},
		model.Book{
			Title:     "test0",
			View:      4,
			Completed: 3,
		},
		model.Book{
			Title:     "테스트",
			Author:    "호잇",
			View:      3,
			Completed: 3,
		},
		model.Book{
			Title:     "테스트2",
			Author:    "호잇",
			View:      2,
			Completed: 3,
		},
	}
	model.Save(&book3)

	return token
}

func testEndpoint(token string, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r := router.New()
	req.Header.Add("Authorization", token)
	r.ServeHTTP(w, req)
	return w
}
