package model

import (
	"github.com/go-bongo/bongo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Wordbook struct {
	bongo.DocumentBase `bson:",inline"`
	UserID             bson.ObjectId   `json:"-"`
	Name               string          `json:"name"`
	Entries            []WordbookEntry `json:"entries"`
}

func (wb *Wordbook) BeforeSaveConfigure(sess *mgo.Session) error {
	return nil
}

func (wb *Wordbook) AfterSaveConfigure(sess *mgo.Session) error {
	return nil
}

type WordbookEntry struct {
	UnkownWord
	Star bool `json:"star"`
}
