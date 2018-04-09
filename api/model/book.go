package model

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	bongo.DocumentBase `bson:",inline"`
	UserID             bson.ObjectId `json:"-"`
	Title              string        `json:"title"`
	Picture            string        `json:"picture"`
	Author             string        `json:"author"`
	Chapters           []Chapter     `json:"chapters"`
	View               uint          `json:"view"`
	Completed          uint          `json:"completed"`
}

type Chapter struct {
	Title     string
	ContentID bson.ObjectId
}

type ChapterContent struct {
	bongo.DocumentBase `bson:",inline"`
	Content            string
}
