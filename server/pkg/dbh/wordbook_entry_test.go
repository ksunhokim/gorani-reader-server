package dbh_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func TestWordbookGetEntries(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	id, _ := uuid.Parse(util.TestWordbookUuid)
	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)
	wordbook, err := user.GetWordbook(gorn.Mysql, dbh.UUID{id})
	a.Nil(err)

	entries, err := wordbook.GetEntries(gorn.Mysql)
	a.Nil(err)
	a.Equal(1, len(entries))

	entry := entries[0]
	a.Equal(1, entry.DefinitionId)
	a.Equal("book", *entry.SourceBook)
	a.Equal("asdf", *entry.SourceSentence)
}

func TestWordbookAddEntry(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	id, _ := uuid.Parse(util.TestWordbookUuid)
	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)
	wordbook, err := user.GetWordbook(gorn.Mysql, dbh.UUID{id})
	a.Nil(err)

	entry := dbh.WordbookEntry{
		DefinitionId: 2,
		AddedDate:    dbh.RFCTime{time.Now().UTC()},
	}
	err = wordbook.AddEntry(gorn.Mysql, time.Now().UTC(), &entry)
	a.Nil(err)

	entries, err := wordbook.GetEntries(gorn.Mysql)

	a.Equal(2, len(entries))
}

func TestWordbookUpdateEntries(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	id, _ := uuid.Parse(util.TestWordbookUuid)
	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)
	wordbook, err := user.GetWordbook(gorn.Mysql, dbh.UUID{id})
	a.Nil(err)

	entries, err := wordbook.GetEntries(gorn.Mysql)
	a.Nil(err)
	a.Equal(1, len(entries))

	str := "book2"
	str2 := "asdf2"
	entry := dbh.WordbookEntry{
		WordbookId:     dbh.UUID{id},
		DefinitionId:   2,
		SourceBook:     &str,
		SourceSentence: &str2,
		AddedDate:      dbh.RFCTime{time.Now().UTC()},
	}

	entries = append(entries, entry)
	err = wordbook.UpdateEntries(gorn.Mysql, time.Now().UTC(), entries)
	a.Nil(err)

	entries, err = wordbook.GetEntries(gorn.Mysql)
	a.Nil(err)
	a.Equal(2, len(entries))
}

func TestWordbookUpdateInvalidEntries(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	id, _ := uuid.Parse(util.TestWordbookUuid)
	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)
	wordbook, err := user.GetWordbook(gorn.Mysql, dbh.UUID{id})
	a.Nil(err)

	entries, err := wordbook.GetEntries(gorn.Mysql)
	a.Nil(err)
	a.Equal(1, len(entries))

	err = wordbook.UpdateEntries(gorn.Mysql, time.Now().UTC().Add(time.Hour*-3), entries)
	a.NotNil(err)
}
