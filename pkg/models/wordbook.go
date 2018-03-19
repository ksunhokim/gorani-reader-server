package models

import (
	"database/sql"
	"time"
)

type (
	WordBook struct {
		ID        int
		UserID    int       `db:"user_id"`
		CreatedAt time.Time `db:"created_at"`
		SeenAt    time.Time `db:"seen_at"`
		Name      string
	}

	WordBookDef struct {
		WordBookID int `db:"wordbook_id"`
		No         int `db:"sr_no"`
		DefID      int `db:"def_id"`
		Star       bool
		Def        string
		Part       sql.NullString
	}
)

func (wb WordBook) GetDefs() ([]WordBookDef, error) {
	entries := []WordBookDef{}
	err := db.Select(&entries,
		`SELECT * from defs_of_wordbooks
		WHERE wordbook_id = ?
		ORDER BY sr_no`, wb.ID)
	return entries, err
}

func (wb WordBook) AddDef(id int) error {
	_, err := db.Exec(
		`INSERT INTO wordbook_entries (wordbook_id, def_id)
		VALUES (?,?)`, wb.ID, id)
	return err
}

func (wb WordBook) RemoveDef(id int) error {
	_, err := db.Exec(
		`DELETE FROM wordbook_entries
		WHERE def_id = ?
		AND wordbook_id = ?`, id, wb.ID)
	return err
}

func (wb WordBook) Remove() error {
	_, err := db.Exec(
		`DELETE FROM wordbooks
		WHERE id = ?`, wb.ID)
	return err
}
