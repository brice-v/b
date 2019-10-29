package lexer

import (
	"b/token"
	"testing"
)

func TestNextTokenShort(t *testing.T) {
	input := "=+(){},;`"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.BACKTICK, "`"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenLong(t *testing.T) {
	input := `var five = 5;
	var ten = 10;
	
	var add = fun(x,y) {
		x + y;
	};
	
	val result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	== !=
"foobar"
"foo bar"
[1, 2];
{"foo": "bar"}
macro(x, y) { x + y; };

^~.
<=
>=
not and or &%|
0_10
0b10_111
0B1
0x1
0X1
0o1
0O1abc
0x123_abc_123**
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.NUM, "5"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.NUM, "10"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.VAL, "val"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.FSLASH, "/"},
		{token.ASTERISK, "*"},
		{token.NUM, "5"},
		{token.SEMICOLON, ";"},
		{token.NUM, "5"},
		{token.LT, "<"},
		{token.NUM, "10"},
		{token.GT, ">"},
		{token.NUM, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUM, "5"},
		{token.LT, "<"},
		{token.NUM, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EQ, "=="},
		{token.NEQ, "!="},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.NUM, "1"},
		{token.COMMA, ","},
		{token.NUM, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.MACRO, "macro"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.HAT, "^"},
		{token.TILDE, "~"},
		{token.DOT, "."},
		{token.LTE, "<="},
		{token.GTE, ">="},
		{token.NOT, "not"},
		{token.AND, "and"},
		{token.OR, "or"},
		{token.AMPERSAND, "&"},
		{token.PERCENT, "%"},
		{token.PIPE, "|"},
		{token.NUM, "0_10"},
		{token.NUM, "0b10_111"},
		{token.NUM, "0B1"},
		{token.NUM, "0x1"},
		{token.NUM, "0X1"},
		{token.NUM, "0o1"},
		{token.NUM, "0O1"},
		{token.IDENT, "abc"},
		{token.NUM, "0x123_"},
		{token.IDENT, "abc_123"},
		{token.POW, "**"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}

func TestNextTokenShort2(t *testing.T) {
	input := `abc_fun_123; fun
	&=|=^=*=-=+=/=//>><<%=++--~=
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "abc_fun_123"},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "fun"},
		{token.BITANDEQ, "&="},
		{token.BITOREQ, "|="},
		{token.BITXOREQ, "^="},
		{token.MULEQ, "*="},
		{token.MINUSEQ, "-="},
		{token.PLUSEQ, "+="},
		{token.DIVEQ, "/="},
		{token.FLOORDIV, "//"},
		{token.BITRS, ">>"},
		{token.BITLS, "<<"},
		{token.MODEQ, "%="},
		{token.PLUSPLUS, "++"},
		{token.MINUSMINUS, "--"},
		{token.BITNOTEQ, "~="},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
