package dbh_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/go/pkg/dbh"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
)

func TestUnknownWord(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)

	user, err := dbh.GetUser(gorn.Mysql, util.TestUserId)
	a.Nil(err)

	u := 2
	word := dbh.UnknownWord{
		WordId:         1,
		MemroySentence: util.NewString("asdf"),
		AddedDate:      util.RFCTime{time.Now().UTC().Round(time.Second)},
		Sources: []dbh.UnknownWordSource{
			dbh.UnknownWordSource{
				DefinitionId: 1,
				Book:         util.NewString("asdf"),
				Sentence:     util.NewString("asdf2"),
				WordIndex:    &u,
			},
		},
	}
	err = user.PutUnknownWord(gorn.Mysql, &word)
	a.Nil(err)

	words, err := user.GetUnknownWordWithQuizs(gorn.Mysql)
	a.Nil(err)
	a.Equal(1, len(words))
	word1 := words[0]
	a.Equal(word, word1.UnknownWord)

	a.Nil(word1.QuizDate)
	a.Nil(word1.AverageGrade)

	err = user.PutUnknownWord(gorn.Mysql, &word)
	a.Nil(err)

	words, err = user.GetUnknownWordWithQuizs(gorn.Mysql)
	a.Nil(err)
	a.Equal(1, len(words))
	word1 = words[0]
	a.Equal(word, word1.UnknownWord)

	word2, err := user.GetUnknownWord(gorn.Mysql, 1)
	a.Nil(err)
	a.Equal(word, word2)

	err = word1.Delete(gorn.Mysql)
	a.Nil(err)

	words, err = user.GetUnknownWordWithQuizs(gorn.Mysql)
	a.Nil(err)
	a.Equal(0, len(words))
}
