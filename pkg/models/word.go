package models

import "database/sql"

type (
	Word struct {
		ID     int
		Word   string
		Pron   sql.NullString
		Source string
		Type   sql.NullString
	}

	Def struct {
		ID     int
		WordID int
		Part   string
		Def    string
	}
)
