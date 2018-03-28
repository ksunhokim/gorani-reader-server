package model

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type Wordbook struct {
	bongo.DocumentBase `bson:",inline" json:"-"`
	UserID             bson.ObjectId   `json:"-"`
	Name               string          `json:"name"`
	Entries            []WordbookEntry `json:"entries"`
}

type WordbookEntry struct {
	Star       bool   `json:"star"`
	Word       string `json:"word"`
	Definition uint   `json:"def"`
	Book       string `json:"book"`
}
