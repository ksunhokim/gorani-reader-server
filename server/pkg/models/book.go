package models

type Book struct {
	Id   int    `gorm:"column:book_id"`
	Name string `gorm:"column:book_name"`
}
