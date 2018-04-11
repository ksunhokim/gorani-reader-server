package model

import (
	"github.com/sunho/engbreaker/pkg/dbs"
	"gopkg.in/mgo.v2/bson"
)

type (
	Word struct {
		Word          string       `json:"word" bson:"word,omitempty"`
		Pronunciation string       `json:"pronunciation"`
		Definitions   []Definition `json:"definitions"`
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

	WordRef struct {
		Word       string `json:"word"`
		Definition uint   `json:"definition"`
	}
)

func GetWord(word string) (Word, error) {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	out := Word{}
	err := sess.DB("").C("words").Find(bson.M{
		"word": word,
	}).One(&out)

	return out, err
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
