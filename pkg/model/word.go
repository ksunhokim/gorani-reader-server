package model

import "github.com/go-bongo/bongo"

type Word struct {
	bongo.DocumentBase `bson:",inline"`
	Word               string
	Pronunciation      string
	Definitions        []Definition
}

type Definition struct {
	Definition string
	Part       string
	Examples   []Example
}

type Example struct {
	First  string
	Second string
}
