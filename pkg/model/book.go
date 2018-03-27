package model

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	bongo.DocumentBase `bson:",inline"`
	UserID             bson.ObjectId `json:"user_id"`
	Title              string        `json:"title"`
	Picture            string        `json:"picture"`
	Chapters           []Chapter     `json:"chapters"`
}

type Chapter struct {
	Title     string
	ContentId bson.ObjectId
}

type ChapterContent struct {
	bongo.DocumentBase `bson:",inline"`
	Content            string
}
