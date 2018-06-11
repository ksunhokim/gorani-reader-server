package sentencer

import (
	"strings"

	"github.com/sunho/gorani-reader-server/go/pkg/dbh"
)

type (
	WordId int

	Sentence struct {
		Origin string
		Words  []WordId
	}

	Dictionary map[string]WordId
)

func NewDictionary(words []dbh.Word) Dictionary {
	d := Dictionary{}
	for _, word := range words {
		d[strings.ToLower(word.Word)] = WordId(word.Id)
	}

	return d
}
