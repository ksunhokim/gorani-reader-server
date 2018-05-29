package dbh

import "github.com/jinzhu/gorm"

type Book struct {
	Id   int    `gorm:"column:book_id"`
	Name string `gorm:"column:book_name"`
}

type Sentence struct {
	Id       int    `gorm:"column:sentence_id"`
	Sentence string `gorm:"column:sentence"`
	BookId   *int   `gorm:"column:book_id"`
}

type WordSentence struct {
	WordId     int `gorm:"column:word_id"`
	SentenceId int `gorm:"column:sentence_id"`
	Position   int `gorm:"column:position"`
}

func AddSentence(db *gorm.DB, sentence *Sentence) error {
	err := db.Create(sentence).Error
	return err
}

func AddWordSentence(db *gorm.DB, wordsentence *WordSentence) error {
	err := db.Create(wordsentence).Error
	return err
}
