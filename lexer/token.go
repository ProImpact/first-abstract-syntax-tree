package lexer

import "fmt"

type TokenType int

const (
	_ TokenType = iota
	OPEN_BRACKET
	CLOSE_BRACKET
	STRING
	UNKNOW
	EOF
)

func (t TokenType) String() string {
	switch t {
	case OPEN_BRACKET:
		return "OPEN_BRACKET"
	case CLOSE_BRACKET:
		return "CLOSE_BRACKET"
	case STRING:
		return "STRING"
	case EOF:
		return "EOF"
	case UNKNOW:
		return "UNKNOW"
	}
	return "invalid{not in set}"
}

type Token struct {
	Type TokenType
	Data string
}

func (t Token) String() string {
	return fmt.Sprintf("[type: %s, data: %s]", t.Type, t.Data)
}

var OPEN_BRACKET_TOKEN = NewToken(OPEN_BRACKET, "{")
var CLOSE_BRACKET_TOKEN = NewToken(CLOSE_BRACKET, "}")
var EOF_TOKEN = NewToken(EOF, "0")
var UNKNOW_TOKEN = NewToken(UNKNOW, "~")

func NewToken(typ TokenType, data string) *Token {
	return &Token{
		Type: typ,
		Data: data,
	}
}
