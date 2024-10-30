package selector

import (
	"unicode"
)

type TokenType int

const (
	TokenTypeTag TokenType = iota
	TokenTypeID
	TokenTypeClass
	TokenTypeSpace
	TokenTypeCombinator // >, +, ~
	TokenTypeAttribute  // [attr=value]
	TokenTypeInvalid
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	input string
	pos   int
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		input: input,
		pos:   0,
	}
}

func (t *Tokenizer) nextToken() *Token {
	t.skipWhitespace()

	if t.pos >= len(t.input) {
		return nil
	}

	switch t.input[t.pos] {
	case '#':
		return t.readID()
	case '.':
		return t.readClass()
	case '[':
		return t.readAttribute()
	case '>', '+', '~':
		return t.readCombinator()
	default:
		if unicode.IsLetter(rune(t.input[t.pos])) {
			return t.readTag()
		}
	}

	return &Token{Type: TokenTypeInvalid}
}

func (t *Tokenizer) readID() *Token {
	t.pos++ // Skip #
	start := t.pos
	for t.pos < len(t.input) && isIDChar(t.input[t.pos]) {
		t.pos++
	}
	return &Token{
		Type:  TokenTypeID,
		Value: t.input[start:t.pos],
	}
}

func (t *Tokenizer) readClass() *Token {
	t.pos++ // Skip .
	start := t.pos
	for t.pos < len(t.input) && isIDChar(t.input[t.pos]) {
		t.pos++
	}
	return &Token{
		Type:  TokenTypeClass,
		Value: t.input[start:t.pos],
	}
}

func (t *Tokenizer) readTag() *Token {
	start := t.pos
	for t.pos < len(t.input) && isIDChar(t.input[t.pos]) {
		t.pos++
	}
	return &Token{
		Type:  TokenTypeTag,
		Value: t.input[start:t.pos],
	}
}

func (t *Tokenizer) readAttribute() *Token {
	start := t.pos
	t.pos++ // Skip [
	for t.pos < len(t.input) && t.input[t.pos] != ']' {
		t.pos++
	}
	if t.pos < len(t.input) {
		t.pos++ // Skip ]
	}
	return &Token{
		Type:  TokenTypeAttribute,
		Value: t.input[start:t.pos],
	}
}

func (t *Tokenizer) readCombinator() *Token {
	value := string(t.input[t.pos])
	t.pos++
	return &Token{
		Type:  TokenTypeCombinator,
		Value: value,
	}
}

func (t *Tokenizer) skipWhitespace() {
	for t.pos < len(t.input) && unicode.IsSpace(rune(t.input[t.pos])) {
		t.pos++
	}
}

func isIDChar(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) || c == '-' || c == '_'
}
