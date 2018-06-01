package lang

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

type tokenKind int

const (
	tokenKindCapitalWord = iota
	tokenKindNormalWord
	tokenKindPunc
	tokenKindUnknown
	tokenKindBlank
	tokenKindEos
	tokenKindEof
)

type token struct {
	kind tokenKind
	lit  string
}

type tokenizer struct {
	s *bufio.Scanner
}

func newTokenizer(r io.Reader) *tokenizer {
	s := bufio.NewScanner(r)
	s.Split(split)
	return &tokenizer{
		s: s,
	}
}

func (t *tokenizer) tokenize() (toks []token) {
	for {
		tok := t.scan()
		toks = append(toks, tok)
		if tok.kind == tokenKindEof {
			return
		}

		if strings.HasSuffix(tok.lit, "\"") {
			toks = append(toks, t.scanQuoteEos(toks)...)
			continue
		}

		if strings.HasSuffix(tok.lit, "”") {
			toks = append(toks, t.scanQuoteEos(toks)...)
			continue
		}
	}
}

func backseekToken(toks []token, kind tokenKind, n int) token {
	for i := 1; i <= n; i++ {
		if tok := toks[len(toks)-i]; tok.kind == kind {
			return tok
		}
	}
	return token{tokenKindEof, ""}
}

func (t *tokenizer) scanQuoteEos(toks []token) (out []token) {
	for {
		tok := t.scan()
		out = append(out, tok)
		switch tok.kind {
		case tokenKindBlank:
			continue
		case tokenKindCapitalWord:
			if len(toks) <= 2 ||
				backseekToken(toks, tokenKindEos, 3).kind != tokenKindEof {
			} else {
				out = append(out[:len(out)-1], token{tokenKindEos, ""}, out[len(out)-1])
			}
			return
		case tokenKindEof:
			out = out[:len(out)-1]
			return
		default:
			return
		}
	}
}

func (t *tokenizer) scan() (tok token) {
	if t.s.Scan() {
		s := t.s.Bytes()
		u, _ := utf8.DecodeRune(s)
		tok.lit = string(s)

		if unicode.IsSpace(u) {
			fmt.Println(u)
			tok.kind = tokenKindBlank
			return
		}

		if unicode.IsPunct(u) {
			tok.kind = tokenKindPunc
			return
		}

		if unicode.IsUpper(u) {
			tok.kind = tokenKindCapitalWord
			return
		}

		if unicode.IsLower(u) {
			tok.kind = tokenKindNormalWord
			return
		}

		tok.kind = tokenKindUnknown
		return
	}

	return token{kind: tokenKindEof}
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

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
