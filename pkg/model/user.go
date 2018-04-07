package model

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	bongo.DocumentBase `bson:",inline"`
	Email              string   `json:"email"`
	Nickname           string   `json:"nickname"`
	AuthProvider       string   `json:"auth_provider"`
	AuthID             string   `json:"auth_id"`
	Wordbooks          []string `json:"wordbooks"`
}

func (user *User) LiftWordbook(wordbook string) {
	user.DeleteWordbook(wordbook)
	user.Wordbooks = append([]string{wordbook}, user.Wordbooks...)
}

func (user *User) DeleteWordbook(wordbook string) {
	for i, book := range user.Wordbooks {
		if book == wordbook {
			user.Wordbooks = append(user.Wordbooks[:i], user.Wordbooks[i+1:]...)
			return
		}
	}
}

func (user *User) AddWordbook(wordbook string) {
	user.Wordbooks = append([]string{wordbook}, user.Wordbooks...)
}

func (user *User) GetWordbook(wordbook string) (Wordbook, error) {
	book := Wordbook{}
	err := Get(&book, bson.M{
		"userid": user.GetId(),
		"name":   wordbook,
	})
	return book, err
}
