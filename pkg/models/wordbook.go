package models

import (
	"database/sql"
	"time"
)

type (
	WordBook struct {
		ID        int       `json:"-"`
		UserID    int       `db:"user_id" json:"-"`
		CreatedAt time.Time `db:"created_at" json:"created_at"`
		SeenAt    time.Time `db:"seen_at" json:"seen_at"`
		Name      string    `json:"name"`
	}

	WordBookDef struct {
		WordBookID int            `db:"wordbook_id" json:"-"`
		No         int            `db:"sr_no" json:"no"`
		DefID      int            `db:"def_id" json:"def_id"`
		Star       sql.NullBool   `json:"star"`
		Def        string         `json:"def"`
		Part       sql.NullString `json:"part"`
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

func (wb WordBook) Remove() error {
	_, err := db.Exec(
		`DELETE FROM wordbooks
		WHERE id = ?`, wb.ID)
	return err
}

func (wd WordBookDef) Remove() error {
	_, err := db.Exec(
		`DELETE FROM wordbook_entries
		WHERE wordbook_id = ?
		AND def_id = ?`, wd.WordBookID, wd.DefID)
	return err
}
