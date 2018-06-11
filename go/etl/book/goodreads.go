package book

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
)

func init() {
	ReviewProviders = append(ReviewProviders, GoodReadsProvider{})
	GenreProviders = append(GenreProviders, GoodReadsProvider{})
}

type GoodReadsProvider struct{}

func createGoodReadsDoc(isbn string) (*goquery.Document, error) {
	url := "https://www.goodreads.com/search?q=%s"
	url = fmt.Sprintf(url, isbn)

	cli := util.CreateClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (GoodReadsProvider) Name() string {
	return "goodreads"
}

func (GoodReadsProvider) Rating(isbn string) (float32, error) {
	doc, err := createGoodReadsDoc(isbn)
	if err != nil {
		return 0, err
	}

	rateNode := doc.Find(".rating .average").First()
	rate := rateNode.Text()
	f, _ := strconv.ParseFloat(rate, 32)

	return float32(f), nil
}

func (GoodReadsProvider) Reviews(isbn string, maxresult int) (reviews []string, err error) {
	doc, err := createGoodReadsDoc(isbn)
	if err != nil {
		return
	}

	doc.Find(".reviewText.stacked span span").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if !strings.Contains(text, "spoiler") {
			reviews = append(reviews, text)
		}
	})

	return
}

func (GoodReadsProvider) Genre(isbn string) (genres []string, err error) {
	doc, err := createGoodReadsDoc(isbn)
	if err != nil {
		return
	}

	doc.Find(".elementList .left").Each(func(i int, s *goquery.Selection) {
		ss := s.Find(".bookPageGenreLink")
		genre := ss.Last().Text()
		if genre != "" {
			genres = append(genres, genre)
		}
	})

	return
}
