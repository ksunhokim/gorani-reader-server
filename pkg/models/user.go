package models

import (
	"time"
)

type User struct {
	ID        int
	CreatedAt time.Time `db:"created_at"`
	Username  string
	Email     string
}

func AddUser(u User) error {
	_, err := db.Exec(
		`INSERT INTO users (email, username)
		 VALUES (?,?)`, u.Email, u.Username)
	return err
}

func GetUser(email string) (User, error) {
	u := User{}
	err := db.Get(&u,
		`SELECT * from users
		WHERE email = ?`, email)
	return u, err
}

func (u User) AddWordBook(wb WordBook) error {
	_, err := db.Exec(
		`INSERT INTO wordbooks (user_id, name)
		VALUES (?,?)`, u.ID, wb.Name)
	return err
}

func (u User) GetWordBooks() ([]WordBook, error) {
	entries := []WordBook{}
	err := db.Select(&entries,
		`SELECT * FROM wordbooks
		WHERE user_id = ?`, u.ID)
	return entries, err
}

func (u User) GetWordBook(name string) (WordBook, error) {
	entry := WordBook{}
	err := db.Get(&entry,
		`SELECT * FROM wordbooks
		WHERE name = ?
		AND user_id = ?`, name, u.ID)
	return entry, err
}
