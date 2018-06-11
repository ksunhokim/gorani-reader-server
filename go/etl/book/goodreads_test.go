package book_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/etl/book"
)

func TestGoodReadsGenre(t *testing.T) {
	a := assert.New(t)
	p := book.GoodReadsProvider{}
	gs, err := p.Genre("9780316501590")
	a.Nil(err)
	sol := []string{"Light Novel", "Fantasy", "Manga", "Horror", "Action"}
	a.Equal(sol, gs)
}

func TestGoodReadsRating(t *testing.T) {
	a := assert.New(t)
	p := book.GoodReadsProvider{}
	rate, err := p.Rating("9780316501590")
	a.Nil(err)
	a.NotEqual(0, rate)
}

func TestGoodReadsReviews(t *testing.T) {
	a := assert.New(t)
	p := book.GoodReadsProvider{}
	reviews, err := p.Reviews("9780316501590", 5)
	a.Nil(err)
	a.NotEqual(0, len(reviews))
	a.NotEqual("", reviews[0])
}
