package lang

import (
	"bufio"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenKind int

const (
	TokenKindEof = iota
	TokenKindEos
	TokenKindCapitalWord
	TokenKindNormalWord
	TokenKindPunc
	TokenKindUnknown
	TokenKindBlank
)

type Token struct {
	Kind TokenKind
	Lit  string
}

type Tokenizer struct {
	DotSpecialCases []map[string]bool
	index           int
	s               *bufio.Scanner
	toks            []Token
	buffer          []Token
}

func NewTokenizer(r io.Reader) *Tokenizer {
	s := bufio.NewScanner(r)
	s.Split(split)
	return &Tokenizer{
		DotSpecialCases: []map[string]bool{},
		s:               s,
		buffer:          []Token{},
		toks:            []Token{},
	}
}

func isWordToken(t Token) bool {
	return t.Kind == TokenKindCapitalWord || t.Kind == TokenKindNormalWord
}

func (t *Tokenizer) Tokenize() []Token {
	for {
		tok := t.read()
		t.toks = append(t.toks, tok)
		if tok.Kind == TokenKindEof {
			return t.toks
		}

		if tok.Lit == "-" && isWordToken(t.lastToken(1)) {
			tok2 := t.read()
			if isWordToken(tok2) {
				t.toks = t.toks[:len(t.toks)-1]
				t.toks[len(t.toks)-1].Lit += tok.Lit + tok2.Lit
				continue
			}
			t.unread()
		}

		if tok.Lit == "." && isWordToken(t.lastToken(1)) {
			if lit := t.scanDotSpecialCase(); lit != "" {
				t.toks = t.toks[:len(t.toks)-1]
				t.toks[len(t.toks)-1].Lit = lit
				continue
			}
		}

		if strings.HasSuffix(tok.Lit, "”") || strings.HasSuffix(tok.Lit, "\"") {
			if t.scanQuoteEos() {
				t.toks = append(t.toks, Token{TokenKindEos, ""})
				continue
			}
		}

		if strings.HasSuffix(tok.Lit, ".") || strings.HasSuffix(tok.Lit, "?") || strings.HasSuffix(tok.Lit, "!") {
			t.toks = append(t.toks, Token{TokenKindEos, ""})
			continue
		}
	}
}

func (t *Tokenizer) backseekToken(Kind TokenKind, n int) Token {
	if n >= len(t.toks) {
		return Token{TokenKindEof, ""}
	}

	for i := 1; i <= n; i++ {
		if tok := t.toks[len(t.toks)-i]; tok.Kind == Kind {
			return tok
		}
	}
	return Token{TokenKindEof, ""}
}

func (t *Tokenizer) lastToken(n int) Token {
	if len(t.toks) <= n {
		return Token{TokenKindEof, ""}
	}

	return t.toks[len(t.toks)-n-1]
}

func (t *Tokenizer) hasDotSpecialCase(index int, word string) (bool, bool) {
	if index >= len(t.DotSpecialCases) {
		return false, false
	}

	v, ok := t.DotSpecialCases[index][strings.ToLower(word)]
	return v, ok
}

func (t *Tokenizer) scanDotSpecialCase() (lit string) {
	tok := t.lastToken(1)

	var (
		index int = 1
		clit  string
		readn int
		state int
	)
	if v, ok := t.hasDotSpecialCase(0, tok.Lit); !ok {
		return
	} else {
		if v {
			lit = tok.Lit + "."
		} else {
			clit = tok.Lit + "."
		}
	}

	for {
		tok = t.read()
		readn++
		if state == 0 {
			if isWordToken(tok) {
				t.unread()
				readn--
				state = 1
			} else if tok.Kind == TokenKindBlank {
				state = 1
			} else {
				break
			}
		} else if state == 1 {
			if isWordToken(tok) {
				if v, ok := t.hasDotSpecialCase(index, tok.Lit); ok {
					index++
					state = 2

					if v {
						lit += clit + tok.Lit
						clit = ""
						readn = 0
					} else {
						clit += tok.Lit
					}
				} else {
					break
				}
			} else {
				break
			}
		} else if state == 2 {
			if tok.Lit == "." {
				if clit == "" {
					lit += "."
					readn = 0
				} else {
					clit += "."
				}
				state = 0
			} else {
				break
			}
		}
	}

	for i := 0; i < readn; i++ {
		t.unread()
	}
	return
}

func (t *Tokenizer) scanQuoteEos() bool {
	for {
		tok := t.read()
		defer t.unread()

		switch tok.Kind {
		case TokenKindBlank:
			continue
		case TokenKindCapitalWord:
			if len(t.toks) >= 3 && t.backseekToken(TokenKindEos, 3).Kind == TokenKindEof {
				return true
			}
			return false
		default:
			return false
		}
	}
}

func (t *Tokenizer) unread() {
	if t.index != 0 {
		t.index--
	}
}

func (t *Tokenizer) read() Token {
	if t.index == len(t.buffer) {
		var tok Token
		if t.s.Scan() {
			s := t.s.Bytes()
			u, _ := utf8.DecodeRune(s)
			tok.Lit = string(s)

			if unicode.IsSpace(u) {
				tok.Kind = TokenKindBlank
			} else if unicode.IsPunct(u) {
				tok.Kind = TokenKindPunc
			} else if unicode.IsUpper(u) {
				tok.Kind = TokenKindCapitalWord
			} else if unicode.IsLower(u) {
				tok.Kind = TokenKindNormalWord
			} else {
				tok.Kind = TokenKindUnknown
			}
		} else {
			tok = Token{TokenKindEof, ""}
		}
		t.buffer = append(t.buffer, tok)
	}

	tok := t.buffer[t.index]
	t.index++
	return tok
}

func isAlphabet(u rune) bool {
	return 'a' <= u && u <= 'z' || 'A' <= u && u <= 'Z'
}

func isUnknown(u rune) bool {
	return !unicode.IsPunct(u) && !isAlphabet(u) && !unicode.IsSpace(u)
}

func splitOne(eof bool, data []byte, start int, u rune, shouldAdvance func(u rune) bool) (int, []byte, error) {
	for width, i := 0, start; i < len(data) && shouldAdvance(u); i += width {
		u, width = utf8.DecodeRune(data[i:])
		if !shouldAdvance(u) {
			return i, data[:i], nil
		}
	}

	if eof {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func split(data []byte, atEOF bool) (advance int, Token []byte, err error) {
	u, start := utf8.DecodeRune(data)
	if unicode.IsSpace(u) {
		return splitOne(atEOF, data, start, u, unicode.IsSpace)
	}

	if unicode.IsPunct(u) {
		return splitOne(atEOF, data, start, u, unicode.IsPunct)
	}

	if isAlphabet(u) {
		return splitOne(atEOF, data, start, u, isAlphabet)
	}

	if isUnknown(u) && start != 0 {
		return splitOne(atEOF, data, start, u, isUnknown)
	}

	return
}
