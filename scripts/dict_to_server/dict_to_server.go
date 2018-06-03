package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

type OWord struct {
	Word          string `json:"word"`
	Pronunciation string `json:"pronunciation,omitempty"`
	Definitions   []ODef `json:"definitions,omitempty`
}

type ODef struct {
	Definition string     `json:"definition"`
	POS        string     `json:"pos"`
	Examples   []OExample `json:"examples,omitempty`
}

type OExample struct {
	Foreign string `json:"foreign"`
	Native  string `json:"native,ompitempty"`
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

func req(url string, word OWord) error {
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}

	text, err := json.Marshal(word)
	if err != nil {
		return err
	}
	resp, _ := netClient.Post(url, "application/json", bytes.NewReader(text))
	if resp.StatusCode != 200 {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%v", string(bytes))
	}
	return nil
}

func main() {
	url := "http://127.0.0.1:5982/word"
	iwords := make(map[string]IWord)
	bytes, err := ioutil.ReadFile("raw/proned_output.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &iwords)
	if err != nil {
		panic(err)
	}

	for _, word := range iwords {
		oword := OWord{
			Word:          word.Word,
			Pronunciation: word.Pronunciation,
			Definitions:   []ODef{},
		}
		for _, def := range word.Definitions {
			odef := ODef{
				Definition: def.Definition,
				POS:        dealPOS(def.POS),
				Examples:   []OExample{},
			}
			for _, ex := range def.Examples {
				oex := OExample{
					Foreign: ex.Foreign,
					Native:  ex.Native,
				}
				odef.Examples = append(odef.Examples, oex)
			}
			oword.Definitions = append(oword.Definitions, odef)
		}

		err := req(url, oword)
		if err != nil {
			log.Println(err)
		}
	}
}
