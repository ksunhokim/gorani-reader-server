package lang_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader/server/pkg/lang"
)

func TestSentence(t *testing.T) {
	a := assert.New(t)
	bytes, err := ioutil.ReadFile("test.txt")
	a.Nil(err)

	str := string(bytes)
	combined := strings.Replace(str, "\n", " ", -1)
	sentences := strings.Split(str, "\n")

	arr := lang.SplitSentences(combined)
	for i := range sentences {
		if arr[i] != sentences[i] {
			t.Error("SplitSentences not working")
		}
	}
}
