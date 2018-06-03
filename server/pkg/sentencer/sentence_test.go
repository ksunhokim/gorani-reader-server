package sentencer_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/sentencer"
)

func splitSentences(str string) (sens []string) {
	t := sentencer.NewTokenizer(strings.NewReader(str))
	toks := t.Tokenize()

	n := 0
	sens = []string{""}
	for _, tok := range toks {
		if tok.Kind == sentencer.TokenKindEos {
			sens = append(sens, "")
			n++
			continue
		}
		sens[n] += tok.Lit
	}
	return
}

func TestSentence(t *testing.T) {
	a := assert.New(t)
	bytes, err := ioutil.ReadFile("test.txt")
	a.Nil(err)

	str := string(bytes)
	sentences := strings.Split(str, "\n")
	combined := strings.Replace(str, "\n", "", -1)

	arr := splitSentences(combined)
	a.Equal(len(sentences), len(arr)-1)
	fmt.Println(strings.Join(arr, "\n"))
	for i := range sentences {
		if arr[i] != sentences[i] {
			t.Error("SplitSentences not working")
		}
	}
}
