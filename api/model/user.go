package model

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	bongo.DocumentBase `bson:",inline"`
	Email              string          `json:"email"`
	Nickname           string          `json:"nickname"`
	AuthProvider       string          `json:"auth_provider"`
	AuthID             string          `json:"auth_id"`
	Books              []bson.ObjectId `json:"-"`
	Wordbooks          []bson.ObjectId `json:"-"`
	Unkown             bson.ObjectId   `json:"-"`
}

func (user *User) GetWordbook(wordbook string) (Wordbook, error) {
	book := Wordbook{}
	err := Get(&book, bson.M{
		"userid": user.GetId(),
		"name":   wordbook,
	})
	return book, err
}
