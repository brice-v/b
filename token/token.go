package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT  = "IDENT"
	NUM    = "NUM"
	STRING = "STRING"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	FSLASH    = "/"
	HAT       = "^"
	AMPERSAND = "&"
	TILDE     = "~"
	PIPE      = "|"
	PERCENT   = "%"

	LT  = "<"
	LTE = "<="
	GT  = ">"
	GTE = ">="

	EQ  = "=="
	NEQ = "!="

	POW = "**"

	// List Tokens
	LBRACKET = "["
	RBRACKET = "]"

	// Hash Token
	COLON = ":"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	VAL      = "VAL"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	MACRO    = "MACRO"

	AND = "AND"
	OR  = "OR"
	NOT = "NOT"
)

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"var":    VAR,
	"val":    VAL,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
	"and":    AND,
	"or":     OR,
	"not":    NOT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
