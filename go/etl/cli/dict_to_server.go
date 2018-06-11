package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sync"

	pb "github.com/sunho/gorani-reader-server/go/proto/etl"
)

type IWord struct {
	Word          string `json:"word"`
	Pronunciation string `json:"pron"`
	Definitions   []IDef `json:"defs"`
}

type IDef struct {
	Definition string     `json:"def"`
	POS        string     `json:"pos"`
	Examples   []IExample `json:"examples"`
}

type IExample struct {
	Foreign string `json:"foreign"`
	Native  string `json:"native"`
}

//ENUM('verb', 'aux', 'tverb', 'iverb', 'noun', 'adj', 'adv', 'abr', 'prep', 'symbol', 'pronoun', 'conj', 'suffix', 'prefix', 'det')
var POSMap = map[string]string{
	"동사":  "verb",
	"조동사": "aux",
	"자동사": "iverb",
	"타동사": "tverb",
	"명사":  "noun",
	"형용사": "adj",
	"부사":  "adv",
	"약어":  "abr",
	"전치사": "prep",
	"기호":  "symbol",
	"대명사": "pronoun",
	"접속사": "conj",
	"접미사": "suffix",
	"접두사": "prefix",
	"한정사": "det",
}

func dealPOS(pos string) string {
	if p, ok := POSMap[pos]; ok {
		return p
	}
	return ""
}

func dictToServer(addr string, dict string) error {
	iwords := make(map[string]IWord)
	cli, conn, err := makeClient(addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	bytes, err := ioutil.ReadFile(dict)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &iwords)
	if err != nil {
		return err
	}

	work := make(chan *pb.Word, 10000)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for w := range work {
				_, err := cli.AddWord(context.Background(), &pb.AddWordRequest{Word: w})
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}

	stressPat := regexp.MustCompile("[0-2]+")
	i := 0
	for _, word := range iwords {
		i++
		if i%1000 == 0 {
			log.Println(i, "/", len(iwords))
		}
		pron := stressPat.ReplaceAllString(word.Pronunciation, "")
		oword := &pb.Word{
			Word:          word.Word,
			Pronunciation: pron,
			Definitions:   []*pb.Definition{},
		}
		for _, def := range word.Definitions {
			odef := &pb.Definition{
				Definition: def.Definition,
				Pos:        dealPOS(def.POS),
				Examples:   []*pb.Example{},
			}
			for _, ex := range def.Examples {
				foreign := ex.Foreign
				native := ex.Native
				if strings.Contains(foreign, "출처") {
					continue
				}
				if strings.Contains(native, "반복듣기") {
					native = ""
				}
				oex := &pb.Example{
					Foreign: foreign,
					Native:  native,
				}
				odef.Examples = append(odef.Examples, oex)
			}
			oword.Definitions = append(oword.Definitions, odef)
		}
		work <- oword
	}
	close(work)
	wg.Wait()
	return nil
}
