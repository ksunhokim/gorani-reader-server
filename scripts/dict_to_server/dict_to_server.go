package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	pb "github.com/sunho/gorani-reader/server/proto/etl"
	"google.golang.org/grpc"
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
	Foreign string `json:"first"`
	Native  string `json:"second"`
}

//ENUM('verb', 'aux', 'tverb', 'noun', 'adj', 'adv', 'abr', 'prep', 'symbol', 'pronoun', 'conj', 'suffix', 'prefix', 'det')
var POSMap = map[string]string{
	"동사":  "verb",
	"조동사": "aux",
	"자동사": "tverb",
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

func main() {
	addr := "127.0.0.1:5982"
	iwords := make(map[string]IWord)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err.Error())
	}
	defer conn.Close()

	cli := pb.NewETLClient(conn)
	bytes, err := ioutil.ReadFile("raw/proned_output.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &iwords)
	if err != nil {
		panic(err)
	}

	work := make(chan *pb.Word, 10000)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for w := range work {
				_, err := cli.AddWord(context.Background(), w)
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}

	for _, word := range iwords {
		oword := &pb.Word{
			Word:          word.Word,
			Pronunciation: word.Pronunciation,
			Definitions:   []*pb.Definition{},
		}
		for _, def := range word.Definitions {
			odef := &pb.Definition{
				Definition: def.Definition,
				Pos:        dealPOS(def.POS),
				Examples:   []*pb.Example{},
			}
			for _, ex := range def.Examples {
				oex := &pb.Example{
					Foreign: ex.Foreign,
					Native:  ex.Native,
				}
				odef.Examples = append(odef.Examples, oex)
			}
			oword.Definitions = append(oword.Definitions, odef)
		}
		work <- oword
	}
	close(work)
	wg.Wait()
}
