package book

import (
	"github.com/jinzhu/gorm"
	"github.com/sunho/gorani-reader-server/go/pkg/dbh"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
)

func (b *Book) AddToDB(db *gorm.DB) (err error) {
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	b2 := dbh.Book{
		Isbn:       b.Isbn,
		Name:       b.Name,
		Author:     util.BlankToNil(b.Author),
		CoverImage: util.BlankToNil(b.Cover.Object),
	}
	err = dbh.AddBook(tx, &b2)
	if err != nil {
		return
	}

	for _, sen := range b.Sentences {
		sen2 := dbh.Sentence{
			BookIsbn: util.BlankToNil(b.Isbn),
			Sentence: sen.Origin,
		}
		err = dbh.AddSentence(tx, &sen2)
		if err != nil {
			return
		}

		for i, word := range sen.Words {
			wordsen := dbh.WordSentence{
				WordId:     int(word),
				SentenceId: sen2.Id,
				Position:   i,
			}
			err = dbh.AddWordSentence(tx, &wordsen)
			if err != nil {
				return
			}
		}
	}

	for _, rating := range b.Ratings {
		rat := dbh.BookRating{
			BookIsbn: b.Isbn,
			Provider: rating.Provider,
			Rating:   rating.Rating,
		}
		err = dbh.AddBookRating(tx, &rat)
		if err != nil {
			return
		}
	}

	for _, review := range b.Reviews {
		rev := dbh.BookReview{
			BookIsbn: b.Isbn,
			Provider: review.Provider,
			Review:   review.Comment,
		}
		err = dbh.AddBookReview(tx, &rev)
		if err != nil {
			return
		}
	}

	for _, genre := range b.Genre {
		gen, err := dbh.GetGenreByName(tx, genre)
		if err != nil {
			gen.Name = genre
			err = dbh.AddGenre(tx, &gen)
			if err != nil {
				return err
			}
		}
		gen2 := dbh.BookGenre{
			BookIsbn: b.Isbn,
			Genre:    gen.Code,
		}

		err = dbh.AddBookGenre(tx, &gen2)
		if err != nil {
			return err
		}
	}

	return
}
