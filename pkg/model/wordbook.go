package model

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type Wordbook struct {
	bongo.DocumentBase `bson:",inline"`
	UserID             bson.ObjectId
	Name               string
	Entries            []WordBookEntry
}

type WordBookEntry struct {
	Word            string
	DefinitionIndex uint
	Pronunciation   string
	Star            bool
}
