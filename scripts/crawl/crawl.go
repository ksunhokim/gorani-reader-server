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

	"github.com/jmoiron/sqlx"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

var list = []string{}
var db *sqlx.DB

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

func getDefinition(word string, url string, primary bool, source string) {
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
	} else if strings.Contains(url, "idiomId") {
		typ = "idiom"
	} else {
		typ = "word"
	}
	pron := doc.Find(".word_view .pron .fy .first .fnt_e16").First().Text()
	re, err := db.Exec("INSERT INTO words(word, pron, source, type) VALUES (?,?,?,?)",
		word, pron, source, typ)
	if err != nil {
		log.Println(err)
		return
	}
	wordID, _ := re.LastInsertId()
	doc.Find(".box_wrap1").Each(func(i int, s *goquery.Selection) {
		part := s.Find("h3 .fnt_syn").First().Text()
		s.Find("dl dt").Each(func(i int, ss *goquery.Selection) {
			input := ss.Find("em").First().Children().Not("a").Not("p").Not(".fnt_intro").Text()
			re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
			re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
			final := re_leadclose_whtsp.ReplaceAllString(input, "")
			final = re_inside_whtsp.ReplaceAllString(final, " ")
			re, err = db.Exec("INSERT INTO defs(def, part, word_id) VALUES (?,?,?)",
				final, part, wordID)
			if err != nil {
				log.Println(err)
				return
			}
			defID, _ := re.LastInsertId()
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
				db.Exec("INSERT INTO examples(kor, eng, def_id) VALUES (?,?,?)",
					kor, eng, defID)
			})
		})
	})
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
				pat := regexp.MustCompile("(entryId=([^&]+))|(idiomId=([^&]+))")
				res := pat.FindAllStringSubmatch(href, -1)
				if len(res) == 1 {
					if s.Text() == word {
						getDefinition(s.Text(), "http://endic.naver.com"+href, true, res[0][0])
					} else {
						getDefinition(s.Text(), "http://endic.naver.com"+href, false, res[0][0])
					}
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
	tdb, err := sqlx.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		panic(err)
	}
	db = tdb
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
	close(input)
	log.Println("done")
	wg.Done()
	log.Println(time.Now().Sub(t).Minutes(), " minutes")
	time.Sleep(time.Second * 60)
	log.Println("exiting")
}
