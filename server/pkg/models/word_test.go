package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/models"
)

func TestWord(t *testing.T) {
	gorn := Setup()
	a := assert.New(t)

	word := models.Word{
		Word: "test2",
		Definitions: []models.Definition{
			models.Definition{
				Definition: "test1",
				Examples: []models.Example{
					models.Example{
						Foreign: "asdf",
					},
				},
			},
			models.Definition{
				Definition: "test2",
			},
		},
	}
	err := models.AddWord(gorn.Mysql, &word)
	a.Nil(err)

	word2, err := models.GetWord(gorn.Mysql, word.Id)
	a.Nil(err)

	a.Equal("test2", word2.Word)

	defs, err := word2.GetDefinitions(gorn.Mysql)
	a.Nil(err)

	a.Equal(2, len(defs))
	a.Equal("test1", defs[0].Definition)
	a.Equal("test2", defs[1].Definition)

	examples, err := defs[0].GetExamples(gorn.Mysql)
	a.Nil(err)

	a.Equal(1, len(examples))
	a.Equal("asdf", examples[0].Foreign)
}
