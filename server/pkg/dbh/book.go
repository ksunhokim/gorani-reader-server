package dbh

import "github.com/jinzhu/gorm"

type Book struct {
	Isbn       string  `gorm:"column:book_isbn"`
	Name       string  `gorm:"column:book_name"`
	Author     *string `gorm:"column:book_author"`
	CoverImage *string `gorm:"column:book_cover_image"`
}

type Sentence struct {
	BookIsbn *string `gorm:"column:book_isbn"`
	Id       int     `gorm:"column:sentence_id"`
	Sentence string  `gorm:"column:sentence"`
}

type WordSentence struct {
	WordId     int `gorm:"column:word_id"`
	SentenceId int `gorm:"column:sentence_id"`
	Position   int `gorm:"column:word_sentence_position"`
}

type BookRating struct {
	BookIsbn string  `gorm:"column:book_isbn"`
	Provider int     `gorm:"column:book_rating_provider"`
	Rating   float32 `gorm:"column:rating"`
}

func AddBook(db *gorm.DB, book *Book) error {
	err := db.Create(book).Error
	return err
}

func AddSentence(db *gorm.DB, sentence *Sentence) error {
	err := db.Create(sentence).Error
	return err
}

func AddWordSentence(db *gorm.DB, wordsentence *WordSentence) error {
	err := db.Create(wordsentence).Error
	return err
}

func AddBookRating(db *gorm.DB, review *BookRating) error {
	err := db.Create(review).Error
	return err
}
