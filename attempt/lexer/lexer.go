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

	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(char) + string(l.char)}
		} else {
			tok = token.New(token.ASSIGN, l.char)
		}
	case '+':
		tok = token.New(token.PLUS, l.char)
	case '-':
		tok = token.New(token.MINUS, l.char)
	case '*':
		tok = token.New(token.ASTERISK, l.char)
	case '/':
		tok = token.New(token.SLASH, l.char)
	case '<':
		tok = token.New(token.LT, l.char)
	case '>':
		tok = token.New(token.GT, l.char)
	case '!':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(char) + string(l.char)}
		} else {
			tok = token.New(token.BANG, l.char)
		}
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
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = token.New(token.ILLEGAL, l.char)
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func isLetter(char rune) bool {
	return 'a' <= char && char <= 'z' ||
		'A' <= char && char <= 'Z' ||
		char == '_'
}

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}
