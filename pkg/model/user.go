package model

import "github.com/go-bongo/bongo"

type User struct {
	bongo.DocumentBase `bson:",inline"`
	Email              string   `json:"email"`
	Nickname           string   `json:"nickname"`
	AuthProvider       string   `json:"auth_provider"`
	AuthID             string   `json:"auth_id"`
	Wordbooks          []string `json:"wordbooks"`
}
