package model

import "github.com/go-bongo/bongo"

type Word struct {
	bongo.DocumentBase `bson:",inline" json"-"`
	Word               string       `json:"word"`
	Pronunciation      string       `json:"pronunciation"`
	Definitions        []Definition `json:"definitions"`
}

type Definition struct {
	Definition string    `json:"definition"`
	Part       string    `json:"part"`
	Examples   []Example `json:"examples"`
}

type Example struct {
	First  string `json:"first"`
	Second string `json:"second"`
}
