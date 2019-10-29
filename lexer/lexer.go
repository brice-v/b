package lexer

import (
	"b/token"
)

type Lexer struct {
	input   string
	pos     int // current position
	nextPos int
	ch      byte // current char under examination

	inputLen int
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.pos
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.pos
	nextChar := l.peekChar()
	if l.ch == '0' {
		if nextChar == 'b' ||
			nextChar == 'B' ||
			nextChar == 'x' ||
			nextChar == 'X' ||
			nextChar == 'o' ||
			nextChar == 'O' {
			// skipping over the 2 chars
			l.readChar()
			l.readChar()
		}
	}
	for {
		ok := isDigit(l.ch) ||
			(l.ch == '.' && isDigit(l.peekChar())) ||
			(isDigit(l.ch) && l.peekChar() == '.')
		if !ok {
			break
		}
		l.readChar()
	}
	return l.input[position:l.pos]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		if l.peekChar() == '+' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUSPLUS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUSEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.PLUS, l.ch)
		}
	case '-':
		if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUSMINUS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUSEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.MINUS, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		if l.peekChar() == '/' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.FLOORDIV, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.DIVEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.FSLASH, l.ch)
		}
	case '*':
		if l.peekChar() == '*' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.POW, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MULEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASTERISK, l.ch)
		}
	case '<':
		if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.BITLS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.BITRS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '^':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.BITXOREQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.HAT, l.ch)
		}
	case '~':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.BITNOTEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.TILDE, l.ch)
		}
	case '.':
		tok = newToken(token.DOT, l.ch)
	case '%':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MODEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.PERCENT, l.ch)
		}
	case '&':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.BITANDEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.AMPERSAND, l.ch)
		}
	case '|':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.BITOREQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.PIPE, l.ch)
		}
	case '`':
		tok = newToken(token.BACKTICK, l.ch)
	case '?':
		tok = newToken(token.QUESTION, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.NUM
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekCharAtOffset(offset int) byte {
	nextPos := l.pos + offset
	if nextPos >= l.inputLen {
		return 0
	} else {
		return l.input[nextPos]
	}
}

func (l *Lexer) peekChar() byte {
	if l.nextPos >= l.inputLen {
		return 0
	} else {
		return l.input[l.nextPos]
	}
}

// readChar consumes the character
func (l *Lexer) readChar() {
	if l.nextPos >= l.inputLen {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPos]
	}
	l.pos = l.nextPos
	l.nextPos++
}

func (l *Lexer) readString() string {
	position := l.pos + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return l.input[position:l.pos]
}

func New(input string) *Lexer {
	inputLen := len(input)
	l := &Lexer{input: input, inputLen: inputLen}
	l.readChar()
	return l
}
