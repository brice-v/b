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
	BACKTICK  = "`"
	QUESTION  = "?"

	LT  = "<"
	LTE = "<="
	GT  = ">"
	GTE = ">="

	EQ  = "=="
	NEQ = "!="

	POW        = "**"
	FLOORDIV   = "//"
	PLUSEQ     = "+="
	MINUSEQ    = "-="
	MULEQ      = "*="
	DIVEQ      = "/="
	BITANDEQ   = "&="
	BITOREQ    = "|="
	BITNOTEQ   = "~="
	BITXOREQ   = "^="
	MODEQ      = "%="
	BITLS      = "<<"
	BITRS      = ">>"
	PLUSPLUS   = "++"
	MINUSMINUS = "--"

	RARROW    = "=>"
	DOTDOT    = ".."
	DOTDOTDOT = "..."
	POUND     = "#"
	MLCOMMENT = "###"

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
	ELIF     = "ELIF"
	RETURN   = "RETURN"
	MACRO    = "MACRO"

	AND = "AND"
	OR  = "OR"
	NOT = "NOT"

	IN  = "IN"
	FOR = "FOR"

	// Types are keywords? this may not be correct
	INT        = "INT"
	UINT       = "UINT"
	FLOAT      = "FLOAT"
	STRINGTYPE = "STRINGTYPE"
	OBJECT     = "OBJ"
	ENUM       = "ENUM"
	LIST       = "LIST"
	MAP        = "MAP"
	CHANNEL    = "CHAN"
	ANY        = "ANY"
	BOOLEAN    = "BOOL"
	CHARACTER  = "CHAR"
	RUNE       = "RUNE"
)

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"var":    VAR,
	"val":    VAL,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"elif":   ELIF,
	"return": RETURN,
	"macro":  MACRO,
	"and":    AND,
	"or":     OR,
	"not":    NOT,
	"in":     IN,
	"for":    FOR,
	"int":    INT,
	"uint":   UINT,
	"float":  FLOAT,
	"str":    STRINGTYPE,
	"obj":    OBJECT,
	"enum":   ENUM,
	"list":   LIST,
	"map":    MAP,
	"chan":   CHANNEL,
	"any":    ANY,
	"bool":   BOOLEAN,
	"char":   CHARACTER,
	"rune":   RUNE,
}

// LookupIdent will return an identifier token if the identifier
// is not found in the keywords lookup table
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
