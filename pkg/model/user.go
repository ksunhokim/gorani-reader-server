package model

import "github.com/go-bongo/bongo"

type User struct {
	bongo.DocumentBase `bson:",inline"`
	Email              string
	Nickname           string
	AuthProvider       string
	AuthID             string
}
