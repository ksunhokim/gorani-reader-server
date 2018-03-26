package model

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	bongo.DocumentBase `bson:",inline"`
	UserID             bson.ObjectId
	Title              string
	Picture            string
	Chapters           []Chapter
}

type Chapter struct {
	Title     string
	ContentId bson.ObjectId
}

type ChapterContent struct {
	bongo.DocumentBase `bson:",inline"`
	Content            string
}
