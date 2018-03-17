package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int
	CreatedAt time.Time `db:"created_at"`
	Username  string
	Email     string
}

func GetUser(db *sqlx.DB, id int) (User, error) {
	u := User{}
	err := db.Get(&u,
		`SELECT * from users
		WHERE id = ?`, id)
	return u, err
}

func (u User) AddWordBook(db *sqlx.DB, wb WordBook) error {
	_, err := db.Exec(
		`INSERT INTO wordbooks (user_id, name)
		VALUES (?,?)`, u.ID, wb.Name)
	return err
}

func (u User) GetWordBooks(db *sqlx.DB) ([]WordBook, error) {
	entries := []WordBook{}
	err := db.Select(&entries,
		`SELECT * FROM wordbooks
		WHERE user_id = ?`, u.ID)
	return entries, err
}
