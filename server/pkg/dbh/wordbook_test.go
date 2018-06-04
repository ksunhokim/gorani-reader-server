package dbh_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func TestGetWordbook(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	id, _ := uuid.Parse(util.TestWordbookUuid)
	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)
	wordbook, err := user.GetWordbook(gorn.Mysql, dbh.UUID{id})
	a.Nil(err)

	a.Equal("test", wordbook.Name)
}

func TestAddWordbook(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)

	wordbook := dbh.Wordbook{
		Id:       dbh.UUID{uuid.New()},
		SeenDate: dbh.RFCTime{time.Now().UTC()},
		Name:     "asdf",
	}

	err = user.AddWordbook(gorn.Mysql, &wordbook)
	a.Nil(err)

	id := wordbook.Id

	wordbook, err = user.GetWordbook(gorn.Mysql, id)
	a.Nil(err)

	a.Equal("asdf", wordbook.Name)
	a.Equal(user.Id, wordbook.UserId)
}

func TestGetWordbooks(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)

	wordbooks, err := user.GetWordbooks(gorn.Mysql)
	a.Nil(err)

	a.Equal(1, len(wordbooks))

	wordbook := wordbooks[0]
	a.Equal("test", wordbook.Name)
}

func TestUpdateWordbook(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	id, _ := uuid.Parse(util.TestWordbookUuid)
	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)
	wordbook, err := user.GetWordbook(gorn.Mysql, dbh.UUID{id})
	a.Nil(err)

	a.Equal("test", wordbook.Name)

	wordbook.Name = "hoi"
	wordbook.Update(gorn.Mysql)

	wordbook, err = user.GetWordbook(gorn.Mysql, dbh.UUID{id})
	a.Nil(err)

	a.Equal("hoi", wordbook.Name)
}

func TestDeleteWordbook(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)

	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)

	wordbooks, err := user.GetWordbooks(gorn.Mysql)
	a.Nil(err)

	a.Equal(1, len(wordbooks))

	wordbook := wordbooks[0]
	a.Equal("test", wordbook.Name)

	wordbook.Delete(gorn.Mysql)

	wordbooks, err = user.GetWordbooks(gorn.Mysql)
	a.Nil(err)

	a.Equal(0, len(wordbooks))
}
