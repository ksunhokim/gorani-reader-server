package lang

import (
	"github.com/PuerkitoBio/goquery"
)

func ExtractSentencesFromHtml(doc *goquery.Document) ([]string, error) {
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
	})
	return []string{}, nil
}

func SplitSentences(str string) []string {
	return []string{}
}
