package dbh

import "github.com/jinzhu/gorm"

type Book struct {
	Isbn       string  `gorm:"column:book_isbn;primary_key"`
	Name       string  `gorm:"column:book_name"`
	Author     *string `gorm:"column:book_author"`
	CoverImage *string `gorm:"column:book_cover_image"`
}

func (Book) TableName() string {
	return "book"
}

type Sentence struct {
	BookIsbn *string `gorm:"column:book_isbn"`
	Id       int     `gorm:"column:sentence_id;primary_key"`
	Sentence string  `gorm:"column:sentence"`
}

func (Sentence) TableName() string {
	return "sentence"
}

// TODO insert stream
type WordSentence struct {
	WordId     int `gorm:"column:word_id"`
	SentenceId int `gorm:"column:sentence_id"`
	Position   int `gorm:"column:word_position"`
}

func (WordSentence) TableName() string {
	return "word_sentence"
}

type BookRating struct {
	BookIsbn string  `gorm:"column:book_isbn"`
	Provider string  `gorm:"column:book_rating_provider"`
	Rating   float32 `gorm:"column:rating"`
}

func (BookRating) TableName() string {
	return "book_rating"
}

type BookGenre struct {
	BookIsbn string `gorm:"column:book_isbn"`
	Genre    string `gorm:"column:genre"`
}

func (BookGenre) TableName() string {
	return "book_genre"
}

type BookReview struct {
	BookIsbn string `gorm:"column:book_isbn"`
	Provider string `gorm:"column:book_review_provider"`
	Review   string `gorm:"column:review"`
}

func (BookReview) TableName() string {
	return "book_review"
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

func AddBookGenre(db *gorm.DB, genre *BookGenre) error {
	err := db.Create(genre).Error
	return err
}

func AddBookReview(db *gorm.DB, review *BookReview) error {
	err := db.Create(review).Error
	return err
}

func FindSentences(db *gorm.DB, word1 Word, word2 Word, maxdistance int) (sentences []Sentence, err error) {
	err = db.Raw(`
		SELECT sentence.*
		FROM
			word_sentence a
		INNER JOIN
			word_sentence b
		ON 
			a.word_id = ? AND
			b.word_id = ? AND
			a.sentence_id = b.sentence_id AND
			ABS(a.word_position - b.word_position) <= ? 
		INNER JOIN
			sentence
		ON
			a.sentence_id = sentence.sentence_id
		GROUP BY
			sentence_id;`, word1.Id, word2.Id, maxdistance).
		Scan(&sentences).Error
	return
}

func (u *User) GetKnownDegreeOfSentence(db *gorm.DB, sen Sentence) (known int, total int, err error) {
	err = db.DB().QueryRow(`
		SELECT COUNT(ws.word_id), COUNT(ws2.word_id)
		FROM
			word_sentence ws,
			word_sentence ws2
		INNER JOIN
			known_word nw
		ON
			ws2.word_id = nw.word_id
		WHERE
			ws.sentence_id = ? AND
			ws2.sentence_id = ? AND
			nw.user_id = ?;`, sen.Id, sen.Id, u.Id).
		Scan(&total, &known)
	return
}
