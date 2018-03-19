package models

import (
	"database/sql"
)

type (
	Word struct {
		ID     int
		Word   string
		Pron   sql.NullString
		Source string
		Type   string
	}

	Def struct {
		ID     int
		WordID int `db:"word_id"`
		Part   sql.NullString
		Def    string
	}
)

func GetWord(id int) (Word, error) {
	word := Word{}
	err := db.Get(&word,
		`SELECT * from words
		WHERE id = ?`, id)
	return word, err
}

func (w Word) GetDefs() ([]Def, error) {
	defs := []Def{}
	err := db.Select(&defs,
		`SELECT * from defs
		WHERE word_id = ?
		ORDER BY id`, w.ID)
	return defs, err
}

func GetDef(id int) (Def, error) {
	def := Def{}
	err := db.Get(&def,
		`SELECT * from defs
		WHERE id = ?`, id)
	return def, err
}

func (d Def) GetWord() (Word, error) {
	word := Word{}
	err := db.Get(&word,
		`SELECT * from words
		WHERE id = ?`, d.WordID)
	return word, err
}
