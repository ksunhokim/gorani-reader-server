package book

import "github.com/sunho/gorani-reader/server/pkg/dbh"

type Dictionary map[string]int

func NewDictionary(words []dbh.Word) Dictionary {
	d := Dictionary{}
	for _, word := range words {
		d[word.Word] = word.Id
	}

	return d
}
