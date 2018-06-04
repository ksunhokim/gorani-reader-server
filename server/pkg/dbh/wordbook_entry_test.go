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
	wordbook, err := user.GetWordbook(gorn.Mysql, util.UUID{id})
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
	wordbook, err := user.GetWordbook(gorn.Mysql, util.UUID{id})
	a.Nil(err)

	entry := dbh.WordbookEntry{
		DefinitionId: 2,
		AddedDate:    util.RFCTime{time.Now().UTC()},
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
	wordbook, err := user.GetWordbook(gorn.Mysql, util.UUID{id})
	a.Nil(err)

	entries, err := wordbook.GetEntries(gorn.Mysql)
	a.Nil(err)
	a.Equal(1, len(entries))

	str := "book2"
	str2 := "asdf2"
	entry := dbh.WordbookEntry{
		WordbookId:     util.UUID{id},
		DefinitionId:   2,
		SourceBook:     &str,
		SourceSentence: &str2,
		AddedDate:      util.RFCTime{time.Now().UTC()},
	}
	entries2 := []dbh.WordbookEntry{}
	for _, e := range entries {
		entries2 = append(entries2, e.WordbookEntry)
	}
	entries2 = append(entries2, entry)
	err = wordbook.UpdateEntries(gorn.Mysql, time.Now().UTC(), entries2)
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
	wordbook, err := user.GetWordbook(gorn.Mysql, util.UUID{id})
	a.Nil(err)

	entries, err := wordbook.GetEntries(gorn.Mysql)
	a.Nil(err)
	a.Equal(1, len(entries))

	entries2 := []dbh.WordbookEntry{}
	for _, e := range entries {
		entries2 = append(entries2, e.WordbookEntry)
	}

	err = wordbook.UpdateEntries(gorn.Mysql, time.Now().UTC().Add(time.Hour*-3), entries2)
	a.NotNil(err)
}
