package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        []rune
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         rune // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '=':
		tok = token.New(token.ASSIGN, l.char)
	case '+':
		tok = token.New(token.PLUS, l.char)
	case '(':
		tok = token.New(token.LPAREN, l.char)
	case ')':
		tok = token.New(token.RPAREN, l.char)
	case '{':
		tok = token.New(token.LBRACE, l.char)
	case '}':
		tok = token.New(token.RBRACE, l.char)
	case ',':
		tok = token.New(token.COMMA, l.char)
	case ';':
		tok = token.New(token.SEMICOLON, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			return tok
		}
		tok = token.New(token.ILLEGAL, l.char)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func isLetter(char rune) bool {
	return 'a' <= char && char <= 'z' ||
		'A' <= char && char <= 'Z' ||
		char == '_'
}
