package dbh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func TestFindRelevantWords(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)

	word, err := dbh.GetWordById(gorn.Mysql, 1)
	a.Nil(err)
	word2, err := dbh.GetWordById(gorn.Mysql, 2)
	a.Nil(err)
	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)
	words, err := user.FindRelevantKnownWords(gorn.Mysql, "test", word, 10)
	a.Nil(err)
	a.Equal(0, len(words))

	user.AddKnownWord(gorn.Mysql, word.Id)
	user.AddKnownWord(gorn.Mysql, word2.Id)

	words, err = user.FindRelevantKnownWords(gorn.Mysql, "test", word, 10)
	a.Nil(err)
	a.Equal(1, len(words))
}
