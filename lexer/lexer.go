package lexer

import (
	"interpretor/token"
)

type Lexer struct {
	input        string
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token
	lexer.clearWhiteSpace()

	switch lexer.ch {
	case '=':
		next := lexer.peekNext()
		if string(next) == token.ASSIGN {
			tok = newToken(token.EQUALS, lexer.ch)
			lexer.readChar()
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case '>':
		next := lexer.peekNext()
		if string(next) == token.ASSIGN {
			tok = newToken(token.LARGEROREQUAL, lexer.ch)
			lexer.readChar()
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case '<':
		next := lexer.peekNext()
		if string(next) == token.ASSIGN {
			tok = newToken(token.SMALLEROREQUAL, lexer.ch)
			lexer.readChar()
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case '-':
		tok = newToken(token.MINUS, lexer.ch)
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '.':
		tok = newToken(token.FULLSTOP, lexer.ch)
	case '(':
		tok = newToken(token.OPENING_BRACKET, lexer.ch)
	case ')':
		tok = newToken(token.CLOSING_BRAKET, lexer.ch)
	case '{':
		tok = newToken(token.OPENING_CURLY, lexer.ch)
	case '}':
		tok = newToken(token.CLOSING_CURLY, lexer.ch)
	case '[':
		tok = newToken(token.OPENING_SQUARE, lexer.ch)
	case ']':
		tok = newToken(token.CLOSING_SQUARE, lexer.ch)
	case '!':
		next := lexer.peekNext()
		if string(next) == token.ASSIGN {
			tok = newToken(token.NOT_EQUAL, lexer.ch)
			lexer.readChar()
		} else {
			tok = newToken(token.EXCLAMATIONMARK, lexer.ch)
		}

	case '`':
	case '\'':
	case '"':
		tok.Type = token.STRING
		tok.Value = lexer.readString()
	case ';':
		tok = newToken(token.SEMICOLUN, lexer.ch)
	case ':':
		tok = newToken(token.COLON, lexer.ch)
	case '\n':
		tok = newToken(token.NEW_LINE, lexer.ch)
	case 0:
		tok = newToken(token.EOF, lexer.ch)
		return tok
	default:
		if isLetter(lexer.ch) {
			tok.Value = lexer.getIdentifier()
			tok.Type = token.GetIdentifier(tok.Value)
			return tok
		} else if isDigit(lexer.ch) {
			tok.Value = lexer.getNumber()
			tok.Type = token.NUMBER
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) readString() string {
	position := lexer.readPosition
	quote := lexer.ch
	lexer.readChar()

	for lexer.ch != quote {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.position = lexer.readPosition
		lexer.readPosition++
		lexer.ch = lexer.input[lexer.position]
	}

}

func (lexer *Lexer) clearWhiteSpace() {
	for lexer.ch == ' ' {
		lexer.readChar()
	}
}

func (lexer *Lexer) getIdentifier() string {
	position := lexer.position

	for isLetter(lexer.ch) || isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) Peek() {

}

func (lexer *Lexer) getNumber() string {
	position := lexer.position

	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) peekNext() byte {
	return lexer.input[lexer.readPosition]
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func newToken(tokenType string, ch byte) token.Token {
	return token.Token{Type: tokenType, Value: string(ch)}
}