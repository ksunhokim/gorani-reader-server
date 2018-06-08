package sentence_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/etl/sentence"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func TestBook(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)

	word1, err := dbh.GetWordById(gorn.Mysql, 1)
	a.Nil(err)
	word2, err := dbh.GetWordById(gorn.Mysql, 2)
	a.Nil(err)
	sens, err := sentence.FindFromBook(gorn.Mysql, word1, word2, 1)
	a.Nil(err)

	a.Equal(1, len(sens))
	a.Equal("test test2", sens[0].Sentence)

	sens, err = sentence.FindFromBook(gorn.Mysql, word1, word2, 0)
	a.Nil(err)

	a.Equal(0, len(sens))
	a.NotNil(nil)
}
