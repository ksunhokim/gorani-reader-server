package view_test

import (
	"os"

	"github.com/markbates/goth"
	"github.com/sunho/engbreaker/pkg/auth"
	"github.com/sunho/engbreaker/pkg/config"
	"github.com/sunho/engbreaker/pkg/dbs"
	"github.com/sunho/engbreaker/pkg/model"
)

func initDB() string {
	os.Setenv("MONGO_DB", "viewtest")
	config.Debug = true
	dbs.Init()
	dbs.MDB.Session.DB("viewtest").DropDatabase()
	model.MigrateIndex()
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
	model.Save(user)
	book := model.Wordbook{
		UserID:  user.GetId(),
		Name:    "test",
		Entries: []model.WordbookEntry{},
	}
	model.Save(&book)
	return token
}
