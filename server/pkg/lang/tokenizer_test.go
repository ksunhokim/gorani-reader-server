package lang

import (
	"fmt"
	"strings"
	"testing"
)

func testTokenizer(text string, toks []Token, t *testing.T) {
	tz := NewTokenizer(strings.NewReader(text))
	tz.DotSpecialCases = []map[string]bool{
		make(map[string]bool),
		make(map[string]bool),
		make(map[string]bool),
	}
	tz.DotSpecialCases[0]["u"] = false
	tz.DotSpecialCases[1]["s"] = true
	tz.DotSpecialCases[2]["a"] = true

	toks2 := tz.Tokenize()
	fmt.Println(toks2)
	if len(toks) != len(toks2) {
		t.Errorf("Lenth not equal")
	}

	for i := range toks {
		if toks[i] != toks2[i] {
			t.Errorf("Fail")
		}
	}
}
func TestTokenizerEasy(t *testing.T) {
	s := "Hello from the.아"
	a := []Token{
		Token{TokenKindCapitalWord, "Hello"},
		Token{TokenKindBlank, " "},
		Token{TokenKindNormalWord, "from"},
		Token{TokenKindBlank, " "},
		Token{TokenKindNormalWord, "the"},
		Token{TokenKindPunc, "."},
		Token{TokenKindEos, ""},
		Token{TokenKindUnknown, "아"},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)
}

func TestTokenizerQuote(t *testing.T) {
	s := `"hohoho" And then`
	a := []Token{
		Token{TokenKindPunc, "\""},
		Token{TokenKindNormalWord, "hohoho"},
		Token{TokenKindPunc, "\""},
		Token{TokenKindEos, ""},
		Token{TokenKindBlank, " "},
		Token{TokenKindCapitalWord, "And"},
		Token{TokenKindBlank, " "},
		Token{TokenKindNormalWord, "then"},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)
}

func TestTokenizerDash(t *testing.T) {
	s := `three-dimensional text`
	a := []Token{
		Token{TokenKindNormalWord, "three-dimensional"},
		Token{TokenKindBlank, " "},
		Token{TokenKindNormalWord, "text"},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)

	s = `Three-dimensional text`
	a = []Token{
		Token{TokenKindCapitalWord, "Three-dimensional"},
		Token{TokenKindBlank, " "},
		Token{TokenKindNormalWord, "text"},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)
}

func TestTokenizerExclaim(t *testing.T) {
	s := `I am the sentence!! HOO`
	a := []Token{
		Token{TokenKindCapitalWord, "I"},
		Token{TokenKindBlank, " "},
		Token{TokenKindNormalWord, "am"},
		Token{TokenKindBlank, " "},
		Token{TokenKindNormalWord, "the"},
		Token{TokenKindBlank, " "},
		Token{TokenKindNormalWord, "sentence"},
		Token{TokenKindPunc, "!!"},
		Token{TokenKindEos, ""},
		Token{TokenKindBlank, " "},
		Token{TokenKindCapitalWord, "HOO"},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)
}

func TestTokenizerUSA(t *testing.T) {
	s := `U. S. A.`
	a := []Token{
		Token{TokenKindCapitalWord, "U.S.A."},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)

	s = `U. S. A`
	a = []Token{
		Token{TokenKindCapitalWord, "U.S.A"},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)

	s = `U. S.`
	a = []Token{
		Token{TokenKindCapitalWord, "U.S."},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)

	s = `U.`
	a = []Token{
		Token{TokenKindCapitalWord, "U"},
		Token{TokenKindPunc, "."},
		Token{TokenKindEos, ""},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)
	s = `U.D.`
	a = []Token{
		Token{TokenKindCapitalWord, "U"},
		Token{TokenKindPunc, "."},
		Token{TokenKindEos, ""},
		Token{TokenKindCapitalWord, "D"},
		Token{TokenKindPunc, "."},
		Token{TokenKindEos, ""},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)
	s = `U.U.S.A`
	a = []Token{
		Token{TokenKindCapitalWord, "U"},
		Token{TokenKindPunc, "."},
		Token{TokenKindEos, ""},
		Token{TokenKindCapitalWord, "U.S.A"},
		Token{TokenKindEof, ""},
	}
	testTokenizer(s, a, t)
}
