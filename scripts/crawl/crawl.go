package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var list = []string{}
var db *gorm.DB

type Word struct {
	gorm.Model
	Word   string
	Pron   string
	Source string
	Type   string
	Def    []Def
}

type Def struct {
	gorm.Model
	WordID  uint
	Part    string
	Def     string
	Example []Example
}

type Example struct {
	gorm.Model
	DefID uint
	Kor   string
	Eng   string
}

func getBody(url string) io.ReadCloser {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return resp.Body
}

func getDefinition(word string, url string, primary bool) {
	log.Println("definition:", url)
	body := getBody(url)
	if body == nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	typ := "word"
	if primary {
		typ = "primary"
	} else if strings.Contains(url, "idiom") {
		typ = "idiom"
	} else {
		typ = "word"
	}
	pron := doc.Find(".word_view .pron .fy .first .fnt_e16").First().Text()
	wordNode := Word{
		Word:   word,
		Pron:   pron,
		Source: url,
		Type:   typ,
	}
	defs := []Def{}
	doc.Find(".box_wrap1").Each(func(i int, s *goquery.Selection) {
		part := s.Find("h3 .fnt_syn").First().Text()
		s.Find("dl dt").Each(func(i int, ss *goquery.Selection) {
			input := ss.Find("em").First().Children().Not("a").Not("p").Not(".fnt_intro").Text()
			re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
			re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
			final := re_leadclose_whtsp.ReplaceAllString(input, "")
			final = re_inside_whtsp.ReplaceAllString(final, " ")
			def := Def{
				Part: part,
				Def:  final,
			}
			exs := []Example{}
			ss.NextUntil("dt").Each(func(i int, sss *goquery.Selection) {
				eng := ""
				kor := ""
				sss.Find("p span").Each(func(i int, ssss *goquery.Selection) {
					if i == 0 {
						eng = ssss.Text()
					} else {
						kor = ssss.Text()
					}
				})
				exs = append(exs, Example{
					Eng: eng,
					Kor: kor,
				})
			})
			def.Example = exs
			defs = append(defs, def)
		})
	})
	wordNode.Def = defs
	if db.NewRecord(wordNode) {
		db.Create(&wordNode)
		db.Save(&wordNode)
	}
}

func getQuery(word string) {
	log.Println("fetch:", word)
	body := getBody("http://endic.naver.com/search.nhn?sLn=kr&query=" + word)
	if body == nil {
		return
	}
	log.Println("query:", word)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc.Find(".word_num").Each(func(i int, s *goquery.Selection) {
		sect, _ := s.Find("h3 img").First().Attr("alt")
		if sect == "단어/숙어" {
			s.Find(".list_e2 dt span a").Each(func(i int, s *goquery.Selection) {
				href, _ := s.Attr("href")
				if s.Text() == word {
					getDefinition(s.Text(), "http://endic.naver.com"+href, true)
				} else {
					getDefinition(s.Text(), "http://endic.naver.com"+href, false)
				}

			})
		}
	})
}

func worker(input chan int) {
	for index := range input {
		word := list[index]
		log.Println(index, "/", len(list), ":", word)
		getQuery(word)
	}
}

func main() {
	tdb, err := gorm.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err)
	}
	db = tdb
	db.Model(&Def{}).AddForeignKey("word_id", "words(id)", "CASCADE", "RESTRICT")
	db.Model(&Example{}).AddForeignKey("def_id", "defs(id)", "CASCADE", "RESTRICT")
	db.Model(&Word{}).AddUniqueIndex("word_unique", "source", "word")
	txt, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	list = strings.Split(string(txt), "=")
	list = list[:len(list)-1]
	log.Printf("%d entries inputed\n", len(list))
	var wg sync.WaitGroup
	input := make(chan int, 1000)
	for i := 0; i < 500; i++ {
		go worker(input)
		wg.Add(1)
	}

	for index, _ := range list {
		input <- index
	}

	log.Println("done")
	wg.Wait()
	log.Println(time.Now().Sub(t).Minutes(), " minutes")
	time.Sleep(time.Second * 60)
	log.Println("exiting")
}
