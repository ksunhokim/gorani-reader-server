package models_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/api/models"
	"github.com/sunho/gorani-reader/server/api/util"
)

func TestGetWordbook(t *testing.T) {
	gorn := Setup()
	a := assert.New(t)
	id, _ := uuid.Parse(TestWordbookUuid)
	bytes := util.UuidToBytes(id)
	wordbook, err := models.GetWordbook(gorn.Mysql, bytes)
	a.Nil(err)

	a.Equal("test", wordbook.Name)
	a.Equal(false, wordbook.IsUnknown)
}

func TestAddWordbook(t *testing.T) {
	gorn := Setup()
	a := assert.New(t)
	user, err := models.GetUser(gorn.Mysql, TestUserId)
	a.Nil(err)

	wordbook := models.Wordbook{
		Id:   util.UuidToBytes(uuid.New()),
		Name: "asdf",
	}

	err = user.CreateWordbook(gorn.Mysql, &wordbook)
	a.Nil(err)

	id := wordbook.Id
	a.Equal(16, len(id))

	wordbook, err = models.GetWordbook(gorn.Mysql, id)
	a.Nil(err)

	a.Equal("asdf", wordbook.Name)
	a.Equal(user.Id, wordbook.UserId)
	a.Equal(false, wordbook.IsUnknown)
}

func TestGetWordbooks(t *testing.T) {
	gorn := Setup()
	a := assert.New(t)
	user, err := models.GetUser(gorn.Mysql, TestUserId)
	a.Nil(err)

	wordbooks, err := user.GetWordbooks(gorn.Mysql)
	a.Nil(err)

	a.Equal(1, len(wordbooks))

	wordbook := wordbooks[0]
	a.Equal("test", wordbook.Name)
}

func TestD