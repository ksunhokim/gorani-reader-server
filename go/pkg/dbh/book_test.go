package dbh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/go/pkg/dbh"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
)

func TestFindSentences(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)

	word1, err := dbh.GetWordById(gorn.Mysql, 1)
	a.Nil(err)
	word2, err := dbh.GetWordById(gorn.Mysql, 2)
	a.Nil(err)
	sens, err := dbh.FindSentences(gorn.Mysql, word1, word2, 1)
	a.Nil(err)

	a.Equal(1, len(sens))
	a.Equal("test test2", sens[0].Sentence)

	sens, err = dbh.FindSentences(gorn.Mysql, word1, word2, 0)
	a.Nil(err)

	a.Equal(0, len(sens))
}

func TestGetKnownDegree(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)

	word1, err := dbh.GetWordById(gorn.Mysql, 1)
	a.Nil(err)
	word2, err := dbh.GetWordById(gorn.Mysql, 2)
	a.Nil(err)
	sens, err := dbh.FindSentences(gorn.Mysql, word1, word2, 1)
	a.Nil(err)

	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)
	sen := sens[0]

	err = user.AddKnownWord(gorn.Mysql, word1.Id)
	a.Nil(err)

	known, total, err := user.GetKnownDegreeOfSentence(gorn.Mysql, sen)
	a.Nil(err)

	a.Equal(1, known)
	a.Equal(2, total)
}
