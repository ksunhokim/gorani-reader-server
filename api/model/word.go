package model

import (
	"github.com/sunho/engbreaker/pkg/dbs"
	"gopkg.in/mgo.v2/bson"
)

type (
	Word struct {
		Id            string       `json:"word"`
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
)

func GetWord(word string) (Word, error) {
	sess := dbs.MDB.Copy()
	defer sess.Close()

	word_ := Word{}
	err := sess.DB("").C("words").Find(bson.M{
		"_id": word,
	}).One(&word_)

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
