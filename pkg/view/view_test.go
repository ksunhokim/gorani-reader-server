package view_test

import (
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/markbates/goth"
	"github.com/sunho/gorani-reader/pkg/auth"
	"github.com/sunho/gorani-reader/pkg/config"
	"github.com/sunho/gorani-reader/pkg/dbs"
	"github.com/sunho/gorani-reader/pkg/model"
	"github.com/sunho/gorani-reader/pkg/router"
	httpexpect "gopkg.in/gavv/httpexpect.v1"
)

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
	dbs.MDB.DB("").DropDatabase()
	dbs.MDB.DB("").C("words").Insert(
		model.Word{
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
		})
	token := auth.GetTokenOrRegister(
		goth.User{
			Provider: "admin",
			UserID:   "hohoho",
			NickName: "test",
			Email:    "asd@asd.asdf",
		},
	)
	user, _ := auth.ParseToken(token)
	dbs.MDB.DB("").C("wordbooks").Insert(model.Wordbook{
		UpdatedAt: time.Now(),
		UserId:    user.Id,
		Name:      "test",
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
	})

	/*chapter := model.ChapterContent{
		Content: "<div>호이</div>",
	}

	book2 := model.Book{
		UserID:    user.GetId(),
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
	user.Books = []bson.ObjectId{book2.GetId()}
	model.Save(&user)*/
	return token
}
