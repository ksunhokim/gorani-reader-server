package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
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
		Star       sql.NullBool
		Def        string
		Part       sql.NullString
	}
)

func (wb WordBook) GetDefs(db *sqlx.DB) ([]WordBookDef, error) {
	entries := []WordBookDef{}
	err := db.Select(&entries,
		`SELECT * from defs_of_wordbooks
		WHERE wordbook_id = ?
		ORDER BY sr_no`, wb.ID)
	return entries, err
}

func (wb WordBook) AddDef(db *sqlx.DB, id int) error {
	_, err := db.Exec(
		`INSERT INTO wordbook_entries (wordbook_id, def_id)
		VALUES (?,?)`, wb.ID, id)
	return err
}

func (wb WordBook) Remove(db *sqlx.DB) error {
	_, err := db.Exec(
		`DELETE FROM wordbooks
		WHERE id = ?`, wb.ID)
	return err
}

func (wd WordBookDef) Remove(db *sqlx.DB) error {
	_, err := db.Exec(
		`DELETE FROM wordbook_entries
		WHERE wordbook_id = ?
		AND def_id = ?`, wd.WordBookID, wd.DefID)
	return err
}
