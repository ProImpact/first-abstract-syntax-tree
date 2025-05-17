package lexer

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"unicode"
)

type Lexer struct {
	buff        []rune
	index       int
	buffSize    int
	current     *Token
	generations map[int][]string
	root        *Node
}

func New(data []rune) *Lexer {
	return &Lexer{
		buff:     data,
		index:    0,
		buffSize: len(data),
		current:  nil,
	}
}

func Init(src io.Reader) (*Lexer, error) {
	l := &Lexer{}
	data, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	l.buff = bytes.Runes(data)
	l.buffSize = len(l.buff)
	l.index = -1
	l.generations = make(map[int][]string)
	l.root = &Node{}
	return l, nil
}

func (l *Lexer) Next() *Token {
	for {
		l.index++
		if l.index == l.buffSize {
			return EOF_TOKEN
		}
		if unicode.IsSpace(l.buff[l.index]) {
			continue
		}
		switch l.buff[l.index] {
		case '{':
			return OPEN_BRACKET_TOKEN
		case '}':
			return CLOSE_BRACKET_TOKEN
		default:
			if l.buff[l.index] == '"' {
				buff := bytes.Buffer{}
				_, err := buff.WriteRune(l.buff[l.index])
				if err != nil {
					log.Fatal(err)
				}
				l.index++
				for ; l.buff[l.index] != '"'; l.index++ {
					_, err = buff.WriteRune(l.buff[l.index])
					if err != nil {
						log.Fatal(err)
					}
				}
				if l.buff[l.index] == '"' { // closing tags
					_, err = buff.WriteRune(l.buff[l.index])
					if err != nil {
						log.Fatal(err)
					}
					return &Token{
						Type: STRING,
						Data: buff.String(),
					}
				}
			}
			log.Printf("unknow character at index: %d", l.index)
			return &Token{
				Type: UNKNOW,
				Data: string(l.buff[l.index]),
			}

		}
	}
}

func (l *Lexer) advance() {
	l.current = l.Next()
}

func (l *Lexer) Parse() {
	l.advance()
	l.parseBlock(l.root, 1)
}

func (l *Lexer) parseBlock(root *Node, deep int) {
	if l.current.Type != OPEN_BRACKET {
		return
	}
	l.advance()
loop:
	for ; l.current.Type != EOF; l.advance() {
		switch l.current.Type {
		case STRING:
			_, ok := l.generations[deep]
			if !ok {
				l.generations[deep] = make([]string, 0)
			}
			root.addNode(newNode(l.current.Data))
			l.generations[deep] = append(l.generations[deep], l.current.Data)
			l.parseBlock(root.childs[len(root.childs)-1], deep+1)
		case OPEN_BRACKET:
			l.parseBlock(root.childs[len(root.childs)-1], deep+1)
		case CLOSE_BRACKET:
			break loop
		}
	}
}

func (l *Lexer) PrintNodes() string {
	return l.root.String()
}

func (l Lexer) String() string {
	buff := bytes.Buffer{}
	var err error
	for k, v := range l.generations {
		_, err = buff.WriteString(fmt.Sprintf("Members of the %d generation\n\t", k))
		if err != nil {
			log.Fatal(err)
		}
		_, err = buff.WriteString(fmt.Sprintf("%v\n", v))
		if err != nil {
			log.Fatal(err)
		}

	}
	return buff.String()
}
