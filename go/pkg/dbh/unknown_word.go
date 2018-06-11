package dbh

import (
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
)

type UnknownWord struct {
	UserId         int                 `gorm:"column:user_id;primary_key" json:"-"`
	WordId         int                 `gorm:"column:word_id;primary_key" json:"word_id"`
	MemroySentence string              `gorm:"column:unknown_word_memory_sentence" json:"memory_sentence"`
	AddedDate      util.RFCTime        `gorm:"column:unknown_word_added_date" json:"added_date"`
	Sources        []UnknownWordSource `gorm:"association_save_reference:false;save_associations:false" json:"sources"`
}

func (UnknownWord) TableName() string {
	return "unknown_word"
}

type UnknownWordWithQuiz struct {
	UnknownWord
	QuizDate     *util.RFCTime `gorm:"column:quiz_date" json:"quiz_date,omitempty"`
	AverageGrade *int          `gorm:"column:average_grade" json:"average_grade,omitempty"`
}

type UnknownWordSource struct {
	UserId       int     `gorm:"column:user_id" json:"-"`
	WordId       int     `gorm:"column:word_id" json:"-"`
	DefinitionId int     `gorm:"column:definition_id" json:"definition_id"`
	Book         *string `gorm:"column:unknown_word_source_book" json:"source_book,omitemptry"`
	Sentence     *string `gorm:"column:unknown_word_source_sentence" json:"source_sentence,omitempty"`
	WordIndex    *int    `gorm:"column:unknown_word_source_word_index" json:"source_word_index,omitempty"`
}

func (UnknownWordSource) TableName() string {
	return "unknown_word_source"
}

func (u *User) GetUnknownWord(db *gorm.DB, id int) (word UnknownWord, err error) {
	err = db.First(&word, u.Id, id).Error
	if err != nil {
		return
	}
	err = word.fetchSources(db)
	return
}

func (u *User) GetUnknownWordWithQuizs(db *gorm.DB) (words []UnknownWordWithQuiz, err error) {
	err = db.Raw(`
			SELECT uw.*, q.*
			FROM
				unknown_word uw
			LEFT JOIN
				(SELECT 
					quiz.word_id,
					AVG(quiz.quiz_unknown_word_grade) AS average_grade,
					MAX(quiz.quiz_unknown_word_date) AS quiz_date
				FROM
					quiz_unknown_word quiz
				WHERE
					quiz.user_id = ?
				GROUP BY
					quiz.word_id
				) q
			ON
				uw.word_id = q.word_id
			WHERE
				uw.user_id = ?;`,
		u.Id, u.Id).
		Scan(&words).Error
	if err != nil {
		return
	}

	for i := range words {
		err = words[i].fetchSources(db)
		if err != nil {
			return
		}
	}
	return
}

func (u *User) PutUnknownWord(db *gorm.DB, word *UnknownWord) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	word.UserId = u.Id
	err = word.Delete(tx)
	if err != nil {
		return
	}

	err = tx.Create(word).Error
	if err != nil {
		return
	}

	sources := word.Sources
	for i := range sources {
		sources[i].UserId = u.Id
		sources[i].WordId = word.WordId
		err = tx.Create(&sources[i]).Error
		if err != nil {
			return
		}
	}

	return
}

func (w *UnknownWord) Delete(db *gorm.DB) error {
	err := db.Delete(w).Error
	return err
}

func (w *UnknownWord) fetchSources(db *gorm.DB) error {
	err := db.
		Where("user_id = ? AND word_id = ?", w.UserId, w.WordId).
		Find(&(w.Sources)).Error
	return err
}
