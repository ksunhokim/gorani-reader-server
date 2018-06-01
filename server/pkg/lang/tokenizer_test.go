package lang

import (
	"fmt"
	"strings"
	"testing"
)

func TestTokenizerEasy(t *testing.T) {
	s := "Hello from the.아"
	a := []token{
		token{tokenKindCapitalWord, "Hello"},
		token{tokenKindBlank, " "},
		token{tokenKindNormalWord, "from"},
		token{tokenKindBlank, " "},
		token{tokenKindNormalWord, "the"},
		token{tokenKindPunc, "."},
		token{tokenKindUnknown, "아"},
		token{tokenKindEof, ""},
	}
	tz := newTokenizer(strings.NewReader(s))
	toks := tz.tokenize()
	for i := range toks {
		if toks[i] != a[i] {
			t.Errorf("Fail")
		}
	}
}

func TestTokenizerQuote(t *testing.T) {
	s := `"hohoho" And then`
	a := []token{
		token{tokenKindPunc, "\""},
		token{tokenKindNormalWord, "hohoho"},
		token{tokenKindPunc, "\""},
		token{tokenKindBlank, " "},
		token{tokenKindEos, ""},
		token{tokenKindCapitalWord, "And"},
		token{tokenKindBlank, " "},
		token{tokenKindNormalWord, "then"},
		token{tokenKindEof, ""},
	}
	tz := newTokenizer(strings.NewReader(s))
	toks := tz.tokenize()
	fmt.Println(toks)
	for i := range toks {
		if toks[i] != a[i] {
			t.Errorf("Fail")
		}
	}
}
