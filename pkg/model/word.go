package model

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type (
	Word struct {
		bongo.DocumentBase `bson:",inline" json"-"`
		Word               string       `json:"word"`
		Pronunciation      string       `json:"pronunciation"`
		Definitions        []Definition `json:"definitions"`
	}

	Definition struct {
		Definition string    `json:"definition"`
		Part       string    `json:"part"`
		Examples   []Example `json:"examples"`
	}

	Example struct {
		First  string `json:"first"`
		Second string `json:"second"`
	}
)

func GetWord(word string) (Word, error) {
	word_ := Word{}
	err := Get(&word_, bson.M{
		"word": word,
	})
	return word_, err
}
func ValidateWord(ref WordRef) bool {
	word, err := GetWord(ref.Word)
	if err != nil {
		return false
	}
	if uint(len(word.Definitions)) <= ref.Definition {
		return false
	}
	return true
}
