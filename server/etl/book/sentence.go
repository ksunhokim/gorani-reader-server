package book

import (
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/sunho/gorani-reader/server/pkg/lang"
)

func parseSentences(dict Dictionary, r io.Reader) ([]Sentence, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	strs, err := lang.ExtractSentencesFromHtml(doc)
	for _, str := range strs {

	}
	return out, nil
}
