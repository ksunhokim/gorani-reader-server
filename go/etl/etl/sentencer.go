package etl

import (
	"github.com/sunho/gorani-reader/server/pkg/dbh"
	"github.com/sunho/gorani-reader/server/pkg/sentencer"
)

func (e *Etl) ReloadSentencer() error {
	s, err := e.createSentencer()
	if err != nil {
		return err
	}
	sentencer.SetSentencer(s)
	return nil
}

func (e *Etl) createSentencer() (*sentencer.Sentencer, error) {
	words, err := dbh.GetWords(e.Mysql)
	if err != nil {
		return nil, err
	}

	dict := sentencer.NewDictionary(words)
	dot := sentencer.DotSpecialCases{}
	past := make(map[string]string)
	complete := make(map[string]string)

	if e.Config.DotSpecialCasesJson != nil {

	}

	if e.Config.IrregularPastJson != nil {

	}

	if e.Config.IrregularCompleteJson != nil {

	}

	stemmer := sentencer.NewStemmer(past, complete)

	return sentencer.New(dict, dot, stemmer), nil
}
