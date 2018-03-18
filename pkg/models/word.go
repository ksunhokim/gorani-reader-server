package models

import (
	"database/sql"
	"fmt"
)

type (
	Word struct {
		ID     int            `json:"-"`
		Word   string         `json:"word"`
		Pron   sql.NullString `json:"-"`
		Source string         `json:"-"`
		Type   sql.NullString `json:"type"`
	}

	Def struct {
		ID     int    `json:"id"`
		WordID int    `db:"word_id" json:"-"`
		Part   string `json:"part"`
		Def    string `json:"def"`
	}
)

func GetWords(word string) ([]Word, error) {
	words := []Word{}
	err := db.Select(&words,
		`SELECT * from words
		WHERE word = ?
		ORDER BY type, id`, word)
	return words, err
}

func GetWord(word string, index int) (Word, error) {
	words, _ := GetWords(word)
	if index >= len(words) {
		return Word{}, fmt.Errorf("word for %s:%d doesn't exist", word, index)
	}
	return words[index], nil
}

func (w Word) GetDefs() ([]Def, error) {
	defs := []Def{}
	err := db.Select(&defs,
		`SELECT * from defs
		WHERE word_id = ?
		ORDER BY id`, w.ID)
	return defs, err
}
