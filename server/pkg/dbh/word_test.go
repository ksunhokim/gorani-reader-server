package dbh_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/util"
)

func TestWord(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)

	word := dbh.Word{
		Word: "test3",
		Definitions: []dbh.Definition{
			dbh.Definition{
				Definition: "test1",
				Examples: []dbh.Example{
					dbh.Example{
						Foreign: "asdf",
					},
				},
			},
			dbh.Definition{
				WordId:     100,
				Definition: "test2",
			},
		},
	}
	err := dbh.AddWord(gorn.Mysql, &word)
	a.Nil(err)

	word2, err := dbh.GetWordById(gorn.Mysql, word.Id)
	a.Nil(err)

	a.Equal("test3", word2.Word)

	defs := word2.Definitions
	a.Nil(err)

	a.Equal(2, len(defs))
	a.Equal("test1", defs[0].Definition)
	a.Equal("test2", defs[1].Definition)

	examples := defs[0].Examples
	a.Nil(err)

	a.Equal(1, len(examples))
	a.Equal("asdf", examples[0].Foreign)
}

func TestGetWords(t *testing.T) {
	gorn := util.SetupTestGorani()
	a := assert.New(t)
	words, err := dbh.GetWords(gorn.Mysql)
	a.Nil(err)

	a.Equal(2, len(words))
	a.Equal("test", words[0].Word)
}
