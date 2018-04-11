package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Title     string    `json:"title"`
	Picture   string    `json:"picture"`
	Author    string    `json:"author"`
	Chapters  []Chapter `json:"chapters"`
	View      uint      `json:"view"`
	Completed uint      `json:"completed"`
}

type Chapter struct {
	Title     string
	ContentID bson.ObjectId
}

type ChapterContent struct {
	Content string
}
